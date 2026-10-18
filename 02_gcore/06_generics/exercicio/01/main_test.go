package main

import(
	"testing"
	"fmt"
)

func TestContem_Inteiros_Sucesso(t *testing.T) {

	numeros  := []int{10, 20, 30}
	resultado :=  Contem(numeros, 20)
	esperado:= true

	if resultado!= esperado{
		fmt.Println("O valor esperado e true e for fornecido: ", esperado)
	}

	
}


func TestContem_Strings_Sucesso(t *testing.T) {

	frutas := []string{"Banana", "Maça", "Uva", "banana", "Laranja"}
	resultado :=  Contem(frutas, "Banana")
	esperado:= true

	if resultado!= esperado{
		fmt.Println("O valor esperado e true e for fornecido: ", esperado)
	}

	
}

func TestContem_Strings_Falha(t *testing.T) {

	frutas := []string{"Banana", "Maça", "Uva", "banana", "Laranja"}
	resultado :=  Contem(frutas, "Abacaxi")
	esperado:= false

	if resultado!= esperado{
		fmt.Println("O valor esperado e true e for fornecido: ", esperado)
	}

	
}
func TestContem_Inteiros_Falha(t *testing.T) {

	numeros  := []int{10, 20, 30}
	resultado :=  Contem(numeros, 45)
	esperado:= false

	if resultado!= esperado{
		fmt.Println("O valor esperado e true e for fornecido: ", esperado)
	}

	
}
