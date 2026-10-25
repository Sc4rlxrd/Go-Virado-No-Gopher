package main

import (
	"fmt"
	"time"
)

type Cotacao struct {
	Loja  string
	Preco float64
}

// Produtor: Envia cotações para o canal (chan<- indica canal de apenas envio)
func coletarPrecoEmLote(lojas []string, canalOut chan<- Cotacao) {

	for _, loja := range lojas {
		// simula a coleta com Selenium/HTTP
		time.Sleep(99 * time.Millisecond)
		canalOut <- Cotacao{Loja: loja, Preco: 299.99}
	}
	// REGRA IMPORTANTE: O produtor fecha o canal após terminar os envios!
	defer close(canalOut)

}

func main() {
	fmt.Println("=== 📡 Aula 03: Channels e Comunicação Concorrente ===")

	lojas := []string{"Amazon", "MercadoLivre", "Kabum", "Magalu", "Boa Dica"}

	// 1. Criando um Buffered Channel com capacidade para o número total de lojas
	canalCotacoes := make(chan Cotacao, len(lojas))

	// 2. Disparando a coleta em uma goroutine produtora
	go coletarPrecoEmLote(lojas, canalCotacoes)

	// 3. Consumindo os dados no fluxo principal com 'for range'
	fmt.Println("🛒 Recebendo cotações do price-agent via channel:")
	for cotacao := range canalCotacoes {
		fmt.Printf(" • [%-12s] R$ %.2f\n", cotacao.Loja, cotacao.Preco)
	}

	fmt.Println("✅ Canal fechado e todas as cotações foram processadas!")
}
