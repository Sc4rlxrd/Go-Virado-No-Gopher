package collector

import (
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/domain"
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/storage"
	"context"
	"fmt"
	"sync"
	"time"
)

type EngineColeta struct {
	cache        *storage.CacheCotacoes
	numWorkers   int
	timeoutGeral time.Duration
}

func NovaEngineColeta(cache *storage.CacheCotacoes, numWorkers int, timeout time.Duration) *EngineColeta {
	return &EngineColeta{
		cache:        cache,
		numWorkers:   numWorkers,
		timeoutGeral: timeout,
	}
}

// Worker que processa tarefas do canal de jobs com suporte a context
func (e *EngineColeta) worker(ctx context.Context, id int, jobs <-chan domain.SolicitacaoBusca, result chan<- domain.ResultadoColeta, wg *sync.WaitGroup) {

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			for _, loja := range job.Lojas {
				select {
				case <-ctx.Done():
					result <- domain.ResultadoColeta{
						Erro: fmt.Errorf("worker %d cancelado via contexto para loja %s: %w", id, loja, ctx.Err()),
					}
					return

				case <-time.After(10 * time.Millisecond): // Simula latência de rede/scraping
					cotacao := domain.ItemCotacao{
						ID:       id,
						Produto:  job.Produto,
						Loja:     loja,
						Preco:    199.90 + float64(job.ID),
						CriadoEm: time.Now(),
					}
					e.cache.Salvar(cotacao)
					result <- domain.ResultadoColeta{
						Cotacao: cotacao,
					}
				}
			}
		}
	}
}

// ExecutarLote orquestra o Worker Pool e consolida (Fan-In) os resultados
func (e *EngineColeta) ExecutarLote(ctxPai context.Context, solicitacoes []domain.SolicitacaoBusca) ([]domain.ItemCotacao, []error) {
	if len(solicitacoes) == 0 {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(ctxPai, e.timeoutGeral)
	defer cancel()

	jobs := make(chan domain.SolicitacaoBusca, len(solicitacoes))
	results := make(chan domain.ResultadoColeta, len(solicitacoes)*5)

	var wg sync.WaitGroup

	// 1. Inicializa o Worker Pool
	for w := 1; w <= e.numWorkers; w++ {
		wg.Add(1)
		go e.worker(ctx, w, jobs, results, &wg)
	}
	// 2. Alimenta o canal de jobs
	for _, req := range solicitacoes {
		jobs <- req
	}
	close(jobs)

	// 3. Goroutine para fechar o canal de resultados após término dos workers
	go func() {
		wg.Wait()
		close(results)
	}()

	var cotacoes []domain.ItemCotacao
	var errs []error

	// 4. Consome os resultados até fechar o canal
	for res := range results {
		if res.Erro != nil {
			errs = append(errs, res.Erro)
			continue // sem cotacao vazia
		}
		cotacoes = append(cotacoes, res.Cotacao)

	}
	// O select do worker sorteia; isto garante que o chamador saiba que o lote foi interrompido.
	if err := ctx.Err(); err != nil {
		errs = append(errs, fmt.Errorf("lote interrompido: %w", err))
	}
	return cotacoes, errs
}
