package main

import (
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/collector"
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/domain"
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/storage"
	"context"
	"errors"
	"math"
	"runtime"
	"slices"
	"testing"
	"time"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

// aguardar repete a condição até ser verdadeira ou o prazo (2s) acabar.
func aguardar(t *testing.T, descricao string, cond func() bool) {
	t.Helper()
	limite := time.Now().Add(2 * time.Second)
	for time.Now().Before(limite) {
		if cond() {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatalf("condição não atendida a tempo: %s", descricao)
}

func totalLojas(sols []domain.SolicitacaoBusca) int {
	n := 0
	for _, s := range sols {
		n += len(s.Lojas)
	}
	return n
}

func TestIntegracaoEngineColetaConcorrente(t *testing.T) {
	testes := []struct {
		nome           string
		solicitacoes   []domain.SolicitacaoBusca
		numWorkers     int
		timeout        time.Duration
		qtdCotacoesEsp int
		noCacheEsp     []string // "Loja|Produto" esperados no cache (só nos cenários de sucesso)
		deveDarTimeout bool
	}{
		{
			nome: "Execução com Sucesso em Lote Normal",
			solicitacoes: []domain.SolicitacaoBusca{
				{ID: 1, Produto: "Mouse", Lojas: []string{"Amazon", "Kabum"}},
				{ID: 2, Produto: "Teclado", Lojas: []string{"Shopee"}},
			},
			numWorkers:     2,
			timeout:        500 * time.Millisecond,
			qtdCotacoesEsp: 3,
			noCacheEsp:     []string{"Amazon|Mouse", "Kabum|Mouse", "Shopee|Teclado"},
			deveDarTimeout: false,
		},
		{
			nome: "Mais workers que solicitações",
			solicitacoes: []domain.SolicitacaoBusca{
				{ID: 1, Produto: "Monitor", Lojas: []string{"Magalu"}},
			},
			numWorkers:     8,
			timeout:        500 * time.Millisecond,
			qtdCotacoesEsp: 1,
			noCacheEsp:     []string{"Magalu|Monitor"},
			deveDarTimeout: false,
		},
		{
			nome:           "Lote sem solicitações",
			solicitacoes:   nil,
			numWorkers:     2,
			timeout:        500 * time.Millisecond,
			qtdCotacoesEsp: 0,
			noCacheEsp:     nil,
			deveDarTimeout: false,
		},
		{
			nome: "Timeout Curto (Aborta com Cancelamento)",
			solicitacoes: []domain.SolicitacaoBusca{
				{ID: 1, Produto: "Servidor", Lojas: []string{"L1", "L2", "L3", "L4", "L5"}},
			},
			numWorkers:     2,
			timeout:        5 * time.Millisecond,
			deveDarTimeout: true,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			cache := storage.NovoCacheCotacoes()
			engine := collector.NovaEngineColeta(cache, tt.numWorkers, tt.timeout)

			cotacoes, erros := engine.ExecutarLote(context.Background(), tt.solicitacoes)

			// Nenhuma cotação vazia pode vazar para o resultado (vale para todos os cenários)
			for _, c := range cotacoes {
				if c.Loja == "" {
					t.Errorf("cotação vazia no resultado: %+v", c)
				}
			}

			if tt.deveDarTimeout {
				// errors.Join de lista vazia é nil, então "sem erro nenhum" também reprova aqui
				if !errors.Is(errors.Join(erros...), context.DeadlineExceeded) {
					t.Fatalf("erros = %v; esperado conter %v", erros, context.DeadlineExceeded)
				}
				// Não exige "exatamente 0" (dependeria de timing): basta não ter completado
				if len(cotacoes) >= totalLojas(tt.solicitacoes) {
					t.Errorf("coletou %d cotações; o timeout deveria ter abortado antes de completar", len(cotacoes))
				}
				return
			}

			if len(erros) > 0 {
				t.Fatalf("não esperava erros, obteve: %v", erros)
			}
			assertIgual(t, len(cotacoes), tt.qtdCotacoesEsp)

			// O cache tem que refletir exatamente o que foi coletado
			var noCache []string
			for _, c := range cache.ListarTodos() {
				noCache = append(noCache, c.Loja+"|"+c.Produto)
			}
			slices.Sort(noCache)
			esperado := slices.Clone(tt.noCacheEsp)
			slices.Sort(esperado)
			if !slices.Equal(noCache, esperado) {
				t.Errorf("cache = %v; esperado %v", noCache, esperado)
			}

			// Cada cotação tem que casar com a sua solicitação (resultados de workers
			// diferentes não podem se misturar). Depende da simulação: preço = 199.90 + ID.
			idPorProduto := map[string]int{}
			for _, s := range tt.solicitacoes {
				idPorProduto[s.Produto] = s.ID
			}
			for _, c := range cotacoes {
				esperadoPreco := 199.90 + float64(idPorProduto[c.Produto])
				if math.Abs(c.Preco-esperadoPreco) > 1e-9 {
					t.Errorf("%s/%s: preço %.2f; esperado %.2f", c.Loja, c.Produto, c.Preco, esperadoPreco)
				}
			}
		})
	}
}

func TestIntegracaoEngineContextoPaiCancelado(t *testing.T) {
	cache := storage.NovoCacheCotacoes()
	engine := collector.NovaEngineColeta(cache, 2, 500*time.Millisecond)
	sols := []domain.SolicitacaoBusca{
		{ID: 1, Produto: "Mouse", Lojas: []string{"Amazon", "Kabum"}},
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	cotacoes, erros := engine.ExecutarLote(ctx, sols)

	if !errors.Is(errors.Join(erros...), context.Canceled) {
		t.Errorf("erros = %v; esperado conter %v", erros, context.Canceled)
	}
	assertIgual(t, len(cotacoes), 0)
}

// Rodar o lote várias vezes, inclusive abortando por timeout, não pode deixar goroutines soltas.
func TestIntegracaoEngineSemLeaks(t *testing.T) {
	sols := []domain.SolicitacaoBusca{
		{ID: 1, Produto: "Servidor", Lojas: []string{"L1", "L2", "L3", "L4", "L5"}},
		{ID: 2, Produto: "Mouse", Lojas: []string{"Amazon", "Kabum"}},
	}
	antes := runtime.NumGoroutine()

	for i := 0; i < 20; i++ {
		cache := storage.NovoCacheCotacoes()
		engine := collector.NovaEngineColeta(cache, 3, 5*time.Millisecond)
		engine.ExecutarLote(context.Background(), sols)
	}

	aguardar(t, "goroutines voltaram ao patamar inicial", func() bool {
		return runtime.NumGoroutine() <= antes
	})
}
