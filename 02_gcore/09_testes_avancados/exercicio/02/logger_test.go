package main

import (
	"testing"
	"strings"
)

func assertSemErro(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
}

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

func TestFormatadores(t *testing.T) {

	testes := []struct {
		nome     string
		servico  string
		msg      string
		esperado string
	}{
		{"caso comum", "auth", "usuario logado", "[auth] usuario logado"},
		{"mensagem vazia", "db", "", "[db] "},
		{"serviço vazio", "", "ok", "[] ok"},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			assertIgual(t, FormatadorComPlus(tt.servico, tt.msg), tt.esperado)
			assertIgual(t, FormatadorComSprintf(tt.servico, tt.msg), tt.esperado)
		})
	}
}

func BenchmarkFormatadorPlus(b *testing.B) {
	for b.Loop() {
		FormatadorComPlus("auth", "usuario logado com sucesso")
	}
}

func BenchmarkFormatadorSprintf(b *testing.B) {
	for b.Loop() {
		FormatadorComSprintf("auth", "usuario logado com sucesso")
	}
}

var (
	servicoBench = strings.Repeat("a", 4)
	msgBench     = strings.Repeat("b", 26)
)

func BenchmarkFormatadorSprintfVar(b *testing.B) {
	for b.Loop() {
		FormatadorComSprintf(servicoBench, msgBench)
	}
}