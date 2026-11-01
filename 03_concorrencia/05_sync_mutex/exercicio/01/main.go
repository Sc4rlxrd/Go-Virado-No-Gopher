package main

import (
	"fmt"
	"sync"
)

type ContadorMetricas struct {
	mu      sync.Mutex
	sucesso int
	falha   int
}

func (c *ContadorMetricas) IncrementarSucesso() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sucesso++
}

func (c *ContadorMetricas) IncrementarFalha() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.falha++
}

func (c *ContadorMetricas) ObterMetricas() (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.sucesso, c.falha
}

func main() {
	fmt.Println("=== Ex 01: Contador de Métricas ===")
	cont := &ContadorMetricas{}
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cont.IncrementarSucesso()
		}(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cont.IncrementarFalha()
		}(i)
	}
	wg.Wait()
	sucesso, falha := cont.ObterMetricas()
	fmt.Printf("Sucesso: %d, Falha: %d\n", sucesso, falha)

}
