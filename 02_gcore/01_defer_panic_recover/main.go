package main

import (
	"fmt"
)

func processar() {
	defer fmt.Println("Executando limpeza")

	fmt.Println("Iniciando processamento")
	panic("falha inesperada")

	// Esta linha não será executada.
}

func operacaoSegura() {
	defer func() {
		if valor := recover(); valor != nil {
			fmt.Println("Pânico recuperado:", valor)
		}
	}()

	fmt.Println("Iniciando operação")
	panic("falha inesperada")

	// Esta linha não será executada.
}

func main() {
	defer fmt.Println("Primeiro defer")
	defer fmt.Println("Segundo defer")
	defer fmt.Println("Terceiro defer")

	fmt.Println("----- Uso do panic -----")
	// processar()
	fmt.Println("Fim do programa")

	fmt.Println("----- Uso do recover ----")
	operacaoSegura()
	fmt.Println("O programa continuou")
}
