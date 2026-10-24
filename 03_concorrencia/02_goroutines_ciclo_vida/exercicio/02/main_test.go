package main

import (
	"fmt"
	"slices"
	"sync"
	"testing"
	"time"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %+v, esperado %+v", obtido, esperado)

	}
}

func TestChecarHealth(t *testing.T) {
	testes := []struct {
		nome     string
		servico  string
		esperado StatusServico
	}{
		{"Auth: online", "Auth", StatusServico{Nome: "Auth", Online: true, CodigoHTTP: 200}},
		{"Payment: online", "Payment", StatusServico{Nome: "Payment", Online: true, CodigoHTTP: 200}},
		{"Catalog: online", "Catalog", StatusServico{Nome: "Catalog", Online: true, CodigoHTTP: 200}},
		{"Catalog-service (prefixo): online", "Catalog-service", StatusServico{Nome: "Catalog-service", Online: true, CodigoHTTP: 200}},
		{"LegacyDB: offline com 500", "LegacyDB", StatusServico{Nome: "LegacyDB", Online: false, CodigoHTTP: 500}},

		// Comportamento atual da sua implementação (o enunciado não define esses casos):
		{"desconhecido: offline com 500", "Posgresql", StatusServico{Nome: "Posgresql", Online: false, CodigoHTTP: 500}},
		{"vazio: offline com 500", "", StatusServico{Nome: "", Online: false, CodigoHTTP: 500}},
		{"minúsculas: case-sensitive, offline", "auth", StatusServico{Nome: "auth", Online: false, CodigoHTTP: 500}},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			assertIgual(t, ChecarHealth(tt.servico), tt.esperado)
		})
	}
}

func TestVerificarTodosHealthChecks(t *testing.T) {
	testes := []struct {
		nome     string
		entrada  []string
		esperado []StatusServico
	}{
		{
			nome:    "3 serviços: mantém a ordem da entrada",
			entrada: []string{"Auth", "LegacyDB", "Payment"},
			esperado: []StatusServico{
				{Nome: "Auth", Online: true, CodigoHTTP: 200},
				{Nome: "LegacyDB", Online: false, CodigoHTTP: 500},
				{Nome: "Payment", Online: true, CodigoHTTP: 200},
			},
		},
		{
			nome:     "1 serviço",
			entrada:  []string{"Catalog"},
			esperado: []StatusServico{{Nome: "Catalog", Online: true, CodigoHTTP: 200}},
		},
		{"lista vazia: resultado vazio", []string{}, []StatusServico{}},
		{"nil: resultado vazio", nil, []StatusServico{}},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			obtido := VerificarTodosHealthChecks(tt.entrada)
			if !slices.Equal(obtido, tt.esperado) {
				t.Errorf("obtido %+v, esperado %+v", obtido, tt.esperado)

			}
		})
	}

}

// Estresse: muitas goroutines terminando fora de ordem.
// Se algum resultado cair no índice errado, ou o wg.Wait() faltar, este teste pega.
func TestVerificarTodosHealthChecksPreservaOrdem(t *testing.T) {
	const n = 200

	servicos := make([]string, n)
	for i := range servicos {
		if i%2 == 0 {
			servicos[i] = fmt.Sprintf("Auth-%d", i)
		} else {
			servicos[i] = fmt.Sprintf("LegacyDB-%d", i)
		}
	}

	obtido := VerificarTodosHealthChecks(servicos)

	if len(obtido) != n {
		t.Fatalf("len(resultados) = %d; esperado %d", len(obtido), n)
	}
	for i, r := range obtido {
		if r.Nome != servicos[i] {
			t.Fatalf("índice %d: Nome = %q; esperado %q", i, r.Nome, servicos[i])
		}
		if r.Online != (i%2 == 0) {
			t.Fatalf("índice %d (%s): Online = %v; esperado %v", i, r.Nome, r.Online, i%2 == 0)
		}
	}
}

// gerarServicos monta n nomes, metade conhecidos e metade LegacyDB.
func gerarServicos(n int) []string {
	servicos := make([]string, n)
	for i := range servicos {
		if i%2 == 0 {
			servicos[i] = fmt.Sprintf("Auth-%d", i)
		} else {
			servicos[i] = fmt.Sprintf("LegacyDB-%d", i)
		}
	}
	return servicos
}

// Versão de referência, só para comparar com a concorrente.
func verificarSequencial(servicos []string) []StatusServico {
	resultados := make([]StatusServico, len(servicos))
	for i, nome := range servicos {
		resultados[i] = ChecarHealth(nome)
	}
	return resultados
}

func BenchmarkVerificarTodosHealthChecks(b *testing.B) {
	for _, n := range []int{3, 100, 10_000} {
		servicos := gerarServicos(n) // setup fora do loop medido

		b.Run(fmt.Sprintf("concorrente/n=%d", n), func(b *testing.B) {
			for b.Loop() {
				VerificarTodosHealthChecks(servicos)
			}
		})

		b.Run(fmt.Sprintf("sequencial/n=%d", n), func(b *testing.B) {
			for b.Loop() {
				verificarSequencial(servicos)
			}
		})
	}
}

func checarComLatencia(nome string) StatusServico {
	time.Sleep(time.Millisecond) // simula uma chamada de rede
	return ChecarHealth(nome)
}

func verificarConcorrenteCom(servicos []string, checar func(string) StatusServico) []StatusServico {
	var wg sync.WaitGroup
	resultados := make([]StatusServico, len(servicos))
	for i, nome := range servicos {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resultados[i] = checar(nome)
		}()
	}
	wg.Wait()
	return resultados
}

func verificarSequencialCom(servicos []string, checar func(string) StatusServico) []StatusServico {
	resultados := make([]StatusServico, len(servicos))
	for i, nome := range servicos {
		resultados[i] = checar(nome)
	}
	return resultados
}

func BenchmarkComLatencia(b *testing.B) {
	servicos := gerarServicos(20)

	b.Run("concorrente", func(b *testing.B) {
		for b.Loop() {
			verificarConcorrenteCom(servicos, checarComLatencia)
		}
	})
	b.Run("sequencial", func(b *testing.B) {
		for b.Loop() {
			verificarSequencialCom(servicos, checarComLatencia)
		}
	})
}
