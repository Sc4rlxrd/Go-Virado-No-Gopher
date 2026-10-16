package main

import (
	"errors"
	"fmt"
)

type ExecutadorComando interface {
	Executar(cmd string) (string, error)
}

type SSHExecutador struct {
	IP    string
	Porta int
	User  string
}

func (s *SSHExecutador) Executar(cmd string) (string, error) {

	return "[SSH] Executado com sucesso: ", nil

}

type MockExecutador struct {
	DeveFalhar    bool
	SaidaSimulada string
}

func (m *MockExecutador) Executar(cmd string) (string, error) {
	if m.DeveFalhar {
		return "", errors.New("falha simulada no comando")
	}
	return m.SaidaSimulada, nil
}

func DeployApp(e ExecutadorComando) error {
	
	_, err := e.Executar("docker compose up -d")
	if err != nil {
		return err
	}
	return nil
}

func main() {

	fmt.Println("---- Exercício 03 -----")

	var executorReal ExecutadorComando = &SSHExecutador{
		IP:    "153.163.103.219",
		Porta: 22,
		User:  "Scarlxrd",
	}
	err := DeployApp(executorReal)
	if err != nil {
		fmt.Println("Erro no deploy real:", err)
	} else {
		fmt.Println("Deploy real realizado com sucesso!")
	}

}
