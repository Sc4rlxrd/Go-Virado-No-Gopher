package main

import (
	"testing"
	"time"
)

func TestDetectarVazamento(t *testing.T) {

	testes := []struct {
		nome     string
		extras   int // goroutines em background a disparar antes de checar
		limite   int
		esperado bool // vazamento esperado?
	}{

		{"Sem goroutines extras: sem vazamento", 0, 5, false},
		{"10 goroutines em background: vazamento", 10, 5, true},
		{"20 goroutines em background: vazamento", 20, 5, true},
		{"30 goroutines em background: vazamento", 30, 5, true},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			for i := 0; i < tt.extras; i++ {
				go time.Sleep(100 * time.Millisecond)
			}
			vazou, total := DetectarVazamento(tt.limite)
			if vazou != tt.esperado {
				t.Errorf("DetectarVazamento(%d) = %v (total %d); esperado %v",
					tt.limite, vazou, total, tt.esperado)
			}
			time.Sleep(150 * time.Millisecond) // deixa as extras terminarem antes do próximo caso
		})
	}
}

func TestDetectarVazamentoLimiteExato(t *testing.T) {

	_, base := DetectarVazamento(0)
	vazou, total := DetectarVazamento(base)
	if vazou {
		t.Errorf("total %d com limite %d: não deveria ser vazamento", total, base)
	}
	vazou, total = DetectarVazamento(base - 1)
	if !vazou {
		t.Errorf("total %d com limite %d: deveria ser vazamento", total, base-1)
	}
}
