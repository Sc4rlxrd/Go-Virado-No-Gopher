package main

import (
	"maps"
	"testing"
)

func TestValidarUrl(t *testing.T) {
	testes := []struct {
		nome     string
		url      string
		esperado bool
	}{
		{"https: válida", "https://golang.org/doc/", true},
		{"http: inválida", "http://example.com", false},
		{"vazia: inválida", "", false},
		{"sem esquema: inválida", "golang.org", false},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			if obtido := ValidarURL(tt.url); obtido != tt.esperado {
				t.Errorf("ValidarUrl(%q) = %v; esperado %v", tt.url, obtido, tt.esperado)
			}
		})
	}
}

func TestValidarListUrls(t *testing.T) {
	urls := []string{
		"http://example.com/api/v1/recurso",
		"https://proxy.golang.org/",
		"https://golang.org/doc/",
		"http://api.github.com/users/google",
	}

	esperado := map[string]bool{
		"http://example.com/api/v1/recurso":  false,
		"https://proxy.golang.org/":          true,
		"https://golang.org/doc/":            true,
		"http://api.github.com/users/google": false,
	}

	obtido := ValidarListaURLs(urls)

	for !maps.Equal(obtido, esperado) {
		t.Errorf("ValidarListaURLs() = %v; esperado %v", obtido, esperado)

	}
}
