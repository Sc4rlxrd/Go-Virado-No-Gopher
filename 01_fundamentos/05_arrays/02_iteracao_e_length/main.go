package main

import "fmt"

func main(){

	names := [6]string{"Gopher", "Fedora", "OracleCloud","AWS", "Azure","Scarlxrd"}

	fmt.Println("--- Tamanho e Iteração com len() ---")
	fmt.Printf("Tamanho do array (len): %d\n", len(names))


	for i := 0; i < len(names); i++ {
		fmt.Printf("Elemento no índice: [%d] = %s\n", i, names[i])
	}

	// DEMONSTRAÇÃO DE CÓPIA DE VALOR EM GO

	fmt.Println("--- Cópia de Valor em Arrays ---")

	copyNames := names // Cópia de valor do array 'names' para 'copyNames'
	copyNames[5] = "Oracle Cloud" // Modificando o array copiado

	fmt.Printf("Array original: %v\n", names)
	fmt.Printf("Array copiado: %v\n", copyNames)

}
