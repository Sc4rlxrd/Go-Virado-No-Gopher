package main

import (
	"fmt"
	"sync"
	"time"
)

type ItemPreco struct {
	Produto string
	Preco   float64
}

// Simula a requisição de coleta do price-agent em uma loja
func coletarPrecoProduto(produto string) ItemPreco {
	time.Sleep(20 * time.Millisecond)
	return ItemPreco{
		Produto: produto,
		Preco:   199.99,
	}
}

func main() {

	fmt.Println("=== ⏳ Aula 02: Ciclo de Vida de Goroutines e sync.WaitGroup ===")

	produtos := []string{
		"Teclado Redragon",
		"Mouse vertical logi",
		"Monitor 27pol samsung 75hz",
	}

	// 1. Declarando o WaitGroup para sincronização perfeita
	var wg sync.WaitGroup

	fmt.Println("🚀 Iniciando varredura concorrente de preços...")

	for _, prod := range produtos {
		// Incrementa o contador do WaitGroup ANTES do disparo da goroutine
		wg.Add(1)

		// Disparando closure anônima passando 'prod' como parâmetro explícito
		go func(nomeProduto string) {
			// Garantimos que wg.Done() será executado ao sair da goroutine
			defer wg.Done()

			item := coletarPrecoProduto(nomeProduto)
			fmt.Printf("🛒 [price-agent] Coletado: %-25s | R$ %.2f\n", item.Produto, item.Preco)
		}(prod) // Passagem por valor para evitar a armadilha do loop!
	}
	fmt.Println("⏳ Aguardando a finalização de todos os coletores...")
	// Bloqueia até que todas as goroutines chamem wg.Done()
	wg.Wait()
	fmt.Println("✅ Todas as coletas foram concluídas com sucesso! O programa pode encerrar de forma segura.")
}
