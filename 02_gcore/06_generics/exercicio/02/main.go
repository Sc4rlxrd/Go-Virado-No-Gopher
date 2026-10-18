package main

import "fmt"

type Fila[T any] struct {
	itens []T
}

func (f *Fila[T]) Enfileirar(item T) {
	f.itens = append(f.itens, item)
}

func (f *Fila[T]) Desenfileirar() (T, bool) {

	if len(f.itens) == 0 {
		var zero T
		return zero, false
	}
	primeiro := f.itens[0]
	f.itens = f.itens[1:]
	return primeiro, true
}
func main() {

	fmt.Println(" === Ex 02 ===")
	var inteiros Fila[int]
	for i := 0; i <= 5; i++ {
		inteiros.Enfileirar(i + 20)
	}
	fmt.Println("Valores da fila: ", inteiros)
	for i := 0; i <= 3; i++ {
		valor, ok := inteiros.Desenfileirar()
		fmt.Println("Valor: ", valor, "| Sucesso: ", ok)
	}

	fmt.Println("Valores da fila:", inteiros)

	var frutas Fila[string]
	frutas.Enfileirar("Banana")
	frutas.Enfileirar("Maçã")
	frutas.Enfileirar("Uva")
	frutas.Enfileirar("Laranja")

	fmt.Println("Valores da fila:", frutas)
	valor,ok:= frutas.Desenfileirar()
	fmt.Println("Primeira fruta:", valor, "| Sucesso: ", ok )
	fmt.Println("Valores da fila:", frutas)
    fmt.Println("Remoção em fila vazia:", valor, ok)




}
