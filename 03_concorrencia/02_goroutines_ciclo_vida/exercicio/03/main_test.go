package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestEnriquecerItem(t *testing.T) {
	original := ItemColeta{ID: 7, NomeProduto: "Teclado"}

	obtido := EnriquecerItem(original)

	esperado := ItemColeta{ID: 7, NomeProduto: "Teclado", Processado: true, Tag: "COLETA_PENDENTE"}
	if obtido != esperado {
		t.Errorf("EnriquecerItem() = %+v; esperado %+v", obtido, esperado)
	}
	if original.Processado || original.Tag != "" {
		t.Errorf("a entrada foi alterada: %+v", original)
	}
}

func TestProcessarLoteConcorrenteVazio(t *testing.T) {
	testes := []struct {
		nome    string
		entrada []ItemColeta
	}{
		{"nil", nil},
		{"slice vazio", []ItemColeta{}},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			obtido := ProcessarLoteConcorrente(tt.entrada)

			if obtido == nil {
				t.Fatal("esperado slice vazio não nulo, obtido nil")
			}
			if len(obtido) != 0 {
				t.Errorf("len = %d; esperado 0", len(obtido))
			}
		})
	}
}

func TestProcessarLoteConcorrenteCincoItens(t *testing.T) {
	entrada := []ItemColeta{
		{ID: 1, NomeProduto: "Teclado Redragon"},
		{ID: 2, NomeProduto: "Mouse Gamer Logitech"},
		{ID: 3, NomeProduto: "Headset HyperX"},
		{ID: 4, NomeProduto: "Monitor UltraWide LG"},
		{ID: 5, NomeProduto: "Cadeira Gamer ThunderX3"},
	}
	esperado := make([]ItemColeta, len(entrada))
	for i, it := range entrada {
		esperado[i] = ItemColeta{ID: it.ID, NomeProduto: it.NomeProduto, Processado: true, Tag: "COLETA_PENDENTE"}
	}
	copiaEntrada := slices.Clone(entrada)

	obtido := ProcessarLoteConcorrente(entrada)

	if !slices.Equal(obtido, esperado) {
		t.Errorf("obtido %+v\nesperado %+v", obtido, esperado)
	}
	for i, it := range obtido {
		if !it.Processado {
			t.Errorf("índice %d (%s): Processado = false", i, it.NomeProduto)
		}
		if it.ID != entrada[i].ID {
			t.Errorf("índice %d: ID = %d; esperado %d (ordem quebrada)", i, it.ID, entrada[i].ID)
		}
	}
	if !slices.Equal(entrada, copiaEntrada) {
		t.Errorf("a entrada foi modificada: %+v", entrada)
	}
}

// Estresse: muitos itens, para o -race e a checagem de ordem terem o que pegar.
func TestProcessarLoteConcorrenteEstresse(t *testing.T) {
	const n = 500
	entrada := make([]ItemColeta, n)
	for i := range entrada {
		entrada[i] = ItemColeta{ID: i, NomeProduto: fmt.Sprintf("Produto-%d", i)}
	}

	obtido := ProcessarLoteConcorrente(entrada)

	if len(obtido) != n {
		t.Fatalf("len = %d; esperado %d", len(obtido), n)
	}
	for i, it := range obtido {
		if it.ID != i || !it.Processado || it.Tag != "COLETA_PENDENTE" {
			t.Fatalf("índice %d: %+v", i, it)
		}
	}
}
