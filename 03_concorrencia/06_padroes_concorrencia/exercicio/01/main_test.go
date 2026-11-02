package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

func gerarURLs(n int) []string {
	urls := make([]string, n)
	for i := range urls {
		if i%3 == 2 { // uma a cada três é http
			urls[i] = fmt.Sprintf("http://loja%d.com", i)
		} else {
			urls[i] = fmt.Sprintf("https://loja%d.com", i)
		}
	}
	return urls
}

func TestExecutarWorkerPool(t *testing.T) {
	testes := []struct {
		nome       string
		qtdURLs    int
		numWorkers int
	}{
		{"enunciado: 10 URLs e 3 workers", 10, 3},
		{"20 URLs e 4 workers", 20, 4},
		{"mais workers que URLs", 3, 8},
		{"1 worker processa tudo", 7, 1},
		{"sem URLs", 0, 3},
		{"sem URLs e sem workers", 0, 0},
		{"1000 URLs e 10 workers", 1000, 10},
		{"1000 URLs e 1 worker", 1000, 1},
		{"1000 URLs e 100 workers", 1000, 100},
		{"9999 URLs e 10 workers", 9999, 10},
		{"10000 URLs e 10 workers", 10000, 10},
		{"10001 URLs e 1 workers", 10001, 1},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			urls := gerarURLs(tt.qtdURLs)

			resp := ExecutarWorkerPool(urls, tt.numWorkers)

			assertIgual(t, len(resp), tt.qtdURLs) // sem duplicatas e sem perdas

			slices.SortFunc(resp, func(a, b RespostaURL) int { return a.ID - b.ID })

			for i, r := range resp {
				assertIgual(t, r.ID, i+1) // cada ID aparece exatamente uma vez
				assertIgual(t, r.URL, urls[i])
				assertIgual(t, r.Valida, strings.Contains(urls[i], "https://"))

				if r.WorkerId < 1 || r.WorkerId > tt.numWorkers {
					t.Errorf("ID %d: WorkerId = %d; esperado entre 1 e %d", r.ID, r.WorkerId, tt.numWorkers)
				}
			}
		})
	}
}

// Com tarefas demoradas e vários workers, mais de um worker deve participar.
func TestExecutarWorkerPoolUsaVariosWorkers(t *testing.T) {
	resp := ExecutarWorkerPool(gerarURLs(20), 4)

	usados := map[int]bool{}
	for _, r := range resp {
		usados[r.WorkerId] = true
	}
	if len(usados) < 2 {
		t.Errorf("apenas %d worker(s) participaram; esperado mais de um", len(usados))
	}
}
