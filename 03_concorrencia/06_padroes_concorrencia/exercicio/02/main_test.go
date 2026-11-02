package main

import (
	"fmt"
	"slices"
	"testing"
)

func assertIgual[T comparable](t *testing.T, esperado, atual T) {
	t.Helper()

	if esperado != atual {
		t.Errorf("esperado: %v | atual: %v", esperado, atual)
	}
}

func criarProduto(loja, nome string, preco float64, disponivel bool) Produto {
	return Produto{
		Loja:       loja,
		Nome:       nome,
		Preco:      preco,
		Disponivel: disponivel,
	}
}

// Cria um canal com a quantidade de produtos definida.
func criarCanalProdutos(loja string, quantidade int) <-chan Produto {
	canal := make(chan Produto)

	go func() {
		defer close(canal)

		for i := 0; i < quantidade; i++ {
			canal <- criarProduto(
				loja,
				fmt.Sprintf("Produto %d", i+1),
				float64((i+1)*10),
				i%2 == 0,
			)
		}
	}()

	return canal
}

// Cria a quantidade de canais/lojas e retorna
// quantos produtos devem ser recebidos no total.
func criarLojas(quantidadeLojas, produtosPorLoja int) ([]<-chan Produto, int) {
	canais := make([]<-chan Produto, 0, quantidadeLojas)

	for i := 0; i < quantidadeLojas; i++ {
		loja := fmt.Sprintf("Loja %d", i+1)

		canais = append(
			canais,
			criarCanalProdutos(loja, produtosPorLoja),
		)
	}

	totalProdutos := quantidadeLojas * produtosPorLoja

	return canais, totalProdutos
}

func TestFanIn(t *testing.T) {
	testes := []struct {
		nome            string
		quantidadeLojas int
		produtosPorLoja int
	}{
		{"5 lojas e 3 produtos por loja", 5, 3},
		{"10 lojas e 5 produtos por loja", 10, 5},
		{"varias lojas e 1 produto por loja", 8, 1},
		{"1 loja processa varios produtos", 1, 10},
		{"1 loja e 1 produto", 1, 1},
		{"mais lojas que produtos por loja", 10, 2},
		{"1 loja vazia", 1, 0},
		{"varias lojas vazias", 5, 0},
		{"nenhuma loja", 0, 3},
		{"nenhuma loja e nenhum produto", 0, 0},
		{"100 lojas e 10 produtos por loja", 100, 10},
		{"100 lojas e 100 produtos por loja", 100, 100},
		{"1000 lojas e 1 produto por loja", 1000, 1},
		{"9000 lojas e 1 produto por loja", 9000, 1},
		{"10000 lojas e 1 produto por loja", 10000, 1},
	}

	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			canais, totalEsperado := criarLojas(
				teste.quantidadeLojas,
				teste.produtosPorLoja,
			)

			resultado := FanIn(canais...)

			var atual []Produto

			for produto := range resultado {
				atual = append(atual, produto)
			}

			assertIgual(t, totalEsperado, len(atual))

			// Confere também se todos os produtos esperados
			// foram recebidos, independentemente da ordem.
			esperado := make([]Produto, 0, totalEsperado)

			for loja := 0; loja < teste.quantidadeLojas; loja++ {
				nomeLoja := fmt.Sprintf("Loja %d", loja+1)

				for produto := 0; produto < teste.produtosPorLoja; produto++ {
					esperado = append(esperado, criarProduto(
						nomeLoja,
						fmt.Sprintf("Produto %d", produto+1),
						float64((produto+1)*10),
						produto%2 == 0,
					))
				}
			}

			ordenarProdutos(atual)
			ordenarProdutos(esperado)

			for i := range esperado {
				assertIgual(t, esperado[i], atual[i])
			}
		})
	}
}

func ordenarProdutos(produtos []Produto) {
	slices.SortFunc(produtos, func(a, b Produto) int {
		if a.Loja < b.Loja {
			return -1
		}
		if a.Loja > b.Loja {
			return 1
		}

		if a.Nome < b.Nome {
			return -1
		}
		if a.Nome > b.Nome {
			return 1
		}

		if a.Preco < b.Preco {
			return -1
		}
		if a.Preco > b.Preco {
			return 1
		}

		if a.Disponivel == b.Disponivel {
			return 0
		}
		if !a.Disponivel {
			return -1
		}

		return 1
	})
}
