package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

var (
	errA = errors.New("erro A")
	errB = errors.New("erro B")
)

func tarefaOK(executadas *atomic.Int64) func(context.Context) error {
	return func(ctx context.Context) error {
		executadas.Add(1)
		return nil
	}
}

func tarefaErro(err error, executadas *atomic.Int64) func(context.Context) error {
	return func(ctx context.Context) error {
		executadas.Add(1)
		return err
	}
}

func TestExecutarEAgregarErros(t *testing.T) {
	t.Run("sem erros: 5 tarefas ok retornam slice vazio", func(t *testing.T) {
		var executadas atomic.Int64
		tarefas := make([]func(context.Context) error, 5)
		for i := range tarefas {
			tarefas[i] = tarefaOK(&executadas)
		}

		erros := ExecutarEAgregarErros(context.Background(), tarefas)

		assertIgual(t, len(erros), 0)
		if erros == nil {
			t.Error("esperado slice vazio não nulo, obtido nil")
		}
		assertIgual(t, executadas.Load(), int64(5))
	})

	t.Run("erros parciais: 5 tarefas, 2 erros", func(t *testing.T) {
		var executadas atomic.Int64
		tarefas := []func(context.Context) error{
			tarefaOK(&executadas),
			tarefaErro(errA, &executadas),
			tarefaOK(&executadas),
			tarefaErro(errB, &executadas),
			tarefaOK(&executadas),
		}

		erros := ExecutarEAgregarErros(context.Background(), tarefas)

		assertIgual(t, len(erros), 2)
		if !errors.Is(errors.Join(erros...), errA) {
			t.Errorf("erros %v não contêm %v", erros, errA)
		}
		if !errors.Is(errors.Join(erros...), errB) {
			t.Errorf("erros %v não contêm %v", erros, errB)
		}
		// Todas rodaram até o fim, inclusive as que vieram depois da primeira falha
		assertIgual(t, executadas.Load(), int64(5))
	})

	t.Run("não aborta no primeiro erro: as lentas também terminam", func(t *testing.T) {
		var terminouLenta atomic.Bool
		tarefas := []func(context.Context) error{
			func(ctx context.Context) error { return errA }, // falha na hora
			func(ctx context.Context) error { // demora, e o ctx NÃO foi cancelado
				time.Sleep(40 * time.Millisecond)
				if ctx.Err() == nil {
					terminouLenta.Store(true)
				}
				return nil
			},
		}

		erros := ExecutarEAgregarErros(context.Background(), tarefas)

		assertIgual(t, len(erros), 1)
		if !terminouLenta.Load() {
			t.Error("a tarefa lenta foi interrompida; todas deveriam rodar até o fim")
		}
	})

	t.Run("nil e vazio: slice vazio", func(t *testing.T) {
		assertIgual(t, len(ExecutarEAgregarErros(context.Background(), nil)), 0)
		assertIgual(t, len(ExecutarEAgregarErros(context.Background(), []func(context.Context) error{})), 0)
	})

	t.Run("o ctx é repassado às tarefas", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		tarefas := []func(context.Context) error{
			func(ctx context.Context) error { return ctx.Err() },
		}

		erros := ExecutarEAgregarErros(ctx, tarefas)

		assertIgual(t, len(erros), 1)
		if !errors.Is(erros[0], context.Canceled) {
			t.Errorf("erro = %v; esperado context.Canceled", erros[0])
		}
	})
}

// (rode com -race)
func TestAcumuladorErrosConcorrente(t *testing.T) {
	const (
		goroutines   = 200
		porGoroutine = 10
	)
	var acc AcumuladorErros
	var wg sync.WaitGroup

	for g := range goroutines {
		wg.Go(func() {
			for i := range porGoroutine {
				acc.Adicionar(fmt.Errorf("g%d-e%d", g, i))
				acc.Adicionar(nil)   // nil nunca entra
				_ = acc.TemErros()   // leitura concorrente
				_ = acc.ObterErros() // cópia concorrente com as escritas
			}
		})
	}
	wg.Wait()

	assertIgual(t, len(acc.ObterErros()), goroutines*porGoroutine)
	assertIgual(t, acc.TemErros(), true)
}

func TestAcumuladorErrosCopia(t *testing.T) {
	var acc AcumuladorErros
	acc.Adicionar(errA)

	copia := acc.ObterErros()
	copia[0] = errB // mexer na cópia não pode alterar o acumulador

	assertIgual(t, acc.ObterErros()[0] == errA, true)
	assertIgual(t, len(copia), 1)

	acc.Adicionar(errB)
	assertIgual(t, len(copia), 1) // a cópia não cresce com novas adições
}

func TestAcumuladorErrosVazio(t *testing.T) {
	var acc AcumuladorErros

	assertIgual(t, acc.TemErros(), false)
	assertIgual(t, len(acc.ObterErros()), 0)

	acc.Adicionar(nil)
	assertIgual(t, acc.TemErros(), false) // nil é ignorado
}
