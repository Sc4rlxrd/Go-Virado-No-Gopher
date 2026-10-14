package main

import "testing"

func TestRegistarRequisicao(t *testing.T) {

	api := APIRateLimiter{
		LimiteMaximo:      3,
		RequisicoesAtuais: 1,
	}

	err := api.RegistrarRequisicao()

	if err != nil {
		t.Fatalf("erro inesperado ao registrar requisição: %v", err)
	}

	esperado := 2

	if api.RequisicoesAtuais != esperado {
		t.Errorf(
			"requisições esperadas: %d, recebidas: %d",
			esperado,
			api.RequisicoesAtuais,
		)
	}
}
func TestRegistrarRequisicaoLimiteAlcancado(t *testing.T) {

	api := APIRateLimiter{
		LimiteMaximo:      3,
		RequisicoesAtuais: 3,
	}

	err := api.RegistrarRequisicao()

	if err == nil {
		t.Fatal("esperava erro ao atingir o limite máximo, mas veio nil")
	}

	esperado := 3

	if api.RequisicoesAtuais != esperado {
		t.Errorf(
			"requisições esperadas: %d, recebidas: %d",
			esperado,
			api.RequisicoesAtuais,
		)
	}
}

func TestRegistrarRequisicaoAteOLimite(t *testing.T) {

	api := APIRateLimiter{
		LimiteMaximo:      3,
		RequisicoesAtuais: 1,
	}

	sucessos := 0
	falhas := 0

	for i := 0; i < 4; i++ {
		if err := api.RegistrarRequisicao(); err != nil {
			falhas++
		} else {
			sucessos++
		}
	}

	if sucessos != 2 {
		t.Errorf("sucessos esperados: %d, recebidos: %d", 2, sucessos)
	}

	if falhas != 2 {
		t.Errorf("falhas esperadas: %d, recebidas: %d", 2, falhas)
	}

	if api.RequisicoesAtuais != api.LimiteMaximo {
		t.Errorf(
			"requisições esperadas: %d, recebidas: %d",
			api.LimiteMaximo,
			api.RequisicoesAtuais,
		)
	}
}
