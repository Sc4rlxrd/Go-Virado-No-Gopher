package main

import (
	"testing"
)

func TestObterPortaServidor_Padrao(t *testing.T) {

	t.Setenv("APP_PORT", "")

	porta, err := ObterPortaServidor()

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if porta != 8080 {
		t.Errorf("porta = %d; esperado 8080", porta)
	}
}
func TestObterPortaServidor_PortaConfigurada(t *testing.T) {
	t.Setenv("APP_PORT", "9090")

	porta, err := ObterPortaServidor()

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if porta != 9090 {
		t.Errorf("porta = %d; esperado 9090", porta)
	}
}
