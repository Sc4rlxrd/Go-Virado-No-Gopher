package main

import (
	"context"
	"errors"
	"fmt"
)

type chaveContexto string

const RequestIDChave chaveContexto = "request_id"

var ErrRequestIDAusente = errors.New("request_id ausente no contexto")

// InjetarRequestID devolve um NOVO contexto carregando o request ID.
func InjetarRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDChave, requestID)
}

// ObterRequestID lê o request ID do contexto. ok é false se ausente ou de outro tipo.
func ObterRequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(RequestIDChave).(string)
	return id, ok
}

func BuscarCotacaoComTrace(ctx context.Context, produto string) (string, error) {
	id, ok := ObterRequestID(ctx)
	if !ok || id == "" {
		return "", ErrRequestIDAusente
	}
	return fmt.Sprintf("[REQ: %s] Cotação obtida para %s: R$ 199.90", id, produto), nil
}

func main() {
	fmt.Println("Exercicio 02")

	ctx := InjetarRequestID(context.Background(), "req-123") // guarda o retorno

	msg, err := BuscarCotacaoComTrace(ctx, "Teclado")
	fmt.Println(msg, err)

	// Sem request ID: erro
	_, err = BuscarCotacaoComTrace(context.Background(), "Mouse")
	fmt.Println("sem id:", err)

	// O contexto original não muda: injetar de novo cria outro
	ctx2 := InjetarRequestID(ctx, "req-456")
	id1, _ := ObterRequestID(ctx)
	id2, _ := ObterRequestID(ctx2)
	fmt.Println("original:", id1, "| derivado:", id2)
}
