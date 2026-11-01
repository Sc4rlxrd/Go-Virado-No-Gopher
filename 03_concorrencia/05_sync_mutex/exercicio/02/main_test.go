package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

// insere um item com a idade desejada, sem precisar dormir.
func semar(c *CachePrecos, produto string, preco float64, idade time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.dados[produto] = CotacaoItem{
		Preco:    preco,
		CriadoEm: time.Now().Add(-idade),
	}
}

func TestCachePrecosObter(t *testing.T) {

	testes := []struct {
		nome          string        // nome do cenário
		produto       string        // produto consultado
		idade         time.Duration // idade do "Mouse" semeado no cache
		ttl           time.Duration // ttl usado na consulta
		esperadoPreco float64       // preço esperado
		esperadoOk    bool          // encontrou e válido?
	}{
		{"item novo dentro do ttl: retorna", "Mouse", time.Second, time.Hour, 199.99, true},
		{"item expirado: rejeita", "Mouse", 2 * time.Hour, time.Hour, 0, false},
		{"produto inexistente: rejeita", "Teclado", time.Second, time.Hour, 0, false},
		{"ttl curto expira item de 10s", "Mouse", 10 * time.Second, 5 * time.Second, 0, false},
		{"ttl longo mantém item de 10s", "Mouse", 10 * time.Second, time.Minute, 199.99, true},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			cache := NovoCachePrecos()
			semar(cache, "Mouse", 199.99, tt.idade)

			preco, ok := cache.Obter(tt.produto, tt.ttl)
			assertIgual(t, preco, tt.esperadoPreco)
			assertIgual(t, ok, tt.esperadoOk)
		})
	}
}

// Salvar de verdade (com time.Now) e sobrescrever
func TestCachePrecosSalvar(t *testing.T) {
	c := NovoCachePrecos()
	c.Salvar("Mouse", 199.99)
	preco, ok := c.Obter("Mouse", time.Hour)
	assertIgual(t, preco, 199.99)
	assertIgual(t, ok, true)

	// sobrescrever
	c.Salvar("Mouse", 299.99)
	preco, ok = c.Obter("Mouse", time.Hour)
	assertIgual(t, preco, 299.99)
	assertIgual(t, ok, true)
}

// LimparExpirados remove só o que passou do ttl
func TestCachePrecosLimparExpirados(t *testing.T) {
	testes := []struct {
		nome            string        // nome do cenário
		produto         string        // produto consultado
		idade           time.Duration // idade do "Mouse" semeado no cache
		ttl             time.Duration // ttl usado na consulta
		esperadoPreco   float64       // preço esperado
		esperadoOk      bool          // encontrou e válido?
		esperadoTamanho int           // tamanho do cache após a limpeza
	}{
		{"item novo dentro do ttl: mantém", "Mouse", time.Second, time.Hour, 199.99, true, 1},
		{"item expirado: remove", "Mouse", 2 * time.Hour, time.Hour, 0, false, 0},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			cache := NovoCachePrecos()
			semar(cache, tt.produto, 199.99, tt.idade)

			cache.LimparExpirados(tt.ttl)

			preco, ok := cache.Obter(tt.produto, tt.ttl)
			assertIgual(t, preco, tt.esperadoPreco)
			assertIgual(t, ok, tt.esperadoOk)
			assertIgual(t, len(cache.dados), tt.esperadoTamanho)
		})
	}
}

// Teste concorrente do enunciado: leitores, escritores e limpeza ao mesmo tempo
func TestCachePrecosConcorrente(t *testing.T) {
	const n = 100
	c := NovoCachePrecos()
	semar(c, "velho", 99, 2*time.Hour) // as limpezas concorrentes devem removê-lo

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		produto := fmt.Sprintf("produto-%d", i)

		wg.Add(3)
		go func() { // escritor
			defer wg.Done()
			c.Salvar(produto, float64(i))
		}()
		go func() { // leitor (pode ou não ver o item, sem asserção aqui)
			defer wg.Done()
			c.Obter(produto, time.Hour)
		}()
		go func() { // limpeza com ttl de 1h: não pode apagar itens recém-salvos
			defer wg.Done()
			c.LimparExpirados(time.Hour)
		}()
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		produto := fmt.Sprintf("produto-%d", i)
		preco, ok := c.Obter(produto, time.Hour)
		assertIgual(t, ok, true)
		assertIgual(t, preco, float64(i))
	}
	_, ok := c.Obter("velho", 24*time.Hour)
	assertIgual(t, ok, false)
}

// Benchmark de Performance (testing.B)
func BenchmarkCachePrecosObter(b *testing.B) {
	c := NovoCachePrecos()
	c.Salvar("Mouse", 199.99)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Obter("Mouse", time.Hour)
		}
	})
}
