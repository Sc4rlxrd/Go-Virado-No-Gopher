package main

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"
)

type AcumuladorErros struct {
	mu    sync.Mutex
	erros []error
}

func (a *AcumuladorErros) Adicionar(err error) {
	if err == nil {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.erros = append(a.erros, err)
}

func (a *AcumuladorErros) ObterErros() []error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return slices.Clone(a.erros)
}

func (a *AcumuladorErros) TemErros() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return len(a.erros) > 0
}

func ExecutarEAgregarErros(ctx context.Context, tarefas []func(context.Context) error) []error {
	var (
		wg  sync.WaitGroup
		acc AcumuladorErros
	)

	for _, tarefa := range tarefas {
		wg.Add(1)
		go func(t func(context.Context) error) {
			defer wg.Done()
			acc.Adicionar(t(ctx)) // o nil é ignorado dentro do Adicionar
		}(tarefa)
	}

	wg.Wait()

	erros := acc.ObterErros()
	if erros == nil {
		return []error{} // slice vazio, não nil, para quem for iterar ou usar len()
	}
	return erros
}

func main() {
	fmt.Println(" --- Ex 02 --- ")

	tarefas := []func(context.Context) error{
		func(ctx context.Context) error { return nil },
		func(ctx context.Context) error { return errors.New("loja A fora do ar") },
		func(ctx context.Context) error { return nil },
		func(ctx context.Context) error { return errors.New("preço negativo em B") },
		func(ctx context.Context) error { return nil },
	}

	erros := ExecutarEAgregarErros(context.Background(), tarefas)
	fmt.Printf("%d erros coletados:\n", len(erros))
	for _, e := range erros {
		fmt.Println(" -", e)
	}
}
