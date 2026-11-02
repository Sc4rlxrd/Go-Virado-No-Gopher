package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrTokenVazio    = errors.New("Token vazio")
	ErrTokenInvalido = errors.New("Token inválido")
)

const atrasoVerificacao = 50 * time.Millisecond

func ValidarTokenComContexto(ctx context.Context, token string) (bool, error) {
	if token == "" {
		return false, ErrTokenVazio
	}

	type resultado struct {
		valido bool
		err    error
	}

	// Buffer 1: se o ctx vencer a disputa, a goroutine ainda consegue
	// depositar o resultado e terminar, sem ficar presa para sempre.

	ch := make(chan resultado, 1)
	go func() {
		time.Sleep(atrasoVerificacao)
		if token == "token-valido-123" {
			ch <- resultado{true, nil}
			return
		}
		ch <- resultado{false, ErrTokenInvalido}
	}()

	select {
	case <-ctx.Done():
		return false, ctx.Err() // context.DeadlineExceeded ou context.Canceled
	case r := <-ch:
		return r.valido, r.err
	}
}

func main() {
	fmt.Println("Ex 01")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	fmt.Println(ValidarTokenComContexto(ctx, "token-valido-123"))

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel2()
	fmt.Println(ValidarTokenComContexto(ctx2, "token-valido-123"))

}
