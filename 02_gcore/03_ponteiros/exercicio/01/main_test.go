package main

import (
	"testing"
)
func TestReiniciarPorta(t *testing.T) {

	resultado := 3000
	ReiniciarPorta(&resultado)
	esperado := 8080

	if resultado != esperado {
    	t.Errorf("resultado esperado: %d, recebido: %d", esperado, resultado)
	}

}
