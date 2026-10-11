package main

import "fmt"

func main() {

	fmt.Println(" --- Números Pares de 0 a 100 ---")

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Número Par: %d\n", i)
		}
	}
	fmt.Println(" --- Fim da Lista ---")
}
