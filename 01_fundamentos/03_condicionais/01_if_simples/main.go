package main

import "fmt"

func main() {

	// variável de controle
	age := 11
	isAuthorized := age >= 18

	fmt.Println("--- Validação de Autorização ---")

	// Em Go, NÃO se usam parênteses em volta da condição.
	// // As chaves {} são OBRIGATÓRIAS, mesmo para uma única linha de código.
	if isAuthorized {
		fmt.Println("Acesso liberado: Usuário é maior de idade.")
	}

	if !isAuthorized {
		fmt.Println("Acesso negado: Usuário é menor de idade.")
	}

	fmt.Println("Processamento finalizado.")
}
