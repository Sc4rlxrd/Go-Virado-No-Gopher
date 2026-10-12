package main

import (
	"fmt"
)

type APIRateLimiter struct {
	LimiteMaximo      int
	RequisicoesAtuais int
}

func (api *APIRateLimiter) RegistrarRequisicao() error {

	if api.RequisicoesAtuais >= api.LimiteMaximo {
		return fmt.Errorf("Limite Max alcançado")
	}

	api.RequisicoesAtuais++
	return nil

}
func main() {

	fmt.Println("---- Exercício 02 -----")
	api := APIRateLimiter{
		LimiteMaximo:      3,
		RequisicoesAtuais: 1,
	}

	fmt.Println(" --- Status da api atual ---")
	fmt.Printf("Limite Max: %d\nLimite de Requisições: %d\n", api.LimiteMaximo, api.RequisicoesAtuais)
	for i := 0; i < 4; i++ {
		err := api.RegistrarRequisicao()

		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Requisição registrada!")
		}
	}
	fmt.Printf("Limite Max: %d\nLimite de Requisições: %d\n", api.LimiteMaximo, api.RequisicoesAtuais)

}
