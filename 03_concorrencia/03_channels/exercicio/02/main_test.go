package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func ordenarPorLoja(rs []ResultadoCotacao) {
	slices.SortFunc(rs, func(a, b ResultadoCotacao) int {
		return strings.Compare(a.Loja, b.Loja)
	})
}

func TestColetarTodasCotacoes(t *testing.T) {
	testes := []struct {
		nome         string
		lojas        []string
		esperado     []ResultadoCotacao // já em ordem alfabética por Loja
		errosEsperad int
	}{
		{
			nome:  "todas válidas",
			lojas: []string{"Magalu", "Amazon", "Kabum"},
			esperado: []ResultadoCotacao{
				{Loja: "Amazon", Preco: 199.90},
				{Loja: "Kabum", Preco: 199.90},
				{Loja: "Magalu", Preco: 199.90},
			},
			errosEsperad: 0,
		},
		{
			nome:  "uma inexistente no meio",
			lojas: []string{"Magalu", "LojaInexistente", "Amazon"},
			esperado: []ResultadoCotacao{
				{Loja: "Amazon", Preco: 199.90},
				{Loja: "LojaInexistente", Err: ErrLojaNaoEncontrada},
				{Loja: "Magalu", Preco: 199.90},
			},
			errosEsperad: 1,
		},
		{
			nome:  "todas inexistentes",
			lojas: []string{"LojaInexistente", "LojaInexistente"},
			esperado: []ResultadoCotacao{
				{Loja: "LojaInexistente", Err: ErrLojaNaoEncontrada},
				{Loja: "LojaInexistente", Err: ErrLojaNaoEncontrada},
			},
			errosEsperad: 2,
		},
		{
			nome:         "entrada vazia",
			lojas:        []string{},
			esperado:     []ResultadoCotacao{},
			errosEsperad: 0,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			obtido, erros := ColetarTodasCotacoes(tt.lojas)
			ordenarPorLoja(obtido)

			if erros != tt.errosEsperad {
				t.Errorf("erros = %d; esperado %d", erros, tt.errosEsperad)
			}
			if !slices.Equal(obtido, tt.esperado) {
				t.Errorf("obtido %+v\nesperado %+v", obtido, tt.esperado)
			}
		})
	}
}

// Estresse: muitas goroutines enviando ao mesmo tempo no mesmo canal.
func TestColetarTodasCotacoesEstresse(t *testing.T) {
	const n = 300
	lojas := make([]string, n)
	for i := range lojas {
		if i%3 == 0 {
			lojas[i] = "LojaInexistente"
		} else {
			lojas[i] = fmt.Sprintf("Loja-%d", i)
		}
	}

	resultados, erros := ColetarTodasCotacoes(lojas)

	if len(resultados) != n {
		t.Fatalf("len = %d; esperado %d", len(resultados), n)
	}
	if esperado := (n + 2) / 3; erros != esperado { // índices 0,3,6,... = 100
		t.Errorf("erros = %d; esperado %d", erros, esperado)
	}
}