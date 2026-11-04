package main

import (
	"context"
	"errors"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

// fakeComAtrasos simula lojas com atrasos controlados e conta quantas foram canceladas.
func fakeComAtrasos(atrasos map[string]time.Duration, canceladas *atomic.Int64) buscador {
	return func(ctx context.Context, loja string) ResultadoBusca {
		timer := time.NewTimer(atrasos[loja])
		defer timer.Stop()

		select {
		case <-timer.C:
			return ResultadoBusca{Loja: loja, Preco: float64(atrasos[loja] / time.Millisecond)}
		case <-ctx.Done():
			canceladas.Add(1)
			return ResultadoBusca{Loja: loja, Err: ctx.Err()}
		}
	}
}

// aguardar repete a condição até ela ser verdadeira ou o prazo acabar.
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

func TestBuscarPrimeiroPreco(t *testing.T) {
	lento := 10 * time.Second // só termina se for cancelada

	testes := []struct {
		nome       string
		lojas      []string
		atrasos    map[string]time.Duration
		esperada   string
		canceladas int64
	}{
		{
			nome:       "a mais rápida vence",
			lojas:      []string{"A", "B", "C", "D", "E"},
			atrasos:    map[string]time.Duration{"A": lento, "B": 10 * time.Millisecond, "C": lento, "D": lento, "E": lento},
			esperada:   "B",
			canceladas: 4,
		},
		{
			nome:       "a rápida é a última da lista",
			lojas:      []string{"A", "B", "C"},
			atrasos:    map[string]time.Duration{"A": lento, "B": lento, "C": 10 * time.Millisecond},
			esperada:   "C",
			canceladas: 2,
		},
		{
			nome:       "uma loja só: ninguém para cancelar",
			lojas:      []string{"A"},
			atrasos:    map[string]time.Duration{"A": 5 * time.Millisecond},
			esperada:   "A",
			canceladas: 0,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			var canceladas atomic.Int64

			r, err := buscarPrimeiro(context.Background(), tt.lojas, fakeComAtrasos(tt.atrasos, &canceladas))

			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}
			assertIgual(t, r.Loja, tt.esperada)

			// As perdedoras só saem pelo cancelamento: o lento é 10s, o teste espera no máximo 2s
			aguardar(t, "perdedoras canceladas", func() bool { return canceladas.Load() == tt.canceladas })
		})
	}
}

func TestBuscarPrimeiroPrecoListaVazia(t *testing.T) {
	for _, lojas := range [][]string{nil, {}} {
		r, err := BuscarPrimeiroPreco(context.Background(), lojas)

		if !errors.Is(err, ErrNenhumaLojaFornecida) {
			t.Errorf("erro = %v; esperado %v", err, ErrNenhumaLojaFornecida)
		}
		assertIgual(t, r.Loja, "")
	}
}

func TestBuscarPrimeiroPrecoContextoPaiCancelado(t *testing.T) {
	atrasos := map[string]time.Duration{"A": 10 * time.Second, "B": 10 * time.Second}
	var canceladas atomic.Int64

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(20*time.Millisecond, cancel)

	_, err := buscarPrimeiro(ctx, []string{"A", "B"}, fakeComAtrasos(atrasos, &canceladas))

	if !errors.Is(err, context.Canceled) {
		t.Errorf("erro = %v; esperado context.Canceled", err)
	}
	aguardar(t, "as duas goroutines canceladas", func() bool { return canceladas.Load() == 2 })
}

func TestBuscarPrimeiroPrecoTodasFalham(t *testing.T) {
	falha := func(ctx context.Context, loja string) ResultadoBusca {
		return ResultadoBusca{Loja: loja, Err: errors.New("fora do ar")}
	}

	_, err := buscarPrimeiro(context.Background(), []string{"A", "B", "C"}, falha)

	if !errors.Is(err, ErrTodasFalharam) {
		t.Errorf("erro = %v; esperado %v", err, ErrTodasFalharam)
	}
}

func TestBuscarPrimeiroPrecoLojaFalhaMasOutraResponde(t *testing.T) {
	mista := func(ctx context.Context, loja string) ResultadoBusca {
		if loja == "quebrada" {
			return ResultadoBusca{Loja: loja, Err: errors.New("500")} // falha imediata
		}
		time.Sleep(20 * time.Millisecond)
		return ResultadoBusca{Loja: loja, Preco: 99}
	}

	r, err := buscarPrimeiro(context.Background(), []string{"quebrada", "boa"}, mista)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	assertIgual(t, r.Loja, "boa") // a falha rápida não "ganha" a corrida
}

// Sem vazamento: depois de retornar e cancelar, o número de goroutines volta ao que era
func TestBuscarPrimeiroPrecoSemVazamento(t *testing.T) {
	antes := runtime.NumGoroutine()

	for i := 0; i < 20; i++ { // repete para um vazamento pequeno aparecer na soma
		var canceladas atomic.Int64
		atrasos := map[string]time.Duration{"A": 10 * time.Millisecond, "B": 10 * time.Second, "C": 10 * time.Second}
		buscarPrimeiro(context.Background(), []string{"A", "B", "C"}, fakeComAtrasos(atrasos, &canceladas))
	}

	aguardar(t, "goroutines voltaram ao patamar inicial", func() bool {
		return runtime.NumGoroutine() <= antes
	})
}

// A versão simulada de verdade: a loja mais rápida do mapa é a MercadoLivre
func TestBuscarPrimeiroPrecoSimulado(t *testing.T) {
	lojas := []string{"Amazon", "MercadoLivre", "Kabum", "Magalu", "Shopee"}

	inicio := time.Now()
	r, err := BuscarPrimeiroPreco(context.Background(), lojas)
	duracao := time.Since(inicio)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	assertIgual(t, r.Loja, "MercadoLivre")
	assertIgual(t, r.Preco, 199.90)
	if duracao > 150*time.Millisecond { // a mais lenta leva 300ms: não pode ter esperado por ela
		t.Errorf("levou %v; deveria retornar logo após a mais rápida (~30ms)", duracao)
	}
}
