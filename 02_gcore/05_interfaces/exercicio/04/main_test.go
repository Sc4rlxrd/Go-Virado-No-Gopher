package main

import "testing"

func TestExibirRegiao(t *testing.T) {

	resultado := AWSProvedor{
		Provider: "AWS",
		Region:   "Ohio",
	}
	ExibirRegiao(resultado)

	esperado := AWSProvedor{
		Provider: "AWS",
		Region:   "Ohio",
	}
	ExibirRegiao(esperado)

	if resultado != esperado {
		t.Error("Erro")
	}

}

func TestAWSProvedor_ObterInformacoes(t *testing.T) {
	aws := AWSProvedor{
		Provider: "AWS",
		Region:   "Ohio",
	}

	info := aws.ObterInformacoes()

	if info.Provider != "AWS" {
		t.Errorf("provider esperado %q, recebido %q", "AWS", info.Provider)
	}

	if info.Regiao != "Ohio" {
		t.Errorf("região esperada %q, recebida %q", "Ohio", info.Regiao)
	}
}


func TestOracleProvedor_ObterInformacoes(t *testing.T) {
	oracle := OracleProvedor{
		Provider: "ORACLE CLOUD",
		Location: "sa-saopaulo-1",
	}

	info := oracle.ObterInformacoes()

	if info.Provider != "ORACLE CLOUD" {
		t.Errorf("provider esperado %q, recebido %q", "ORACLE CLOUD", info.Provider)
	}

	if info.Regiao != "sa-saopaulo-1" {
		t.Errorf("região esperada %q, recebida %q", "sa-saopaulo-1", info.Regiao)
	}
}

func TestAzureProvedor_ObterInformacoes(t *testing.T) {
	azure := AzureProvedor{
		Provider: "AZURE CLOUD",
		Location: "Texas",
	}

	info := azure.ObterInformacoes()

	if info.Provider != "AZURE CLOUD" {
		t.Errorf("provider esperado %q, recebido %q", "AZURE CLOUD", info.Provider)
	}

	if info.Regiao != "Texas" {
		t.Errorf("região esperada %q, recebida %q", "Texas", info.Regiao)
	}
}

func TestGoogleCloud_ObterInformacoes(t *testing.T) {
	googlecloud := GoogleCloud{
		Provider: "GOOGLE CLOUD",
		Region:   "BRAZIL-SÃOPAULO - 01",
	}

	info := googlecloud.ObterInformacoes()

	if info.Provider != "GOOGLE CLOUD" {
		t.Errorf("provider esperado %q, recebido %q", "GOOGLE CLOUD", info.Provider)
	}

	if info.Regiao != "BRAZIL-SÃOPAULO - 01" {
		t.Errorf("região esperada %q, recebida %q", "BRAZIL-SÃOPAULO - 01", info.Regiao)
	}
}
