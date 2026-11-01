package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

// 1. Table-driven: comportamento sequencial
func TestRateLimiterPermitir(t *testing.T) {
	testes := []struct {
		nome     string
		limite   int
		clientes []string // sequência de chamadas a Permitir
		esperado []bool
	}{
		{"limite 2, mesmo cliente", 2, []string{"a", "a", "a"}, []bool{true, true, false}},
		{"clientes têm contagens independentes", 1, []string{"a", "b", "a", "b"}, []bool{true, true, false, false}},
		{"limite 0 nega tudo", 0, []string{"a"}, []bool{false}},
		{"limite 3, quatro chamadas", 3, []string{"a", "a", "a", "a"}, []bool{true, true, true, false}},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			rl := NovoRateLimiter(tt.limite)

			for i, cliente := range tt.clientes {
				assertIgual(t, rl.Permitir(cliente), tt.esperado[i])
			}
		})
	}
}

func TestRateLimiterResetar(t *testing.T) {
	rl := NovoRateLimiter(1)

	assertIgual(t, rl.Permitir("a"), true)
	assertIgual(t, rl.Permitir("a"), false) // atingiu o limite
	assertIgual(t, rl.Permitir("b"), true)

	rl.Resetar("a")

	assertIgual(t, rl.Permitir("a"), true)  // liberado de novo
	assertIgual(t, rl.Permitir("b"), false) // o reset de "a" não afetou "b"

	rl.Resetar("nunca-visto") // resetar cliente inexistente não pode dar panic
}

// 2. Teste do enunciado: 20 goroutines, limite 10 => exatamente 10 true e 10 false
func TestRateLimiterConcorrente(t *testing.T) {
	const (
		limite     = 10
		goroutines = 20
	)
	rl := NovoRateLimiter(limite)

	var permitidas, negadas atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rl.Permitir("cliente-fedora") {
				permitidas.Add(1)
			} else {
				negadas.Add(1)
			}
		}()
	}
	wg.Wait()

	assertIgual(t, permitidas.Load(), int64(limite))
	assertIgual(t, negadas.Load(), int64(goroutines-limite))
	assertIgual(t, rl.contagem["cliente-fedora"], limite)
}

// Vários clientes disputando ao mesmo tempo: cada um deve ter exatamente o seu limite
func TestRateLimiterConcorrenteVariosClientes(t *testing.T) {
	const (
		limite     = 5
		clientes   = 4
		porCliente = 30
	)
	rl := NovoRateLimiter(limite)

	permitidas := make([]atomic.Int64, clientes)
	var wg sync.WaitGroup

	for c := 0; c < clientes; c++ {
		id := fmt.Sprintf("cliente-%d", c)
		for i := 0; i < porCliente; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if rl.Permitir(id) {
					permitidas[c].Add(1)
				}
			}()
		}
	}
	wg.Wait()

	for c := 0; c < clientes; c++ {
		assertIgual(t, permitidas[c].Load(), int64(limite))
	}
}

func BenchmarkRateLimiterPermitir(b *testing.B) {
	rl := NovoRateLimiter(1 << 30) // limite enorme: mede o custo do lock, não o bloqueio

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rl.Permitir("cliente-fedora")
		}
	})
}
