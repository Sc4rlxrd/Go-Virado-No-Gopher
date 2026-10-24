package main

import (
	"fmt"
	"strings"
	"sync"
)

type StatusServico struct {
	Nome       string
	Online     bool
	CodigoHTTP int
}

func ChecarHealth(nome string) StatusServico {

	if strings.HasPrefix(nome, "Auth",) || strings.HasPrefix(nome, "Payment") || strings.HasPrefix(nome, "Catalog") {
		return StatusServico{
			Nome:       nome,
			Online:     true,
			CodigoHTTP: 200,
		}
	}
	return StatusServico{
		Nome:       nome,
		Online:     false,
		CodigoHTTP: 500,
	}
}

func VerificarTodosHealthChecks(servicos []string) []StatusServico {
	var wg sync.WaitGroup

	resultados := make([]StatusServico, len(servicos))

	for indice, servico := range servicos {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resultados[indice] = ChecarHealth(u)

		}(servico)
	}
	wg.Wait()
	return resultados
}

func main() {
	fmt.Println(" --- Ex02 --- ")
	servico1 := ChecarHealth("Catalog-service")
	servico2 := ChecarHealth("Posgresql-service")
	fmt.Println("CHECARHEALTH DO SERVIÇOS SEM CONCORRENCIA:")
	fmt.Println("")
	fmt.Printf("O serviço:%v\nOnline:%v\nStatusCode:%v\n", servico1.Nome, servico1.Online, servico1.CodigoHTTP)
	fmt.Println("---")
	fmt.Printf("O serviço:%v\nOnline:%v\nStatusCode:%v\n", servico2.Nome, servico2.Online, servico2.CodigoHTTP)
	fmt.Println("")
	servicos := []string{
		"Auth-Serviço",
		"Payment-Serviço",
		"Catalog-Serviço",
		"Posgresql",
		"auth-serviço",
		"payment-serviço",
		"catalog-serviço",
	}

	fmt.Println("CHECARHEALTH DO SERVIÇOS COM CONCORRENCIA:")
	fmt.Println("")
	for _, servico := range VerificarTodosHealthChecks(servicos) {
		fmt.Printf("Serviço:%v\nOnline:%v\nStatusCode:%v\n", servico.Nome, servico.Online, servico.CodigoHTTP)
		fmt.Println("---")

	}

}
