package main

import (
	"slices"
	"testing"
	"time"
)

type resultado struct {
	precos []float64
	status string
}

// rodar executa o monitor em uma goroutine e devolve o resultado por um canal.
func rodar(dados <-chan float64, done <-chan struct{}) <-chan resultado {
	saida := make(chan resultado, 1)
	go func() {
		p, s := MonitorarCotacoesComCancelamento(dados, done)
		saida <- resultado{p, s}
	}()
	return saida
}

func esperar(t *testing.T, saida <-chan resultado) resultado {
	t.Helper()
	select {
	case r := <-saida:
		return r
	case <-time.After(2 * time.Second):
		t.Fatal("a função não retornou a tempo (travou?)")
		return resultado{}
	}
}

func TestMonitorarCotacoesComCancelamento(t *testing.T) {
	t.Run("cancelamento: 2 preços e depois o sinal", func(t *testing.T) {
		dados := make(chan float64) // sem buffer: o envio só completa quando o monitor recebe
		done := make(chan struct{})
		saida := rodar(dados, done)

		dados <- 199.90
		dados <- 205.50
		close(done) // os 2 já foram consumidos; só chDone está pronto agora

		r := esperar(t, saida)

		if r.status != "cancelado_pelo_usuario" {
			t.Errorf("status = %q; esperado %q", r.status, "cancelado_pelo_usuario")
		}
		if esperado := []float64{199.90, 205.50}; !slices.Equal(r.precos, esperado) {
			t.Errorf("precos = %v; esperado %v", r.precos, esperado)
		}
	})

	t.Run("conclusão normal: canal de dados fechado", func(t *testing.T) {
		dados := make(chan float64)
		done := make(chan struct{}) // nunca acionado
		saida := rodar(dados, done)

		for _, p := range []float64{10, 20, 30} {
			dados <- p
		}
		close(dados)

		r := esperar(t, saida)

		if r.status != "canal_finalizado" {
			t.Errorf("status = %q; esperado %q", r.status, "canal_finalizado")
		}
		if esperado := []float64{10, 20, 30}; !slices.Equal(r.precos, esperado) {
			t.Errorf("precos = %v; esperado %v", r.precos, esperado)
		}
	})

	t.Run("cancelado antes de qualquer preço: slice vazio", func(t *testing.T) {
		dados := make(chan float64)
		done := make(chan struct{})
		saida := rodar(dados, done)

		close(done)

		r := esperar(t, saida)

		if r.status != "cancelado_pelo_usuario" {
			t.Errorf("status = %q; esperado %q", r.status, "cancelado_pelo_usuario")
		}
		if r.precos == nil || len(r.precos) != 0 {
			t.Errorf("precos = %#v; esperado slice vazio não nulo", r.precos)
		}
	})

	t.Run("canal de dados fechado sem enviar nada", func(t *testing.T) {
		dados := make(chan float64)
		done := make(chan struct{})
		saida := rodar(dados, done)

		close(dados)

		r := esperar(t, saida)

		if r.status != "canal_finalizado" {
			t.Errorf("status = %q; esperado %q", r.status, "canal_finalizado")
		}
		if len(r.precos) != 0 {
			t.Errorf("precos = %v; esperado vazio", r.precos)
		}
	})
}