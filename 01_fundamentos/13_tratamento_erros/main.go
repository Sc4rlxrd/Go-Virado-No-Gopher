package main

import (
	"errors"
	"fmt"
)

// Erro pré-definido como variável de pacote (Sábia convenção em Go)
var ErrDivisaoPorZero = errors.New("não é possível dividir por zero")

// Função que retorna o resultado e um tipo 'error'
func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero
	}
	return a / b, nil // nil indica ausência de erro
}

// Função com erro formatado dinamicamente via fmt.Errorf
func validatePort(port int) (string, error) {
	if port <= 1 || port >= 65535 {
		return "", fmt.Errorf("porta %d é inválida: deve estar entre 1 e 65535", port)
	}
	return fmt.Sprintf("Porta %d liberada no Firewall", port), nil
}

func main() {
	fmt.Println("--- Testando Divisão com Trata de Erro ---")
	res, err := division(10, 2)
	if err != nil {
		fmt.Println("Erro inesperado:", err)
	} else {
		fmt.Println("Resultado 10/2:", res)
	}

	// Tentativa de divisão inválida

	_, errZero := division(10, 0)
	if errZero != nil {
		fmt.Println("Erro capturado com sucesso:", errZero)
	}

	fmt.Println("-- Validação de Porta de Firewall ---")
	msg, errPorta := validatePort(655)
	if errPorta != nil {
		fmt.Println("Falha na validação:", errPorta)
	} else {
		fmt.Println("Sucesso:", msg)
	}
}
