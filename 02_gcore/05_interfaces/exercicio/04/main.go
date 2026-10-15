package main

import (
	"fmt"
)

type ProvedorCloud interface {
	ObterInformacoes() InformacoesProvedor
}

type InformacoesProvedor struct {
	Provider string
	Regiao   string
}

type AWSProvedor struct {
	Provider string
	Region   string
}

type OracleProvedor struct {
	Provider string
	Location string
}

type AzureProvedor struct {
	Provider string
	Location string
}

type GoogleCloud struct {
	Provider string
	Region   string
}

func (aws AWSProvedor) ObterInformacoes() InformacoesProvedor {
	return InformacoesProvedor{
		Provider: aws.Provider,
		Regiao:   aws.Region,
	}
}

func (oracle OracleProvedor) ObterInformacoes() InformacoesProvedor {
	return InformacoesProvedor{
		Provider: oracle.Provider,
		Regiao:   oracle.Location,
	}
}

func (azure AzureProvedor) ObterInformacoes() InformacoesProvedor {
	return InformacoesProvedor{
		Provider: azure.Provider,
		Regiao:   azure.Location,
	}
}

func (googlecloud GoogleCloud) ObterInformacoes() InformacoesProvedor {
	return InformacoesProvedor{
		Provider: googlecloud.Provider,
		Regiao:   googlecloud.Region,
	}
}

func ExibirRegiao(p ProvedorCloud) {
	info := p.ObterInformacoes()

	fmt.Printf("Provider: %s\n", info.Provider)
	fmt.Printf("Localidade: %s\n\n", info.Regiao)
}

func main() {

	fmt.Println("---- Exercício 04 -----")

	aws := AWSProvedor{
		Provider: "AWS",
		Region:   "Ohio",
	}
	oracle := OracleProvedor{
		Provider: "ORACLE CLOUD",
		Location: "sa-saopaulo-1",
	}

	azure := AzureProvedor{
		Provider: "AZURE CLOUD",
		Location: "Texas",
	}

	googlecloud := GoogleCloud{
		Provider: "GOOGLE CLOUD",
		Region:   "BRAZIL-SÃOPAULO01",
	}

	// fmt.Println("Aws: ", aws)
	// fmt.Println("Oracle: ", oracle)
	// aws.ObterRegiao()
	// oracle.ObterRegiao()
	ExibirRegiao(aws)
	ExibirRegiao(oracle)
	ExibirRegiao(azure)
	ExibirRegiao(googlecloud)
}
