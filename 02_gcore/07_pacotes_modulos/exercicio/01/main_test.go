package main

import (
	"testing"
)

func TestVersaoSistema(t *testing.T) {

	esperado := "v1.0.0-fedora"

	if VersaoSistema != esperado {
		t.Errorf(
			"VersaoSistema = %q; esperado %q",
			VersaoSistema,
			esperado,
		)
	}

}
