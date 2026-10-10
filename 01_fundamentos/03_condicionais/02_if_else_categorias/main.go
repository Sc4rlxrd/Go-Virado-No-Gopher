package main

import "fmt"

func main() {

	// Idade para caterização em torneio esportivo
	age := 5
	var category string
	
	// Categorias:
	// Infantil: 5 a 10 anos
	// Juvenil: 11 a 15 anos
	// Adulto: 16 a mais

	fmt.Println("--- Categorização do Atleta ---")

	// Em Go, a chave de fechamento '}' e a palavra 'else' 
	// // DEVEM estar na mesma linha: "} else {"

	if age >= 5 && age <= 10 {
		category = "Infantil"
	} else if age >= 11 && age <= 15 {
		category = "Juvenil"
	} else if age >= 16 {
		category = "Adulto"
	} else {
		category = "Idade não permitida para o torneio."
	}

	fmt.Printf("Idade: %d anos -> Classificação: %s\n", age, category)

}
