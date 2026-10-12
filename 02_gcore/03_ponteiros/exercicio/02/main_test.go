package main

import (
	"testing"
)

func TestEscalarPod(t *testing.T) {

	podK8s := PodK8s{
		Nome:     "BookCommerce",
		Replicas: 3,
	}

	err := EscalarPod(&podK8s, 66)

	if err != nil {
		t.Fatalf("erro inesperado ao escalar Pod: %v", err)
	}

	esperado := 66

	if podK8s.Replicas != esperado {
		t.Errorf(
			"réplicas esperadas: %d, recebidas: %d",
			esperado,
			podK8s.Replicas,
		)
	}
}
