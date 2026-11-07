package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

func ProcessarFilaComSeguranca(ctx context.Context, chIn <-chan string, processar func(string) error) (int, error) {
	processados := 0

	for {
		select {
		case <-ctx.Done():
			return processados, ctx.Err()

		case item, ok := <-chIn:
			if !ok { // produtor fechou o canal: fim normal
				return processados, nil
			}
			if err := processar(item); err == nil {
				processados++
			}
		}
	}
}

func main() {
	fmt.Println(" --- Ex 03 --- ")

	fila := make(chan string)
	go func() {
		defer close(fila)
		for _, item := range []string{"a", "b", "ERRO", "c"} {
			fila <- item
		}
	}()

	n, err := ProcessarFilaComSeguranca(context.Background(), fila, func(s string) error {
		if strings.HasPrefix(s, "ERRO") {
			return errors.New("item inválido")
		}
		time.Sleep(5 * time.Millisecond)
		return nil
	})
	fmt.Printf("processados=%d err=%v\n", n, err)
}
