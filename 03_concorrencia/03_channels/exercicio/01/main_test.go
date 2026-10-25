package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestPipelineLogs(t *testing.T) {
	mensagens := []string{"Amazon Sucesso", "Kabum Falha", "Magalu Sucesso"}
	canal := make(chan string, 3)

	go GerarLogs(mensagens, canal)

	obtido := ConsumirLogs(canal)

	esperado := []string{
		"[LOG_PROCESSADO] Amazon Sucesso",
		"[LOG_PROCESSADO] Kabum Falha",
		"[LOG_PROCESSADO] Magalu Sucesso",
	}

	if !slices.Equal(obtido, esperado) {
		t.Errorf("obtido %q\nesperado %q", obtido, esperado)
	}
}

func TestPipelineLogsTamanhosDeBuffer(t *testing.T) {
	mensagens := []string{"a", "b", "c", "d", "e"}
	esperado := []string{
		"[LOG_PROCESSADO] a", "[LOG_PROCESSADO] b", "[LOG_PROCESSADO] c",
		"[LOG_PROCESSADO] d", "[LOG_PROCESSADO] e",
	}

	for _, buffer := range []int{0, 1, 3, 10} {
		t.Run(fmt.Sprintf("buffer=%d", buffer), func(t *testing.T) {
			canal := make(chan string, buffer)

			go GerarLogs(mensagens, canal)
			obtido := ConsumirLogs(canal)

			if !slices.Equal(obtido, esperado) {
				t.Errorf("obtido %q\nesperado %q", obtido, esperado)
			}
		})
	}
}

func TestPipelineLogsVazio(t *testing.T) {
	canal := make(chan string, 3)

	go GerarLogs(nil, canal) // fecha o canal sem enviar nada
	obtido := ConsumirLogs(canal)

	if obtido == nil || len(obtido) != 0 {
		t.Errorf("esperado slice vazio não nulo, obtido %#v", obtido)
	}
}
