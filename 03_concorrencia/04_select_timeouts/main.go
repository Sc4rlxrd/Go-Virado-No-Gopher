package main

import (
	"fmt"
	"time"
)

type CotacaoRapida struct {
	Loja  string
	Preco float64
}

// Simula busca concorrente em duas fontes distintas de preço
func buscarNaLojaA(ch chan<- CotacaoRapida) {
	time.Sleep(80 * time.Millisecond) // Loja mais lenta
	ch <- CotacaoRapida{Loja: "A", Preco: 150.00}
}

func buscarNaLojaB(ch chan<- CotacaoRapida) {
	time.Sleep(30 * time.Millisecond)
	ch <- CotacaoRapida{Loja: "B", Preco: 145.00}
}

func buscarNaLojaC(ch chan<- CotacaoRapida) {
	time.Sleep(25 * time.Millisecond)
	ch <- CotacaoRapida{Loja: "C", Preco: 140.00}

}

func main() {
	fmt.Println("=== ⏱️ Aula 04: select, Timeouts e Operações Não-Bloqueantes ===")

	chA := make(chan CotacaoRapida, 1)
	chB := make(chan CotacaoRapida, 1)
	chC := make(chan CotacaoRapida, 1)

	go buscarNaLojaA(chA)
	go buscarNaLojaB(chB)
	go buscarNaLojaC(chC)

	// 1. Multiplexação: Pega a primeira cotação que responder (Race Pattern)
	select {
	case resultado := <-chA:
		fmt.Printf("🏆 [%s] Venceu a corrida! Preço: R$ %.2f\n", resultado.Loja, resultado.Preco)
	case resultado := <-chB:
		fmt.Printf("🏆 [%s] Venceu a corrida! Preço: R$ %.2f\n", resultado.Loja, resultado.Preco)
	case resultado := <-chC:
		fmt.Printf("🏆 [%s] Venceu a corrida! Preço: R$ %.2f\n", resultado.Loja, resultado.Preco)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("⚠️ Timeout geral: Nenhuma loja respondeu em tempo hábil!")
	}

	// 2. Exemplo de leitura não-bloqueante com 'default'
	chSinal := make(chan string, 1)
	select {
	case msg := <-chSinal:
		fmt.Println("Sinal recebido:", msg)
	default:
		fmt.Println("Nenhum sinal no canal. Continuando execução normal sem travar...")
	}
}
