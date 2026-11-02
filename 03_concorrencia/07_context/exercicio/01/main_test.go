package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

func TestValidarTokenComContexto(t *testing.T) {
	testes := []struct {
		nome        string
		token       string
		criarCtx    func() (context.Context, context.CancelFunc)
		esperadoOk  bool
		esperadoErr error // nil = nenhum erro esperado
	}{
		{
			nome:  "sucesso: token válido antes do timeout",
			token: "token-valido-123",
			criarCtx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 500*time.Millisecond)
			},
			esperadoOk:  true,
			esperadoErr: nil,
		},
		{
			nome:  "token inválido responde antes do timeout",
			token: "outro-token",
			criarCtx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 500*time.Millisecond)
			},
			esperadoOk:  false,
			esperadoErr: ErrTokenInvalido,
		},
		{
			nome:  "token vazio: erro sem esperar o contexto",
			token: "",
			criarCtx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 500*time.Millisecond)
			},
			esperadoOk:  false,
			esperadoErr: ErrTokenVazio,
		},
		{
			nome:  "WithTimeout curto: DeadlineExceeded",
			token: "token-valido-123",
			criarCtx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 5*time.Millisecond)
			},
			esperadoOk:  false,
			esperadoErr: context.DeadlineExceeded,
		},
		{
			nome:  "WithCancel manual durante a verificação: Canceled",
			token: "token-valido-123",
			criarCtx: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(context.Background())
				time.AfterFunc(10*time.Millisecond, cancel) // cancela no meio dos 50ms
				return ctx, cancel
			},
			esperadoOk:  false,
			esperadoErr: context.Canceled,
		},
		{
			nome:  "contexto já cancelado antes da chamada: Canceled",
			token: "token-valido-123",
			criarCtx: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx, cancel
			},
			esperadoOk:  false,
			esperadoErr: context.Canceled,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			ctx, cancel := tt.criarCtx()
			defer cancel()

			inicio := time.Now()
			ok, err := ValidarTokenComContexto(ctx, tt.token)
			duracao := time.Since(inicio)

			if !errors.Is(err, tt.esperadoErr) {
				t.Fatalf("erro obtido = %v, esperado = %v", err, tt.esperadoErr)
			}
			assertIgual(t, ok, tt.esperadoOk)

			// Cancelamento tem que interromper a espera: não pode levar os 50ms inteiros
			if tt.esperadoErr == context.DeadlineExceeded || tt.esperadoErr == context.Canceled {
				if duracao >= atrasoVerificacao {
					t.Errorf("levou %v; deveria abortar bem antes de %v", duracao, atrasoVerificacao)
				}
			}
		})
	}
}
