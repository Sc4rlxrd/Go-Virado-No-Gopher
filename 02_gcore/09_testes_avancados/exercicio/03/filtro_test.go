package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

func assertSlicesIguais(t *testing.T, obtido, esperado []string) {
	t.Helper()
	if !slices.Equal(obtido, esperado) {
		t.Errorf("falha na asserção: obtido %q, esperado %q", obtido, esperado)
	}
}

func TestFiltrarPrivados(t *testing.T) {
	testes := []struct {
		nome     string
		entrada  []string
		esperado []string
		erro     error // nil = nenhum erro esperado
	}{
		// Ramo de erro
		{"nil: rejeita com ErrListaVazia", nil, nil, ErrListaVazia},
		{"slice vazio: rejeita com ErrListaVazia", []string{}, nil, ErrListaVazia},

		// Cada lado do ||
		{"só prefixo 10.: mantém", []string{"10.0.0.1"}, []string{"10.0.0.1"}, nil},
		{"só prefixo 192.168.: mantém", []string{"192.168.1.1"}, []string{"192.168.1.1"}, nil},

		// Mistura: filtra e preserva a ordem
		{
			nome:     "mistura: mantém só privados na ordem original",
			entrada:  []string{"8.8.8.8", "192.168.1.10", "172.16.0.1", "10.0.0.5", "1.1.1.1"},
			esperado: []string{"192.168.1.10", "10.0.0.5"},
		},

		// Nenhum privado: resultado válido, não erro
		{"nenhum privado: slice vazio sem erro", []string{"8.8.8.8", "172.16.0.1"}, []string{}, nil},

		// Fronteiras do prefixo
		{"100. não casa com 10.", []string{"100.1.1.1"}, []string{}, nil},
		{"192.169. não casa com 192.168.", []string{"192.169.0.1"}, []string{}, nil},
		{"192.168 sem ponto final não casa", []string{"192.168"}, []string{}, nil},
	}

	filtro := FiltroIp{}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado, err := filtro.FiltrarPrivados(tt.entrada)

			if !errors.Is(err, tt.erro) {
				t.Fatalf("erro obtido = %v, esperado = %v", err, tt.erro)
			}
			assertSlicesIguais(t, resultado, tt.esperado)
		})
	}
}

// gerarIPs monta uma entrada com n IPs, metade privados e metade públicos.
func gerarIPs(n int) []string {
	ips := make([]string, n)
	for i := range ips {
		if i%2 == 0 {
			ips[i] = "192.168.1." + strconv.Itoa(i%256)
		} else {
			ips[i] = "8.8.8." + strconv.Itoa(i%256)
		}
	}
	return ips
}

func BenchmarkFiltrarPrivados(b *testing.B) {
	filtro := FiltroIp{}

	for _, n := range []int{10, 1_000, 100_000} {
		ips := gerarIPs(n) // setup fora do loop medido

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for b.Loop() {
				_, _ = filtro.FiltrarPrivados(ips)
			}
		})
	}
}
