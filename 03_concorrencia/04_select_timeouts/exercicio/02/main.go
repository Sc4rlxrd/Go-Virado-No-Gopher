package main

import (
	"fmt"
	"time"
)

func MonitorarCotacoesComCancelamento(chDados <-chan float64, chDone <-chan struct{}) ([]float64, string) {
	precos := []float64{}

	for {
		select {
		case <-chDone:
			return precos, "cancelado_pelo_usuario"
		case preco, ok := <-chDados:
			if !ok {
				return precos, "canal_finalizado"
			}

			precos = append(precos, preco)

		case <-time.After(50 * time.Millisecond):
			fmt.Println("⚠️ Timeout")
		}

	}
}

func main() {
	fmt.Println(" --- Ex 02 --- ")

	dados := make(chan float64)
	done := make(chan struct{})

	go func() {
		for _, p := range []float64{199.90, 205.50, 189.00} {
			dados <- p
		}
		close(dados)
	}()

	precos, status := MonitorarCotacoesComCancelamento(dados, done)
	fmt.Printf("preços=%v status=%s\n", precos, status)

	// usando o timeout
	timeout := make(chan float64)
	go func() {
		for _, p := range []float64{222.90, 205.50, 189.00, 245.55, 3544.45} {
			time.Sleep(60 * time.Millisecond)
			timeout <- p
		}
		close(timeout)
	}()

	precos, status = MonitorarCotacoesComCancelamento(timeout, done)
	fmt.Printf("Usando o timeout -> preços=%v status=%s\n", precos, status)

}
