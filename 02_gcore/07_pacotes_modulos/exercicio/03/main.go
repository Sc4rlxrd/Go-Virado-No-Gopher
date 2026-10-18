package main

import (
	hash "Go-Virado-No-Gopher/02_gcore/07_pacotes_modulos/exercicio/03/internal/seguranca"
	"fmt"
)

func main() {

	fmt.Println("--- Ex 03 ---")
	fmt.Println("Gerador de hash: ", hash.GerarHash())
	fmt.Println()
	for i := 0; i <= 4; i++ {
		fmt.Printf("Hash %d: %s\n", i, hash.GerarHash())
	}

}
