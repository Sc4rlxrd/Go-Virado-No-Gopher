package main

import (
	"fmt"
	"sync"
)

type ItemColeta struct {
	ID          int
	NomeProduto string
	Processado  bool
	Tag         string
}

func EnriquecerItem(item ItemColeta) ItemColeta {
	return ItemColeta{
		ID:          item.ID,
		NomeProduto: item.NomeProduto,
		Processado:  true,
		Tag:         "COLETA_PENDENTE",
	}
}

func ProcessarLoteConcorrente(itens []ItemColeta) []ItemColeta {

	if len(itens) == 0 { // cobre nil e vazio
		return []ItemColeta{}
	}

	var wg sync.WaitGroup

	resultado := make([]ItemColeta, len(itens))

	for i, item := range itens {
		wg.Add(1)
		go func(i int, it ItemColeta) {
			defer wg.Done()
			resultado[i] = EnriquecerItem(it) // cada goroutine escreve só no seu índice

		}(i, item)

	}
	wg.Wait()
	return resultado
}

func main() {
	fmt.Println("--- Ex 03 ---")

	produto := EnriquecerItem(ItemColeta{
		ID:          1,
		NomeProduto: "Teclado Redragon",
	})

	fmt.Printf("Produto:%v\nProcessado:%v\nColeta:%v\n", produto.NomeProduto, produto.Processado, produto.Tag)

	produtos := []ItemColeta{
		{
			ID:          1,
			NomeProduto: "Teclado Mancer ",
		},
		{
			ID:          2,
			NomeProduto: "Mouse Gamer Logitech",
		},
		{
			ID:          3,
			NomeProduto: "Headset HyperX",
		},
		{
			ID:          4,
			NomeProduto: "Monitor UltraWide LG",
		},
		{
			ID:          5,
			NomeProduto: "Cadeira Gamer ThunderX3",
		},
	}

	fmt.Println("---")
	for _, item := range ProcessarLoteConcorrente(produtos) {
		fmt.Printf("Produto:%v\nProcessado:%v\nColeta:%v\n---\n", item.NomeProduto, item.Processado, item.Tag)
	}
}
