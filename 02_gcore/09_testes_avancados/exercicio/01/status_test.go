package main

import (
	"testing"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper() // Informa ao Go que esta função é um helper
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

func TestClassificarStatusHTTP(t *testing.T) {

	testes := []struct {
		nome     string // nome de cada teste
		codigo   int    // entrada
		esperado string // saída esperada
	}{
		// Sucesso (2xx)
		{"200 OK: sucesso", 200, "Sucesso"},
		{"299 limite superior: sucesso", 299, "Sucesso"},

		// Erro de cliente (4xx)
		{"400 limite inferior: cliente erro", 400, "Cliente Erro"},
		{"404 not found: cliente erro", 404, "Cliente Erro"},
		{"499 limite superior: cliente erro", 499, "Cliente Erro"},

		// Erro de servidor (5xx)
		{"500 limite inferior: server error", 500, "Server error"},
		{"503 unavailable: server error", 503, "Server error"},
		{"599 limite superior: server error", 599, "Server error"},

		// Fora das faixas tratadas
		{"199 logo abaixo de 2xx: desconhecido", 199, "Desconhecido"},
		{"300 redirecionamento: desconhecido", 300, "Desconhecido"},
		{"399 logo abaixo de 4xx: desconhecido", 399, "Desconhecido"},
		{"600 acima de 5xx: desconhecido", 600, "Desconhecido"},
		{"0 inválido: desconhecido", 0, "Desconhecido"},
		{"negativo: desconhecido", -1, "Desconhecido"},
	}

	// Iteração sobre a tabela executando Subtests
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := ClassificarStatusHTTP(tt.codigo)
			assertIgual(t, resultado, tt.esperado) // compara com o ESPERADO da linha
		})
	}

}
