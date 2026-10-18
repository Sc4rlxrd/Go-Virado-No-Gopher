package main

import (
	"fmt"
	"testing"
)

func TestEnfileirar_Inteiros(t *testing.T) {

	var inteiros Fila[int]
	for i := 0; i <= 5; i++ {
		inteiros.Enfileirar(i + 20)
	}

	resultado := len(inteiros.itens)
	if len(inteiros.itens) == 6 {
		fmt.Println("valor esperado do len 6 fornecido: ", resultado)
	}

}


func TestDesenfileirar_Inteiros(t *testing.T) {

	var inteiros Fila[int]
    for i := 0; i <= 3; i++ {
		valor, ok := inteiros.Desenfileirar()
		fmt.Println("_", valor, "_", ok)
	}

	resultado := len(inteiros.itens)
	if len(inteiros.itens) == 3 {
		fmt.Println("valor esperado do len 3 fornecido: ", resultado)
	}

}

func TestEnfileirar_Strings(t *testing.T) {

	var frutas Fila[string]
		for i := 0; i <= 5; i++ {
		frutas.Enfileirar("")
	}

	resultado := len(frutas.itens)
	if len(frutas.itens) == 6 {
		fmt.Println("valor esperado do len 6 fornecido: ", resultado)
	}

}

func TestDesenfileirar_Strings(t *testing.T) {

	var frutas Fila[string]
    for i := 0; i <= 3; i++ {
		valor,ok:= frutas.Desenfileirar()
		fmt.Println("_", valor, "_", ok)
	}

	resultado := len(frutas.itens)
	if len(frutas.itens) == 3 {
		fmt.Println("valor esperado do len 3 fornecido: ", resultado)
	}

}