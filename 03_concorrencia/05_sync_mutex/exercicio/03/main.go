package main

import (
	"fmt"
	"sync"
)

type RateLimiter struct {
	mu           sync.Mutex
	limiteMaximo int
	contagem     map[string]int
}

func NovoRateLimiter(limiteMaximo int) *RateLimiter {
	return &RateLimiter{
		limiteMaximo: limiteMaximo,
		contagem:     make(map[string]int),
	}
}

func (rateLimiter *RateLimiter) Permitir(ClienteID string) bool {
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()

	if rateLimiter.contagem[ClienteID] < rateLimiter.limiteMaximo {
		rateLimiter.contagem[ClienteID]++
		return true
	}

	return false
}

func (rateLimiter *RateLimiter) Resetar(ClienteID string) {
	rateLimiter.mu.Lock()
	defer rateLimiter.mu.Unlock()

	delete(rateLimiter.contagem, ClienteID)
}
func main() {
	fmt.Println("Exercicio Rate Limiter Example")

	ratelimiter := NovoRateLimiter(2)

	for i := 1; i <= 4; i++ {
		fmt.Printf("cliente-fedora, requisição %d: %t\n", i, ratelimiter.Permitir("cliente-fedora"))
	}
	fmt.Printf("cliente-debian, requisição 1: %t (contagem própria)\n", ratelimiter.Permitir("cliente-debian"))

	ratelimiter.Resetar("cliente-fedora")
	fmt.Printf("cliente-fedora após Resetar: %t\n", ratelimiter.Permitir("cliente-fedora"))
}
