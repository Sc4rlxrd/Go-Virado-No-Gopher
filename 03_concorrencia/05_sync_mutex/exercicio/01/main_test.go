package main

import (
	"sync"
	"testing"
)

// Helper para asserções repetitivas de igualdade
func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

// Dispara nSucesso + nFalha goroutines ao mesmo tempo e espera todas terminarem.
func dispararConcorrente(c *ContadorMetricas, nSucesso, nFalha int) {
	var wg sync.WaitGroup

	for i := 0; i < nSucesso; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.IncrementarSucesso()
		}()
	}
	for i := 0; i < nFalha; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.IncrementarFalha()
		}()
	}

	wg.Wait()
}

func TestContadorMetricasConcorrente(t *testing.T) {
	testes := []struct {
		nome             string // nome do cenário
		sucessos         int    // goroutines que chamam IncrementarSucesso
		falhas           int    // goroutines que chamam IncrementarFalha
		esperadoSucessos int    // total esperado de sucessos
		esperadoFalhas   int    // total esperado de falhas
	}{
		{"enunciado: 50 sucessos e 50 falhas", 50, 50, 50, 50},
		{"nenhuma chamada: zeros", 0, 0, 0, 0},
		{"só sucessos", 100, 0, 100, 0},
		{"só falhas", 0, 100, 0, 100},
		{"assimétrico", 70, 30, 70, 30},
		{"estresse: muitas goroutines", 1000, 1000, 1000, 1000},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			c := &ContadorMetricas{} // contador novo por cenário

			dispararConcorrente(c, tt.sucessos, tt.falhas)

			sucesso, falha := c.ObterMetricas()
			assertIgual(t, sucesso, tt.esperadoSucessos)
			assertIgual(t, falha, tt.esperadoFalhas)
		})
	}
}

func BenchmarkContadorMetricas(b *testing.B) {
	c := &ContadorMetricas{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.IncrementarSucesso()
		}
	})
}
