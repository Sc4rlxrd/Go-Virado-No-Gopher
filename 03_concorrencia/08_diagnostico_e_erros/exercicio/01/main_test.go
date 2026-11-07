package main

import (
	"context"
	"errors"
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
	errBoom     = errors.New("falha na auditoria")
	errPrimeiro = errors.New("primeiro erro")
	errSegundo  = errors.New("segundo erro")
)

// Tarefa que termina com sucesso e conta que executou.
func tarefaOK(executadas *atomic.Int64) TarefaConcorrente {
	return func(ctx context.Context) error {
		executadas.Add(1)
		return nil
	}
}

// Tarefa que falha com err depois de um atraso (ou aborta se o ctx cancelar antes).
func tarefaFalha(err error, atraso time.Duration) TarefaConcorrente {
	return func(ctx context.Context) error {
		select {
		case <-time.After(atraso):
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Tarefa que só termina quando cancelada; conta o cancelamento.
// O limite de 3s só existe para o teste não pendurar se o cancelamento não funcionar.
func tarefaAteCancelar(canceladas *atomic.Int64) TarefaConcorrente {
	return func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			canceladas.Add(1)
			return ctx.Err()
		case <-time.After(3 * time.Second):
			return errors.New("nunca foi cancelada")
		}
	}
}

func TestExecutarComAbortNoPrimeiroErro(t *testing.T) {
	t.Run("sucesso: 3 tarefas sem erro retornam nil", func(t *testing.T) {
		var executadas atomic.Int64
		tarefas := []TarefaConcorrente{tarefaOK(&executadas), tarefaOK(&executadas), tarefaOK(&executadas)}

		err := ExecutarComAbortNoPrimeiroErro(context.Background(), tarefas)

		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		assertIgual(t, executadas.Load(), int64(3))
	})

	t.Run("falha: 1 de 3 erra, as outras são canceladas, erro exato", func(t *testing.T) {
		var canceladas atomic.Int64
		tarefas := []TarefaConcorrente{
			tarefaAteCancelar(&canceladas),
			tarefaFalha(errBoom, 10*time.Millisecond),
			tarefaAteCancelar(&canceladas),
		}

		inicio := time.Now()
		err := ExecutarComAbortNoPrimeiroErro(context.Background(), tarefas)
		duracao := time.Since(inicio)

		if !errors.Is(err, errBoom) {
			t.Fatalf("erro = %v; esperado %v", err, errBoom)
		}
		assertIgual(t, err.Error(), "falha na auditoria") // mensagem exata, não Canceled
		// A função espera todas as goroutines, então as 2 já contaram o cancelamento
		assertIgual(t, canceladas.Load(), int64(2))
		if duracao > time.Second {
			t.Errorf("levou %v; deveria abortar logo após a falha (~10ms)", duracao)
		}
	})

	t.Run("timeout no contexto pai: DeadlineExceeded", func(t *testing.T) {
		var canceladas atomic.Int64
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
		defer cancel()
		tarefas := []TarefaConcorrente{
			tarefaAteCancelar(&canceladas),
			tarefaAteCancelar(&canceladas),
			tarefaAteCancelar(&canceladas),
		}

		err := ExecutarComAbortNoPrimeiroErro(ctx, tarefas)

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("erro = %v; esperado %v", err, context.DeadlineExceeded)
		}
		assertIgual(t, canceladas.Load(), int64(3))
	})

	t.Run("cancelamento manual no contexto pai: Canceled", func(t *testing.T) {
		var canceladas atomic.Int64
		ctx, cancel := context.WithCancel(context.Background())
		time.AfterFunc(20*time.Millisecond, cancel)
		tarefas := []TarefaConcorrente{tarefaAteCancelar(&canceladas), tarefaAteCancelar(&canceladas)}

		err := ExecutarComAbortNoPrimeiroErro(ctx, tarefas)

		if !errors.Is(err, context.Canceled) {
			t.Fatalf("erro = %v; esperado %v", err, context.Canceled)
		}
		assertIgual(t, canceladas.Load(), int64(2))
	})

	t.Run("nil e vazio retornam nil", func(t *testing.T) {
		if err := ExecutarComAbortNoPrimeiroErro(context.Background(), nil); err != nil {
			t.Errorf("nil: erro = %v; esperado nil", err)
		}
		if err := ExecutarComAbortNoPrimeiroErro(context.Background(), []TarefaConcorrente{}); err != nil {
			t.Errorf("vazio: erro = %v; esperado nil", err)
		}
	})

	t.Run("o primeiro erro vence; erros posteriores são ignorados", func(t *testing.T) {
		// Tarefas que IGNORAM o ctx: a segunda ainda falha depois do cancelamento
		tarefas := []TarefaConcorrente{
			func(ctx context.Context) error { time.Sleep(10 * time.Millisecond); return errPrimeiro },
			func(ctx context.Context) error { time.Sleep(60 * time.Millisecond); return errSegundo },
		}

		err := ExecutarComAbortNoPrimeiroErro(context.Background(), tarefas)

		if !errors.Is(err, errPrimeiro) {
			t.Errorf("erro = %v; esperado %v", err, errPrimeiro)
		}
	})

	t.Run("espera todas as goroutines terminarem antes de retornar", func(t *testing.T) {
		var terminou atomic.Bool
		tarefas := []TarefaConcorrente{
			func(ctx context.Context) error { return errBoom }, // falha na hora
			func(ctx context.Context) error { // ignora o ctx e demora
				time.Sleep(50 * time.Millisecond)
				terminou.Store(true)
				return nil
			},
		}

		err := ExecutarComAbortNoPrimeiroErro(context.Background(), tarefas)
		if err != nil {
			return
		}

		if !terminou.Load() {
			t.Error("a função retornou antes de a segunda tarefa terminar")
		}
	})
}

// Estresse: muitas goroutines, metade falhando ao mesmo tempo. Roda com -race.
func TestExecutarComAbortNoPrimeiroErroEstresse(t *testing.T) {
	var executadas atomic.Int64
	tarefas := make([]TarefaConcorrente, 200)
	for i := range tarefas {
		if i%2 == 0 {
			tarefas[i] = tarefaOK(&executadas)
		} else {
			tarefas[i] = func(ctx context.Context) error { return errBoom }
		}
	}

	err := ExecutarComAbortNoPrimeiroErro(context.Background(), tarefas)

	if !errors.Is(err, errBoom) {
		t.Errorf("erro = %v; esperado %v", err, errBoom)
	}
}
