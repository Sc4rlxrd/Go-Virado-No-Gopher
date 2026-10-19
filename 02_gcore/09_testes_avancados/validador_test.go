package main

import (
	"errors"
	"testing"
)

// Helper para asserções repetitivas de igualdade
func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper() // Informa ao Go que esta função é um helper
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

// 1. Table-Driven Test com Subtests (t.Run)
func TestValidarHostFedora(t *testing.T) {
	// Tabela de cenários de teste usando struct anônima
	testes := []struct {
		nome      string // nome do cenário
		inputHost string // entrada
		esperado  string // saída esperada
		erro      error  // nil = nenhum erro esperado
	}{
		{
			nome:      "Host Válido em Minúsculas",
			inputHost: "srv01.fedora.local",
			esperado:  "srv01.fedora.local",
			erro:      nil,
		},
		{
			nome:      "Host Válido com Maiúsculas (Deve normalizar)",
			inputHost: "SRV-APP.FEDORA.LOCAL",
			esperado:  "srv-app.fedora.local",
			erro:      nil,
		},
		{
			nome:      "Host Vazio (Deve dar erro)",
			inputHost: " ",
			esperado:  "",
			erro:      ErrHostInvalido,
		},
		{
			nome:      "Host Fora do Domínio Fedora (Deve dar erro)",
			inputHost: "srv01.oracle.com",
			esperado:  "",
			erro:      ErrHostDominio,
		},
		{
			nome:      "Espaços nas pontas: remove e aceita",
			inputHost: "  srv02.fedora.local  ",
			esperado:  "srv02.fedora.local",
			erro:      nil,
		},
		{
			nome:      "Só o sufixo, sem nome: rejeita com ErrHostInvalido",
			inputHost: ".fedora.local",
			esperado:  "",
			erro:      ErrHostInvalido,
		},
		{
			nome:      "Somínio errado: rejeita com ErrHostDominio",
			inputHost: "srv01.oracle.com",
			esperado:  "",
			erro:      ErrHostDominio,
		},
		
	}

	// Iteração sobre a tabela executando Subtests
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado, err := ValidarHostFedora(tt.inputHost) // chama com a ENTRADA da linha

			if !errors.Is(err, tt.erro) { // compara com o ERRO da linha
				t.Fatalf("erro obtido = %v, esperado = %v", err, tt.erro)
			}
			assertIgual(t, resultado, tt.esperado) // compara com o ESPERADO da linha
		})
	}
}

// 2. Benchmark de Performance (testing.B)
func BenchmarkValidarHostFedora(b *testing.B) {
	hostValido := "srv01.fedora.local"

	for b.Loop() {
		_, _ = ValidarHostFedora(hostValido)
	}

}
