package main

import "fmt"

func Contem[T comparable](slice []T, elemento T) bool {

	for _, item := range slice {
		if item == elemento {
			return true

		}
	}
	return false
}

func main() {

	fmt.Println(" === Ex 01 ===")

	numeros := []int{10, 20, 30}


	fmt.Println("Valor 20 se encontra no Slice: ", Contem(numeros, 20))   // true
	fmt.Println("O valor 40 se encontra no Slice: ", Contem(numeros, 40)) // false

	frutas := []string{"Banana", "Maça", "Uva", "banana", "Laranja"}
	fmt.Println("A fruta Banana se encontra no Slice: ", Contem(frutas, "Banana"))  // true
	fmt.Println("A fruta Abacaxi se encontra no Slice: ", Contem(frutas, "Abacaxi"))  // false

	fmt.Println("Busca em slice vazio:", Contem([]int{}, 10))

	

}
