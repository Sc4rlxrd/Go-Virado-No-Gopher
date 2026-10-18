package seguranca

import (
	"testing"
)

func TestGerarHash_NaoVazio(t *testing.T) {
	resultado := GerarHash()

	if resultado == "" {
		t.Fatal("GerarHash() retornou uma string vazia")
	}
}

func TestGerarHash_DiferentesChamadas(t *testing.T) {
	primeiro := GerarHash()
	segundo := GerarHash()

	if primeiro == segundo {
		t.Errorf("A chamada do gerar hash foram iguais:\nChamada 1:%s\nChamanda 2:%s", primeiro, segundo)
	}
}
