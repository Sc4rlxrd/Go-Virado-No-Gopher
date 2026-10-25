package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

var ErrLojaNaoEncontrada = errors.New("loja nao encontrada")
const precoPadrao = 199.90

type ResultadoCotacao struct {
	Loja  string
	Preco float64
	Err   error
}

func BuscarCotacao(loja string, chOut chan<- ResultadoCotacao) {

	if loja == "LojaInexistente" {
		chOut <- ResultadoCotacao{
			Loja: loja,
			Err:  ErrLojaNaoEncontrada}
		return

	}

	chOut <- ResultadoCotacao{Loja: loja, Preco: precoPadrao, Err: nil}

}

func ColetarTodasCotacoes(lojas []string) ([]ResultadoCotacao, int) {

	ch := make(chan ResultadoCotacao, len(lojas))

	for _, loja := range lojas {
		time.Sleep(5 * time.Millisecond)
		go BuscarCotacao(loja, ch)
	}

	resultados := make([]ResultadoCotacao, 0, len(lojas))
	erros := 0
	for range lojas { // lê exatamente len(lojas) resultados
		r := <-ch
		if r.Err != nil {
			erros++
		}
		resultados = append(resultados, r)
	}
	return resultados, erros
}

func main() {

	fmt.Println(" --- Ex 02 --- ")

	lojas := []string{"Amazon", "MercadoLivre", "LojaInexistente", "Kabum", "Magalu","Boa Dica","Terabyte"}

	resultados, totalErros := ColetarTodasCotacoes(lojas)

	// A ordem de chegada muda a cada execução; ordeno só para a saída ficar estável.
	sort.Slice(resultados, func(i, j int) bool {
		return resultados[i].Loja < resultados[j].Loja
	})

	for _, r := range resultados {
		if r.Err != nil {
			fmt.Printf("❌ %-16s erro: %v\n", r.Loja, r.Err)
			continue
		}
		fmt.Printf("✅ %-16s R$ %.2f\n", r.Loja, r.Preco)
	}

	fmt.Printf("\nTotal de lojas: %d | Com erro: %d\n", len(resultados), totalErros)
}
