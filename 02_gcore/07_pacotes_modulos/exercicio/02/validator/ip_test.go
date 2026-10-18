package validator

import (
	"testing"
)

func TestValidarIP_Sucesso(t *testing.T) {

	var resultado = ValidarIP("47.70.50.105")
	var esperado = true

	if resultado != esperado {
		t.Errorf("ValidarIP() = %v; esperado %v", resultado, esperado)
	}

}

func TestValidarIP_Falha(t *testing.T) {

	var resultado = ValidarIP("999.999.999.999")
	var esperado = false

	if resultado != esperado {
		t.Errorf("ValidarIP() = %v; esperado %v", resultado, esperado)
	}

}
