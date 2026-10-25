package main

import (
	"fmt"
	"sync"
)

// DuplicarCanal lê de chIn e envia cada mensagem para chOut1 e chOut2.
// Quando chIn fecha, fecha os dois canais de saída.
func DuplicarCanal(chIn <-chan string, chOut1, chOut2 chan<- string) {
	defer close(chOut1)
	defer close(chOut2)

	for msg := range chIn {
		chOut1 <- msg
		chOut2 <- msg
	}
}

func main() {
	fmt.Println(" --- Ex 03 --- ")

	mensagens := []string{"deploy iniciado", "erro 500 no checkout", "deploy finalizado"}

	entrada := make(chan string)
	disco := make(chan string, len(mensagens))
	alertas := make(chan string, len(mensagens))

	// Produtor: quem envia em entrada é quem fecha
	go func() {
		defer close(entrada)
		for _, m := range mensagens {
			entrada <- m
		}
	}()

	go DuplicarCanal(entrada, disco, alertas)

	// Dois consumidores independentes, cada um com o seu destino
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for msg := range disco {
			fmt.Println("[DISCO]  ", msg)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range alertas {
			fmt.Println("[ALERTA] ", msg)
		}
	}()

	wg.Wait()
}
