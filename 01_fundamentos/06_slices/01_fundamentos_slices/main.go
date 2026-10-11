package main

import "fmt"

func main() {

	// Slices não possuem tamanho definido no tipo ([]T) // São visões dinâmicas sobre um array subjacente.
	var sliceSemTamanho []int
	fmt.Printf("Slice Nil: %v | len: %d | cap: %d | é nil? %t\n", sliceSemTamanho, len(sliceSemTamanho), cap(sliceSemTamanho), sliceSemTamanho == nil)

	// 1. Inicialização via Composite Literal
	frutas := []string{"Maçã", "Banana", "Laranja"}
	fmt.Printf("Slice Frutas: %v | len: %d | cap: %d\n", frutas, len(frutas), cap(frutas))

	//  Crescimento Dinâmico com append()
	frutas = append(frutas, "Uva", "Manga")
	fmt.Printf("Após append: %v | len: %d | cap: %d\n", frutas, len(frutas), cap(frutas))

	//  Alocação via make(tipo, len, cap)
	//  make aloca o array interno pré-dimensionado
	buffer := make([]byte, 5, 10) // len = 5, cap = 10
	fmt.Printf("Slice via make: %v | len: %d | cap: %d\n", buffer, len(buffer), cap(buffer))

}
