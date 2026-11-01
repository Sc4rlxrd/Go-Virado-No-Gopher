package main

import (
	"fmt"
	"sync"
	"time"
)

type CotacaoItem struct {
	Preco    float64
	CriadoEm time.Time
}

type CachePrecos struct {
	mu    sync.RWMutex
	dados map[string]CotacaoItem
}

func NovoCachePrecos() *CachePrecos {
	return &CachePrecos{
		dados: make(map[string]CotacaoItem),
	}
}

func (c *CachePrecos) Salvar(produto string, preco float64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.dados[produto] = CotacaoItem{
		Preco:    preco,
		CriadoEm: time.Now(),
	}
}

func (c *CachePrecos) Obter(produto string, ttl time.Duration) (float64, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.dados[produto]
	if !exists {
		return 0, false
	}

	if time.Since(item.CriadoEm) > ttl {
		return 0, false
	}

	return item.Preco, true
}

func (c *CachePrecos) LimparExpirados(ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for produto, item := range c.dados {
		if time.Since(item.CriadoEm) > ttl {
			delete(c.dados, produto)
		}
	}
}

func main() {
	fmt.Println("=== Ex 02: Cache de Preços com TTL ===")
	produtos := []string{"Mouse DevOps", "Teclado Fedora", "Monitor 4K", "Headset Go", "Webcam HD"}
	valores := []float64{199.99, 299.99, 499.99, 149.99, 89.99}
	cache := NovoCachePrecos()
	ttl := 5 * time.Second

	// Salvar preços iniciais
	for i, p := range produtos {
		cache.Salvar(p, valores[i])
	}

	// Obter preços antes do TTL expirar
	for _, p := range produtos {
		if preco, ok := cache.Obter(p, ttl); ok {
			fmt.Printf("Preço de %s: %.2f\n", p, preco)
		} else {
			fmt.Printf("Preço de %s não encontrado ou expirado\n", p)
		}
	}

	// Aguardar o TTL expirar
	time.Sleep(ttl + 1*time.Second)

	// Limpar preços expirados
	cache.LimparExpirados(ttl)

	// Tentar obter preços após o TTL expirar
	for _, p := range produtos {
		if preco, ok := cache.Obter(p, ttl); ok {
			fmt.Printf("Preço de %s: %.2f\n", p, preco)
		} else {
			fmt.Printf("Preço de %s não encontrado ou expirado\n", p)
		}
	}
}
