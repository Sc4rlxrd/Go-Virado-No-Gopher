package main

import (
	"fmt"
	"slices"
	"sync"
	"testing"
)

func lerTudo(ch <-chan string) []string {
	lidas := []string{}
	for msg := range ch {
		lidas = append(lidas, msg)
	}
	return lidas
}

func TestDuplicarCanal(t *testing.T) {
	mensagens := []string{"msg1", "msg2", "msg3"}

	entrada := make(chan string, 3)
	for _, m := range mensagens {
		entrada <- m
	}
	close(entrada)

	saida1 := make(chan string, 3)
	saida2 := make(chan string, 3)

	go DuplicarCanal(entrada, saida1, saida2)

	obtido1 := lerTudo(saida1)
	obtido2 := lerTudo(saida2)

	if !slices.Equal(obtido1, mensagens) {
		t.Errorf("saida1 = %q; esperado %q", obtido1, mensagens)
	}
	if !slices.Equal(obtido2, mensagens) {
		t.Errorf("saida2 = %q; esperado %q", obtido2, mensagens)
	}
}

// Com buffers pequenos ou sem buffer, os consumidores precisam rodar
// ao mesmo tempo, senão o teste trava
func TestDuplicarCanalBuffersDiferentes(t *testing.T) {
	mensagens := []string{"a", "b", "c", "d", "e"}

	for _, buffer := range []int{0, 1, 3, 10} {
		t.Run(fmt.Sprintf("buffer=%d", buffer), func(t *testing.T) {
			entrada := make(chan string)
			saida1 := make(chan string, buffer)
			saida2 := make(chan string, buffer)

			go func() {
				defer close(entrada)
				for _, m := range mensagens {
					entrada <- m
				}
			}()
			go DuplicarCanal(entrada, saida1, saida2)

			var (
				wg               sync.WaitGroup
				obtido1, obtido2 []string
			)
			wg.Add(2)
			go func() { defer wg.Done(); obtido1 = lerTudo(saida1) }()
			go func() { defer wg.Done(); obtido2 = lerTudo(saida2) }()
			wg.Wait()

			if !slices.Equal(obtido1, mensagens) {
				t.Errorf("saida1 = %q; esperado %q", obtido1, mensagens)
			}
			if !slices.Equal(obtido2, mensagens) {
				t.Errorf("saida2 = %q; esperado %q", obtido2, mensagens)
			}
		})
	}
}

func TestDuplicarCanalEntradaVazia(t *testing.T) {
	entrada := make(chan string)
	close(entrada) // fechada sem nenhuma mensagem

	saida1 := make(chan string, 1)
	saida2 := make(chan string, 1)

	go DuplicarCanal(entrada, saida1, saida2)

	if obtido := lerTudo(saida1); len(obtido) != 0 {
		t.Errorf("saida1 = %q; esperado vazia", obtido)
	}
	if obtido := lerTudo(saida2); len(obtido) != 0 {
		t.Errorf("saida2 = %q; esperado vazia", obtido)
	}
}
