package main

import (
	"testing"
)

func TestBuscarComoString_Sucesso(t *testing.T) {

	cache := &MemoriaCache{
		dados: map[string]any{
			"nome":  "Golang",
			"idade": 10,
		},
	}

	resultado, err := BuscarComoString(cache, "nome")
	if err != nil {
		t.Fatalf("não esperava um erro, mas recebeu: %v", err)

	} else if resultado != "Golang" {
		t.Errorf("resultado incorreto: esperado %q, recebido %q", "Golang", resultado)
	}
}

func TestBuscarComoString_Falha(t *testing.T) {

	cache := &MemoriaCache{
		dados: map[string]any{
			"nome":  "Golang",
			"idade": 10,
		},
	}

	resultado2, err := BuscarComoString(cache, "idade")
	if err == nil {
		t.Fatal("esperava um erro porque o valor armazenado não é uma string")
	}

	if resultado2 != "" {
		t.Errorf("esperava uma string vazia, mas recebeu %q", resultado2)
	}

	erroEsperado := "o valor armazenado não é uma string"
	erro, ok := err.(error)
	if !ok {
		t.Fatal("esperava receber um valor que implementasse error")
	}

	if erro.Error() != erroEsperado {
		t.Errorf(
			"erro incorreto: esperado %q, recebido %q",
			erroEsperado,
			erro.Error(),
		)
	}
}
