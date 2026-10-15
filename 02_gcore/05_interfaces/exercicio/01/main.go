package main

import (
	"fmt"
)

type ProvedorCloud interface {
	ObterRegiao() string
	ObterProvider() string
}

type AWSProvedor struct {
	Provider string
	Region   string
}

type OracleProvedor struct {
	Provider string
	Location string
}

func (aws AWSProvedor) ObterRegiao() string {

	return aws.Region
}

func (oracle OracleProvedor) ObterRegiao() string {
	return oracle.Location
}

func (aws AWSProvedor) ObterProvider() string {

	return aws.Provider
}

func (oracle OracleProvedor) ObterProvider() string {
	return oracle.Provider
}

func ExibirRegiao(p ProvedorCloud) {
	fmt.Printf("Provider: %s\n", p.ObterProvider())
	fmt.Printf("Localidade: %s\n", p.ObterRegiao())
	fmt.Println()
}

func main() {

	fmt.Println("---- Exercício 01 -----")

	aws := AWSProvedor{
		Provider: "AWS",
		Region:   "Ohio",
	}
	oracle := OracleProvedor{
		Provider: "ORACLE CLOUD",
		Location: "sa-saopaulo-1",
	}

	// fmt.Println("Aws: ", aws)
	// fmt.Println("Oracle: ", oracle)
	// aws.ObterRegiao()
	// oracle.ObterRegiao()
	ExibirRegiao(aws)
	ExibirRegiao(oracle)
}
