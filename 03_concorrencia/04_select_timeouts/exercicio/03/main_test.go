package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// chamarComLimite roda f e falha o teste se ela não voltar em 1s.
// Se a função bloqueasse, o teste falharia em vez de travar o terminal.
func chamarComLimite(t *testing.T, f func() bool) bool {
	t.Helper()

	res := make(chan bool, 1)
	go func() { res <- f() }()

	select {
	case r := <-res:
		return r
	case <-time.After(time.Second):
		t.Fatal("a chamada bloqueou: deveria retornar imediatamente")
		return false
	}
}

func TestTentarEnviarMetrica(t *testing.T) {
	t.Run("sucesso: buffer 1 vazio", func(t *testing.T) {
		ch := make(chan string, 1)

		ok := chamarComLimite(t, func() bool { return TentarEnviarMetrica(ch, "cpu=40") })

		if !ok {
			t.Fatal("retornou false; esperado true")
		}
		if len(ch) != 1 {
			t.Errorf("len(ch) = %d; esperado 1", len(ch))
		}
		if got := <-ch; got != "cpu=40" {
			t.Errorf("valor no canal = %q; esperado %q", got, "cpu=40")
		}
	})

	t.Run("descarte: buffer cheio retorna false sem travar", func(t *testing.T) {
		ch := make(chan string, 1)
		ch <- "primeira" // enche o buffer: 100% cheio

		inicio := time.Now()
		ok := chamarComLimite(t, func() bool { return TentarEnviarMetrica(ch, "descartada") })
		duracao := time.Since(inicio)

		if ok {
			t.Fatal("retornou true; esperado false (canal cheio)")
		}
		if duracao > 100*time.Millisecond { // folga grande; deveria levar microssegundos
			t.Errorf("levou %v; esperado retorno imediato", duracao)
		}
		// A métrica antiga continua lá e a descartada não entrou
		if len(ch) != 1 {
			t.Errorf("len(ch) = %d; esperado 1", len(ch))
		}
		if got := <-ch; got != "primeira" {
			t.Errorf("valor no canal = %q; esperado %q", got, "primeira")
		}
	})
}

func TestTentarEnviarMetricaSequencia(t *testing.T) {
	testes := []struct {
		nome     string
		buffer   int
		envios   int
		esperado []bool
	}{
		{"buffer 3, 5 envios", 3, 5, []bool{true, true, true, false, false}},
		{"buffer 1, 2 envios", 1, 2, []bool{true, false}},
		{"buffer 0, sem receptor: nunca entra", 0, 2, []bool{false, false}},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			ch := make(chan string, tt.buffer)

			for i := 0; i < tt.envios; i++ {
				m := fmt.Sprintf("m%d", i)
				ok := chamarComLimite(t, func() bool { return TentarEnviarMetrica(ch, m) })

				if ok != tt.esperado[i] {
					t.Errorf("envio %d: ok = %v; esperado %v", i, ok, tt.esperado[i])
				}
			}
		})
	}
}

func TestTentarEnviarMetricaCanalNil(t *testing.T) {
	var ch chan string // nil: um envio normal travaria para sempre

	ok := chamarComLimite(t, func() bool { return TentarEnviarMetrica(ch, "x") })

	if ok {
		t.Error("canal nil: retornou true; esperado false")
	}
}

// Muitas goroutines disputando um canal pequeno e sem consumidor:
// exatamente cap(ch) devem entrar, o resto deve ser descartado.
func TestTentarEnviarMetricaConcorrente(t *testing.T) {
	const (
		capacidade = 10
		produtores = 200
	)
	ch := make(chan string, capacidade)

	var enviadas, descartadas atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < produtores; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if TentarEnviarMetrica(ch, fmt.Sprintf("m%d", i)) {
				enviadas.Add(1)
			} else {
				descartadas.Add(1)
			}
		}()
	}
	wg.Wait()

	if got := enviadas.Load(); got != capacidade {
		t.Errorf("enviadas = %d; esperado %d", got, capacidade)
	}
	if got := descartadas.Load(); got != produtores-capacidade {
		t.Errorf("descartadas = %d; esperado %d", got, produtores-capacidade)
	}
	if len(ch) != capacidade {
		t.Errorf("len(ch) = %d; esperado %d", len(ch), capacidade)
	}
}