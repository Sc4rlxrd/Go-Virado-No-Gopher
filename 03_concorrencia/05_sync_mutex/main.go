package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache em memória de cotações protegido por RWMutex

type CotacaoCache struct {
	mu       sync.RWMutex
	cotacoes map[string]float64
}

func NovoCacheCotacoes() *CotacaoCache {
	return &CotacaoCache{
		cotacoes: make(map[string]float64),
	}
}

// Atualizar usa Lock() exclusivo (Escrita)
func (c *CotacaoCache) Atualizar(produto string, preco float64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cotacoes[produto] = preco
}

// Obter usa RLock() compartilhado (Leitura rápida em paralelo)
func (c *CotacaoCache) Obter(produto string) (float64, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	preco, ok := c.cotacoes[produto]
	return preco, ok
}

func main() {
	fmt.Println("=== 🔒 Aula 05: Sincronização com sync.Mutex e sync.RWMutex ===")

	cache := NovoCacheCotacoes()
	var wg sync.WaitGroup

	// 1. Disparando 5 Goroutines de Escrita (Atualizando preços em paralelo)
	produtos := []string{"Mouse DevOps", "Teclado Fedora", "Monitor 4K", "Headset Go", "Webcam HD"}

	for _, p := range produtos {
		wg.Add(1)
		go func(produto string) {
			defer wg.Done()
			cache.Atualizar(produto, 199.99)
		}(p)
	}

	// 2. Disparando 10 Goroutines de Leitura em Paralelo
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(5 * time.Millisecond)
			if preco, ok := cache.Obter("Teclado Fedora"); ok {
				fmt.Printf("📖 [Leitor %d] Preço lido: R$ %.2f\n", id, preco)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Operações de leitura e escrita concluídas sem Data Race!")
}
