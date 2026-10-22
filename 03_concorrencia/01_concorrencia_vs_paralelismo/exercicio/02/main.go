package main

import (
	"fmt"
	"runtime"
	"time"
)

func VerificarLoja(nomeLoja string) {
	fmt.Printf("[price-agent] Coletando preços na loja: %s\n", nomeLoja)
	time.Sleep(5 * time.Millisecond)
}

func main() {
	lojas := []string{"Amazon", "MercadoLivre", "Kabum", "Magalu", "Shopee"}

	fmt.Printf("Goroutines vivas : %d\n", runtime.NumGoroutine())
	for indice := range lojas {
		go VerificarLoja(lojas[indice])
	}

	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Goroutines vivas : %d\n", runtime.NumGoroutine())
}
