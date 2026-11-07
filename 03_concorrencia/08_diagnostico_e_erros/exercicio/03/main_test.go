package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
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

// aguardar repete a condição até ser verdadeira ou o prazo acabar.
func aguardar(t *testing.T, descricao string, cond func() bool) {
	t.Helper()
	limite := time.Now().Add(2 * time.Second)
	for time.Now().Before(limite) {
		if cond() {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatalf("condição não atendida a tempo: %s", descricao)
}

// produtor envia n itens e fecha; aborta se o ctx cancelar.
// Fecha 'terminou' ao sair, para o teste saber que a goroutine não ficou solta.
func produtor(ctx context.Context, n int, saida chan<- string, terminou chan<- struct{}) {
	defer close(terminou)
	for i := 0; i < n; i++ {
		select {
		case saida <- fmt.Sprintf("item-%d", i):
		case <-ctx.Done():
			return
		}
	}
	close(saida)
}

func TestProcessarFilaComSeguranca(t *testing.T) {
	t.Run("normal: 4 itens e canal fechado retornam 4, nil", func(t *testing.T) {
		ch := make(chan string, 4)
		for i := range 4 {
			ch <- fmt.Sprintf("item-%d", i)
		}
		close(ch)

		n, err := ProcessarFilaComSeguranca(context.Background(), ch, func(string) error { return nil })

		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		assertIgual(t, n, 4)
	})

	t.Run("itens com erro não entram na contagem", func(t *testing.T) {
		ch := make(chan string, 5)
		for _, s := range []string{"ok", "ruim", "ok", "ruim", "ok"} {
			ch <- s
		}
		close(ch)

		n, err := ProcessarFilaComSeguranca(context.Background(), ch, func(s string) error {
			if s == "ruim" {
				return errors.New("falhou")
			}
			return nil
		})

		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		assertIgual(t, n, 3)
	})

	t.Run("canal vazio e fechado: 0, nil", func(t *testing.T) {
		ch := make(chan string)
		close(ch)

		n, err := ProcessarFilaComSeguranca(context.Background(), ch, func(string) error { return nil })

		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		assertIgual(t, n, 0)
	})

	t.Run("timeout: 100 itens, contexto expira com contagem parcial", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 55*time.Millisecond)
		defer cancel()

		ch := make(chan string)
		terminou := make(chan struct{})
		go produtor(ctx, 100, ch, terminou)

		n, err := ProcessarFilaComSeguranca(ctx, ch, func(string) error {
			time.Sleep(10 * time.Millisecond) // ~10ms por item: 100 itens levariam 1s
			return nil
		})

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("erro = %v; esperado %v", err, context.DeadlineExceeded)
		}
		if n <= 0 || n >= 100 {
			t.Errorf("contagem = %d; esperado parcial (entre 1 e 99)", n)
		}

		// O produtor também precisa terminar (respeitou o ctx)
		select {
		case <-terminou:
		case <-time.After(time.Second):
			t.Error("o produtor ficou preso: goroutine leak")
		}
	})

	t.Run("cancelamento manual: Canceled com contagem parcial", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		var processados atomic.Int64

		ch := make(chan string)
		terminou := make(chan struct{})
		go produtor(ctx, 100, ch, terminou)

		n, err := ProcessarFilaComSeguranca(ctx, ch, func(string) error {
			if processados.Add(1) == 3 {
				cancel() // cancela durante o 3º item
			}
			return nil
		})

		if !errors.Is(err, context.Canceled) {
			t.Fatalf("erro = %v; esperado %v", err, context.Canceled)
		}
		assertIgual(t, n, 3)
		<-terminou
	})

	t.Run("contexto já cancelado antes da chamada", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		ch := make(chan string) // ninguém envia: só o ctx.Done() está pronto

		n, err := ProcessarFilaComSeguranca(ctx, ch, func(string) error { return nil })

		if !errors.Is(err, context.Canceled) {
			t.Fatalf("erro = %v; esperado %v", err, context.Canceled)
		}
		assertIgual(t, n, 0)
	})
}

// 3. Ausência de leaks: o número de goroutines volta ao patamar inicial
func TestProcessarFilaComSegurancaSemLeaks(t *testing.T) {
	antes := runtime.NumGoroutine()

	for range 20 {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)

		ch := make(chan string)
		terminou := make(chan struct{})
		go produtor(ctx, 1000, ch, terminou)

		ProcessarFilaComSeguranca(ctx, ch, func(string) error {
			time.Sleep(2 * time.Millisecond)
			return nil
		})

		cancel()
		<-terminou
	}

	aguardar(t, "goroutines voltaram ao patamar inicial", func() bool {
		return runtime.NumGoroutine() <= antes
	})
}

// produtor SEM ctx, preso num envio que ninguém lê.
func TestProdutorSemCtxVazaGoroutine(t *testing.T) {
	antes := runtime.NumGoroutine()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	ch := make(chan string) // sem buffer
	go func() {
		ch <- "ninguém vai ler" // fica preso aqui para sempre
	}()

	ProcessarFilaComSeguranca(ctx, ch, func(string) error { return nil })
	time.Sleep(50 * time.Millisecond)

	depois := runtime.NumGoroutine()
	if depois <= antes {
		t.Skip("o leak esperado não aconteceu nesta execução")
	}
	t.Logf("goroutines: antes=%d depois=%d (o produtor sem ctx ficou preso, como esperado)", antes, depois)
	// Solta a goroutine presa para não contaminar os outros testes
	go func() { <-ch }()
}
