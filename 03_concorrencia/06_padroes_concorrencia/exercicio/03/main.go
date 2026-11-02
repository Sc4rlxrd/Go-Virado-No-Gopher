package main

import "fmt"

type ItemBruto struct {
	ID            int
	Nome          string
	PrecoOriginal float64
}

type ItemProcessado struct {
	ID         int
	Nome       string
	PrecoFinal float64
	Valido     bool
}

func EstagioGerador(itens []ItemBruto) <-chan ItemBruto {
	canal := make(chan ItemBruto)

	go func() {
		defer close(canal)
		for _, item := range itens {
			canal <- item
		}
	}()

	return canal
}

func EstagioFiltro(in <-chan ItemBruto) <-chan ItemBruto {
	out := make(chan ItemBruto)

	go func() {
		defer close(out)
		for item := range in {
			if item.PrecoOriginal > 0 {
				out <- item
			}
		}
	}()

	return out
}

func EstagioCalcularDesconto(in <-chan ItemBruto, percentualDesconto float64) <-chan ItemProcessado {
	out := make(chan ItemProcessado)

	go func() {
		defer close(out)
		for item := range in {
			precoFinal := item.PrecoOriginal * (1 - percentualDesconto/100)
			out <- ItemProcessado{
				ID:         item.ID,
				Nome:       item.Nome,
				PrecoFinal: precoFinal,
				Valido:     true,
			}
		}
	}()

	return out
}

func main() {

	fmt.Println("Exercicio 3")
	dados := []ItemBruto{
		{ID: 1, Nome: "Produto A", PrecoOriginal: 100.0},
		{ID: 2, Nome: "Produto B", PrecoOriginal: 200.0},
		{ID: 3, Nome: "Produto C", PrecoOriginal: -50.0}, // Produto inválido
		{ID: 4, Nome: "Produto D", PrecoOriginal: 150.0},
		{ID: 5, Nome: "Produto E", PrecoOriginal: 0.0}, // Produto inválido
	}
	out := EstagioCalcularDesconto(EstagioFiltro(EstagioGerador(dados)), 10.0)
	fmt.Println("Produtos processados:..")
	for item := range out {
		fmt.Printf("Item: %+v\nValor Final: %.2f\n", item.Nome, item.PrecoFinal)
		fmt.Println("---")
	}
}
