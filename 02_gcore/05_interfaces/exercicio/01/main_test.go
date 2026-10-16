package main

import (
	"testing"
)

func TestExibirRegiao(t *testing.T) {

	resultado := AWSProvedor{
		Provider: "AWS",
		Region:   "Ohio",
	}
	ExibirRegiao(resultado)

	esperado := AWSProvedor{
		Provider: "AWS",
		Region:   "Ohio",
	}
	ExibirRegiao(esperado)

	if resultado != esperado {
		t.Error("Erro")
	}

}
