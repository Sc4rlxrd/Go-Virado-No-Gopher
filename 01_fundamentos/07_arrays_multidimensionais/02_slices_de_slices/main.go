package main

import "fmt"

func main() {
	// Slices de slices: cada linha pode ter um tamanho diferente.
	matrizVariavel := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("--- Slices de Slices (Linhas Irregulares) ---")

	for i, linha := range matrizVariavel {
		fmt.Printf("Linha %d (tamanho %d): %v\n", i, len(linha), linha)
	}

	// Alocação dinâmica de uma grade 2D via make.
	linhas, colunas := 3, 4

	grade := make([][]int, linhas)

	for i := range grade {
		grade[i] = make([]int, colunas)
		// Aloca cada linha individualmente.
	}

	fmt.Printf("Grade %dx%d alocada via make: %v\n", linhas, colunas, grade)
}
