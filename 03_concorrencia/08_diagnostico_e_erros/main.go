package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Estrutura customizada para coletar múltiplos erros concorrentes de forma segura
type ColetorErrosConcorrente struct {
	mu    sync.Mutex
	erros []error
}

func (c *ColetorErrosConcorrente) Adicionar(err error) {
	if err == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.erros = append(c.erros, err)
}

func (c *ColetorErrosConcorrente) ObterErros() []error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.erros
}

// Simula consulta a um serviço externo que pode falhar ou expirar
func consultarServicoDevOps(ctx context.Context, id int) error {
	if id ==3 {
		return fmt.Errorf("falha de conexão no node %d", id)
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(20 * time.Millisecond):
		return nil
	}
}

func main() {
	fmt.Println("=== 🩺 Aula 08: Diagnóstico, Deadlocks e Tratamento de Erros ===")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	coletor := &ColetorErrosConcorrente{}

	fmt.Println("🚀 Disparando 5 verificações concorrentes de nós de infraestrutura...")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(nodeID int) {
			defer wg.Done()
			err := consultarServicoDevOps(ctx, nodeID)
			if err != nil {
				coletor.Adicionar(err)
			}
		}(i)
	}
	wg.Wait()

	errosEncontrados := coletor.ObterErros()
	if len(errosEncontrados) > 0 {
		fmt.Printf("⚠️ Foram encontrados %d erro(s) durante a execução:\n", len(errosEncontrados))
		for _, err := range errosEncontrados {
			fmt.Println(" •", err)
		}
	} else {
		fmt.Println("✅ Todos os nós foram verificados sem falhas!")
	}
}
