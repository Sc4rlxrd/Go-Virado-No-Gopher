package main

import "fmt"

func main() {

	valueMax := 50

	fmt.Println(" --- Exibindo os 25 primeiros números ---")
	for i := 0; i <= valueMax; i++ {
		if i == 25 {
			fmt.Printf("Interrompendo a execução no número: %d\n", i)
			break
		}

	}
	fmt.Println(" --- Fim da Lista ---")
}
