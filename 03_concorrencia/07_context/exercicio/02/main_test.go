package main

import (
	"context"
	"errors"
	"sync"
	"testing"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

func TestBuscarCotacaoComTrace(t *testing.T) {
	testes := []struct {
		nome     string
		criarCtx func() context.Context
		produto  string
		esperado string
		erro     error // nil = nenhum erro esperado
	}{
		{
			nome:     "com request ID: retorna a cotação formatada",
			criarCtx: func() context.Context { return InjetarRequestID(context.Background(), "abc-123") },
			produto:  "Teclado",
			esperado: "[REQ: abc-123] Cotação obtida para Teclado: R$ 199.90",
			erro:     nil,
		},
		{
			nome:     "sem request ID: erro de chave ausente",
			criarCtx: func() context.Context { return context.Background() },
			produto:  "Teclado",
			esperado: "",
			erro:     ErrRequestIDAusente,
		},
		{
			nome:     "request ID vazio: erro de chave ausente",
			criarCtx: func() context.Context { return InjetarRequestID(context.Background(), "") },
			produto:  "Teclado",
			esperado: "",
			erro:     ErrRequestIDAusente,
		},
		{
			nome: "chave string pura não colide com a chave tipada",
			criarCtx: func() context.Context {
				return context.WithValue(context.Background(), "request_id", "intruso") //nolint:staticcheck // de propósito
			},
			produto:  "Teclado",
			esperado: "",
			erro:     ErrRequestIDAusente,
		},
		{
			nome: "valor de outro tipo sob a mesma chave: erro",
			criarCtx: func() context.Context {
				return context.WithValue(context.Background(), RequestIDChave, 12345)
			},
			produto:  "Teclado",
			esperado: "",
			erro:     ErrRequestIDAusente,
		},
		{
			nome: "propagação por contexto derivado (WithCancel)",
			criarCtx: func() context.Context {
				ctx, cancel := context.WithCancel(InjetarRequestID(context.Background(), "xyz"))
				t.Cleanup(cancel)
				return ctx
			},
			produto:  "Mouse",
			esperado: "[REQ: xyz] Cotação obtida para Mouse: R$ 199.90",
			erro:     nil,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			obtido, err := BuscarCotacaoComTrace(tt.criarCtx(), tt.produto)

			if !errors.Is(err, tt.erro) {
				t.Fatalf("erro obtido = %v, esperado = %v", err, tt.erro)
			}
			assertIgual(t, obtido, tt.esperado)
		})
	}
}

func TestObterRequestID(t *testing.T) {
	id, ok := ObterRequestID(context.Background())
	assertIgual(t, ok, false)
	assertIgual(t, id, "")

	id, ok = ObterRequestID(InjetarRequestID(context.Background(), "abc"))
	assertIgual(t, ok, true)
	assertIgual(t, id, "abc")
}

// O contexto é imutável: injetar cria outro, o original não muda.
func TestInjetarRequestIDNaoAlteraOOriginal(t *testing.T) {
	base := context.Background()
	comA := InjetarRequestID(base, "A")
	comB := InjetarRequestID(comA, "B") // sobrescreve só no derivado

	_, ok := ObterRequestID(base)
	assertIgual(t, ok, false)

	idA, _ := ObterRequestID(comA)
	assertIgual(t, idA, "A")

	idB, _ := ObterRequestID(comB)
	assertIgual(t, idB, "B")
}

// Contextos são seguros para uso concorrente: várias goroutines lendo o mesmo ctx.
func TestBuscarCotacaoComTraceConcorrente(t *testing.T) {
	ctx := InjetarRequestID(context.Background(), "req-concorrente")

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			msg, err := BuscarCotacaoComTrace(ctx, "Monitor")
			if err != nil {
				t.Errorf("erro inesperado: %v", err)
			}
			if msg != "[REQ: req-concorrente] Cotação obtida para Monitor: R$ 199.90" {
				t.Errorf("msg = %q", msg)
			}
		}()
	}
	wg.Wait()
}
