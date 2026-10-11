package main

import (
	"errors"
	"fmt"
)

type ServidorCloud struct {
	ID     string
	Nome   string
	CPU    float64
	Status string
}

func ValidadorServidor(s ServidorCloud) error {
	if s.ID == "" {
		return errors.New("ID do servidor é obrigatório")
	}
	if s.CPU < 0 || s.CPU > 100 {
		return fmt.Errorf("porcentagem de CPU inváçida: %.2f", s.CPU)
	}
	return nil

}

func AnalisarInfra(servidores []ServidorCloud) (map[string]int, error) {
	relatorio := map[string]int{
		"ONLINE":  0,
		"ALERTA":  0,
		"OFFLINE": 0,
	}
	for _, s := range servidores {
		if err := ValidadorServidor(s); err != nil {
			return nil, fmt.Errorf("erro no servidor %s: %w", s.Nome, err)
		}
		switch {
		case s.CPU > 80.0:
			relatorio["ALERTA"]++
		case s.Status == "STOPPED":
			relatorio["OFFLINE"]++
		default:
			relatorio["ONLINE"]++
		}
	}
	return relatorio, nil
}

func main() {

	fmt.Println("=== gtols v0.1 - Validador de Infraestrutura Fedora/Oracle ===")

	infra := []ServidorCloud{
		{ID: "srv-01", Nome: "Fedora Local", CPU: 45.2, Status: "RUNNING"},
		{ID: "srv-02", Nome: "Oracle VM 1", CPU: 88.5, Status: "RUNNING"},
		{ID: "srv-03", Nome: "Oracle VM 2", CPU: 12.0, Status: "STOPPED"},
	}

	relatorio, err := AnalisarInfra(infra)
	if err != nil {
		fmt.Println("Erro fatal ao analisar infraestrutura:", err)
		return
	}
	fmt.Println("--- Relatório Final de Saúde ---")
	for status, qtd := range relatorio {
		fmt.Printf("%-10s: %d\\n", status, qtd)
	}

}
