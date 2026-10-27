package main


import (
	"testing"
	"time"
	"fmt"
)

func TestAguardarResultadoComTimeout(t *testing.T) {

	testes:= [] struct{
		nome string
		atraso time.Duration // quanto o produtor demora para enviar
		timeout time.Duration
		esperaErro bool
	}{
			{"sucesso: produtor responde antes do timeout", 10 * time.Millisecond, 500 * time.Millisecond, false},
		{"timeout: produtor demora mais que o limite", 300 * time.Millisecond, 20 * time.Millisecond, true},
	}

	for _,tt:= range testes{
		t.Run(tt.nome, func(t *testing.T) {
			ch:= make(chan string,1) // buffer evita goroutine presa no cenário de timeout
			go func() {
				time.Sleep(tt.atraso)
				ch <- "199.90"
			}()
			res, err := AguardarResultadoComTimeout(ch, tt.timeout)

			if tt.esperaErro {
				esperado := fmt.Sprintf("operacao expirou apos %v", tt.timeout)
				if err == nil || err.Error() != esperado {
					t.Fatalf("err = %v; esperado %q", err, esperado)
				}
				if res != "" {
					t.Errorf("res = %q; esperado o zero value (string vazia)", res)
				}
				return
			}
			if err!= nil{
				t.Fatalf("erro inesperado: %v", err)
			}
			if res != "199.90" {
				t.Errorf("res = %q; esperado %q", res, "199.90")
			}
		})
	}
	
}

// A função é genérica: confirma que funciona com outro tipo e devolve o zero value dele.
func TestAguardarResultadoComTimeoutGenerico(t *testing.T) {
	ch:= make(chan int)

	res,err:= AguardarResultadoComTimeout(ch, 10*time.Millisecond)

		if err == nil {
		t.Fatal("esperado erro de expiração")
	}
	if res != 0 {
		t.Errorf("res = %d; esperado 0 (zero value de int)", res)
	}
}