package main

import (
	"context"
	"fmt"
	"time"
)

type ResultadoScaping struct {
	Loja  string
	Preco float64
	Erro  error
}

// Função de Coleta que respeita o cancelamento/timeout do Context
func buscarPrecoComContext(ctx context.Context, loja string, preco float64, tempoSimulado time.Duration) ResultadoScaping {

	chRes := make(chan ResultadoScaping, 1)

	// Dispara o trabalho simulado em goroutine separada
	go func() {
		time.Sleep(tempoSimulado)
		chRes <- ResultadoScaping{
			Loja:  loja,
			Preco: preco,
		}
	}()

	// Select monitorando o canal de resultados E o canal de cancelamento do contexto
	select {
	case <-ctx.Done():
		// Se o tempo estourar ou o contexto for cancelado externamente:
		return ResultadoScaping{
			Loja: loja,
			Erro: fmt.Errorf("operação cancelada pelo contexto: %w", ctx.Err()),
		}
	case res := <-chRes:
		// Se o resultado chegar antes do cancelamento:
		return res
	}
}

func main() {
	fmt.Println("=== 🧭 Aula 07: context.Context — Cancelamento e Timeouts ===")

	// 1. Criando um contexto com Timeout de 100 milissegundos
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // Boas Práticas: Sempre garanta que o cancel() seja chamado via defer

	fmt.Println("🔍 Teste 1: Loja Rápida (deve responder a tempo")
	res1 := buscarPrecoComContext(ctxTimeout, "Amazon", 200.0, 30*time.Millisecond)
	if res1.Erro != nil {
		fmt.Println("Error: ", res1.Erro)
	} else {
		fmt.Printf(" ✅ [%s] Preço obtido: R$ %.2f\n", res1.Loja, res1.Preco)
	}

	fmt.Println("🔍 Teste 2: Loja Lenta (deve ser abortada pelo timeout do contexto)")
	res2 := buscarPrecoComContext(ctxTimeout, "Magalu", 194.0, 200*time.Millisecond)
	if res2.Erro != nil {
		fmt.Println(" ⚠️ Cancelado com sucesso:", res2.Erro)
	} else {
		fmt.Printf(" ✅ [%s] Preço obtido: R$ %.2f\n", res2.Loja, res2.Preco)
	}
}
