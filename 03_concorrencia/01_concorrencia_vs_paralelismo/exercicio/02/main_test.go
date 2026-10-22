package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// capturarSaida roda f e devolve tudo o que ela escreveu em os.Stdout.
func capturarSaida(t *testing.T, f func()) string {
	t.Helper()

	original := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("erro ao criar pipe: %v", err)
	}
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = original

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestVerificarLojaAmazon(t *testing.T) {
	resultado := capturarSaida(t, func() {
		VerificarLoja("Amazon")
	})
	esperado := "[price-agent] Coletando preços na loja: Amazon\n"

	if resultado != esperado {
		t.Errorf("VerificarLoja() = %q; esperado %q", resultado, esperado)
	}
}

func TestVerificarLojaMercadoLivre(t *testing.T) {
	resultado := capturarSaida(t, func() {
		VerificarLoja("MercadoLivre")
	})
	esperado := "[price-agent] Coletando preços na loja: MercadoLivre\n"

	if resultado != esperado {
		t.Errorf("VerificarLoja() = %q; esperado %q", resultado, esperado)
	}
}

func TestVerificarLojaKabum(t *testing.T) {
	resultado := capturarSaida(t, func() {
		VerificarLoja("Kabum")
	})
	esperado := "[price-agent] Coletando preços na loja: Kabum\n"

	if resultado != esperado {
		t.Errorf("VerificarLoja() = %q; esperado %q", resultado, esperado)
	}
}

func TestVerificarLojaMagalu(t *testing.T) {
	resultado := capturarSaida(t, func() {
		VerificarLoja("Magalu")
	})
	esperado := "[price-agent] Coletando preços na loja: Magalu\n"

	if resultado != esperado {
		t.Errorf("VerificarLoja() = %q; esperado %q", resultado, esperado)
	}
}

func TestVerificarLojaShopee(t *testing.T) {
	resultado := capturarSaida(t, func() {
		VerificarLoja("Shopee")
	})
	esperado := "[price-agent] Coletando preços na loja: Shopee\n"

	if resultado != esperado {
		t.Errorf("VerificarLoja() = %q; esperado %q", resultado, esperado)
	}
}

