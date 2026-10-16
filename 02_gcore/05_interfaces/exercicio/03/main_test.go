package main

import (
	"testing"
)

func TestDeployApp_Sucesso(t *testing.T) {

	mock := &MockExecutador{
		DeveFalhar:    false,
		SaidaSimulada: "Container iniciado com sucesso",
	}

	err := DeployApp(mock)
	if err != nil {
		t.Errorf("esperava que o deploy desse sucesso, mas retornou erro: %v", err)
	}
}

func TestDeployApp_Falhar(t *testing.T) {

	mock := &MockExecutador{
		DeveFalhar:    true,
		SaidaSimulada: "",
	}

	err := DeployApp(mock)
	if err == nil {
		t.Error("esperava um erro no deploy, mas o resultado foi sucesso")
	}
}
