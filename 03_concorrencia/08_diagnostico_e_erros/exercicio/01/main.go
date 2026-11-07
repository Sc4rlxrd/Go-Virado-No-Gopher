package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type TarefaConcorrente func(ctx context.Context) error

func ExecutarComAbortNoPrimeiroErro(ctx context.Context, tarefas []TarefaConcorrente) error {
	if len(tarefas) == 0 {
		return nil
	}

	ctxChild, cancel := context.WithCancel(ctx)
	defer cancel()
	var (
		wg          sync.WaitGroup
		once        sync.Once
		primeiroErr error
	)

	for _, tarefa := range tarefas {
		wg.Add(1)
		go func(t TarefaConcorrente) {
			defer wg.Done()
			if err := t(ctxChild); err != nil {
				once.Do(func() {
					primeiroErr = err // 1º: guarda o erro real
					cancel()          // 2º: avisa as outras para abortarem
				})
			}
		}(tarefa)
	}

	wg.Wait()
	return primeiroErr

}

func main() {
	fmt.Print("--- Ex 01 ---")
	tarefas := []TarefaConcorrente{
		func(ctx context.Context) error { // fica esperando até ser cancelada
			<-ctx.Done()
			fmt.Println("Audit A: cancelled")
			return ctx.Err()
		},
		func(ctx context.Context) error { // falha em 20ms
			time.Sleep(20 * time.Millisecond)
			return errors.New("audit B: price outside the range")
		},
		func(ctx context.Context) error {
			<-ctx.Done()
			fmt.Println("Audit C: cancelled")
			return ctx.Err()
		},
	}
	err := ExecutarComAbortNoPrimeiroErro(context.Background(), tarefas)
	fmt.Println("Resultado: ", err)
}
