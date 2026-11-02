package main

import (
	"testing"
)

func assertIgual[T comparable](t *testing.T, esperado, atual T) {
	t.Helper()

	if esperado != atual {
		t.Errorf("esperado: %v | atual: %v", esperado, atual)
	}
}

func TestPipelineTresEstagios(t *testing.T) {
	dados := []ItemBruto{
		{ID: 1, Nome: "Produto A", PrecoOriginal: 100.0},
		{ID: 2, Nome: "Produto B", PrecoOriginal: 200.0},
		{ID: 3, Nome: "Produto C", PrecoOriginal: -50.0},
		{ID: 4, Nome: "Produto D", PrecoOriginal: 150.0},
		{ID: 5, Nome: "Produto E", PrecoOriginal: 0.0},
	}

	esperado := []ItemProcessado{
		{ID: 1, Nome: "Produto A", PrecoFinal: 90.0, Valido: true},
		{ID: 2, Nome: "Produto B", PrecoFinal: 180.0, Valido: true},
		{ID: 4, Nome: "Produto D", PrecoFinal: 135.0, Valido: true},
	}

	out := EstagioCalcularDesconto(
		EstagioFiltro(
			EstagioGerador(dados),
		),
		10.0,
	)

	var atual []ItemProcessado

	for item := range out {
		atual = append(atual, item)
	}

	assertIgual(t, len(esperado), len(atual))

	for i := range esperado {
		assertIgual(t, esperado[i], atual[i])
	}
}
