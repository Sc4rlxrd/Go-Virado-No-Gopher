package main

import "fmt"

func main() {

	// se não setar um tamanho, o compilador entende que é um SLICE, e não um ARRAY.
	// Slices são mais flexíveis que Arrays, pois podem crescer dinamicamente.
	// e para add novos elementos, podemos usar a função 'append()' para adicionar novos elementos ao slice.

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("--- Iteração Idiomática com 'range' ---")
	// numbers = append(numbers, 11, 12, 13, 14, 15)
	// O 'range' retorna dois valores a cada passo: (índice, valor)
	for i, num := range numbers {
		fmt.Printf("Elemento no índice: [%d] = %d\n", i, num)
	}
	fmt.Println("--- Descartando o índice com Blank Identifier (_) ---")
	sun := 0
	for _, num := range numbers {
		sun += num
	}
	fmt.Printf("Soma total dos elementos: %d\n", sun)
}
