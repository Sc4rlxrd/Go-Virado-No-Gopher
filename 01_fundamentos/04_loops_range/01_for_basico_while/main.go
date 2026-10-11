package main

import "fmt"

func main() {

	//  For estilo Clássico (Init; Condition; Post)
	fmt.Println("---- For Classic ----")

	for i := 0; i <= 10; i++ {
		fmt.Printf("Contador C-style: %d\n", i)
	}

	println()
	//  For estilo 'while' (Apenas a Condição)
	// // Em Go, 'while' não existe. Basta omitir o init e o post.

	fmt.Println("---- For estilo While ----")

	count := 0
	for count <= 10 {
		fmt.Printf("Contador estilo While: %d\n", count)
		count++
	}

	println()

	//  For Infinito (estilo 'while(true)')

	fmt.Println("---- For Infinito com Parada ----")
	steps := 0

	for {
		if steps >= 3 {
			fmt.Println("Atingiu o limite. Saindo do loop infinito....")
			break
		}
		fmt.Printf("Passo atual:  %d\n", steps)
		steps++
	}
}
