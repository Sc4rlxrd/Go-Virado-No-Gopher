package main

import (
	"fmt"
	"sync"
	"time"
)

type Produto struct {
	Loja  string
	Nome  string
	Preco float64
	Disponivel bool
}

func FanIn[T any](canais ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(canais))

	for _, c := range canais {
		go func(c <-chan T) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func worker(id int, jobs <-chan Produto, wg *sync.WaitGroup) {
	defer wg.Done()
	for produto := range jobs {
		time.Sleep(120 * time.Millisecond)
		fmt.Printf(
			"[Worker %d] Loja: %-12s | Produto processado: %s (R$ %.2f) - Disponível: %t\n",
			id,
			produto.Loja,
			produto.Nome,
			produto.Preco,
			produto.Disponivel,
		)
	}
}

func main() {
	fmt.Println("Exercicio 02")

	amazon := make(chan Produto)
	go func() {
		amazon <- Produto{Loja: "Amazon", Nome: "Livro 1", Preco: 50.0, Disponivel: true}
		amazon <- Produto{Loja: "Amazon", Nome: "Livro 2", Preco: 60.0, Disponivel: true}
		amazon <- Produto{Loja: "Amazon", Nome: "Livro 3", Preco: 70.0, Disponivel: true}
		amazon <- Produto{Loja: "Amazon", Nome: "Lenovo Ideapad 3 com Ryzen 5 5500U, 8GB RAM, 256GB SSD, Windows 11 - Cinza", Preco: 3000.0, Disponivel: true}
		defer close(amazon)
	}()

	magalu := make(chan Produto)
	go func() {
		magalu <- Produto{Loja: "Magalu", Nome: "Livro 3", Preco: 80.0, Disponivel: true}
		magalu <- Produto{Loja: "Magalu", Nome: "Livro 4", Preco: 90.0, Disponivel: true}
		magalu <- Produto{Loja: "Magalu", Nome: "Notebook Gamer Acer Nitro 5 AN515-45-R92M, AMD Ryzen 7 5800H, 16GB RAM, 512GB SSD, NVIDIA GeForce RTX 3060, Windows 11 - Preto", Preco: 5000.0, Disponivel: true}
		defer close(magalu)
	}()

	casasBahia := make(chan Produto)
	go func() {
		casasBahia <- Produto{Loja: "Casas Bahia", Nome: "Livro 5", Preco: 100.0, Disponivel: true}
		casasBahia <- Produto{Loja: "Casas Bahia", Nome: "Livro 6", Preco: 110.0, Disponivel: true}
		casasBahia <- Produto{Loja: "Casas Bahia", Nome: "Iphone 18 Pro Max, 1TB, Tela 6.7', Câmera Tripla, iOS 18 - Preto", Preco: 10000.0, Disponivel: true}
		defer close(casasBahia)
	}()

	boaDica := make(chan Produto)
	go func() {
		boaDica <- Produto{Loja: "BoaDica", Nome: "Notebook Lenovo ThinkPad T14 Gen 4 (21HESG6C00)", Preco: 12000.0, Disponivel: true}
		boaDica <- Produto{Loja: "BoaDica", Nome: "Smartphone F7 12GB/512GB Black", Preco: 2899.0, Disponivel: true}
		boaDica <- Produto{Loja: "BoaDica", Nome: "Apple iPhone 15 128GB Desbloqueado Preto", Preco: 2648.0, Disponivel: true}
		boaDica <- Produto{Loja: "BoaDica", Nome: "Xiaomi 14T Pro 512GB / 12GB Titan Gray", Preco: 3500.0, Disponivel: true}
		boaDica <- Produto{Loja: "BoaDica", Nome: "Notebook Acer Nitro V1", Preco: 4500.0, Disponivel: true}
		boaDica <- Produto{Loja: "BoaDica", Nome: "Computador Core i5-12400F / 16GB / RTX 3050 / 600W", Preco: 5500.0, Disponivel: true}
		boaDica <- Produto{Loja: "BoaDica", Nome: "Console Nintendo Switch Lite 32GB Cinza (HDH-001)", Preco: 1500.0, Disponivel: true}
		boaDica <- Produto{Loja: "BoaDica", Nome: "Notebook Lenovo ThinkPad L13 (20R4002UBR)", Preco: 3000.0, Disponivel: true}
		defer close(boaDica)
	}()


	teraByte := make(chan Produto)
	go func() {
		teraByte <- Produto{Loja: "TeraByte", Nome: "Notebook Gamer Acer Nitro 5 AN515-45-R92M, AMD Ryzen 7 5800H, 16GB RAM, 512GB SSD, NVIDIA GeForce RTX 3060, Windows 11 - Preto", Preco: 5000.0, Disponivel: true}
		teraByte <- Produto{Loja: "TeraByte", Nome: "Notebook Lenovo ThinkPad T14 Gen 4 (21HESG6C00)", Preco: 12000.0, Disponivel: true}
		teraByte <- Produto{Loja: "TeraByte", Nome: "Smartphone F7 12GB/512GB Black", Preco: 2899.0, Disponivel: true}
		teraByte <- Produto{Loja: "TeraByte", Nome: "Apple iPhone 15 128GB Desbloqueado Preto", Preco: 2648.0, Disponivel: true}
		teraByte <- Produto{Loja: "TeraByte", Nome: "Xiaomi 14T Pro 512GB / 12GB Titan Gray", Preco: 3500.0, Disponivel: true}
		teraByte <- Produto{Loja: "TeraByte", Nome: "Notebook Acer Nitro V1", Preco: 4500.0, Disponivel: true}
		teraByte <- Produto{Loja: "TeraByte", Nome: "Computador Core i5-12400F / 16GB / RTX 3050 / 600W", Preco: 5500.0, Disponivel: true}
		teraByte <- Produto{Loja: "TeraByte", Nome: "Console Nintendo Switch Lite 32GB Cinza (HDH-001)", Preco: 1500.0, Disponivel: false}
		teraByte <- Produto{Loja: "TeraByte", Nome: "Notebook Lenovo ThinkPad L13 (20R4002UBR)", Preco: 3000.0, Disponivel: false}
		defer close(teraByte)
	}()

	produtosAgrupados := FanIn(amazon, magalu, casasBahia, boaDica,teraByte)
	var wg sync.WaitGroup
	numWorkers := 10
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, produtosAgrupados, &wg)
	}

	wg.Wait()

}
