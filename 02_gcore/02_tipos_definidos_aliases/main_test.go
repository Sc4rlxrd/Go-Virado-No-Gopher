package main

import (
	"math"
	"testing"
)

func TestEmMaiusculas(t *testing.T) {

	nome := Nome("Notebook")
	resultado := nome.EmMaiusculas()
	esperado := "NOTEBOOK"
	if resultado != esperado {
		t.Errorf("EmMaiusculas() = %q; esperado %q", resultado, esperado)
	}
}

func TestPesoEmGramas(t *testing.T) {

	peso := Peso(1.75)
	resultado := peso.EmGramas()
	esperado := 1750.0
	if resultado != esperado {
		t.Errorf("EmGramas() = %v; esperado %v", resultado, esperado)
	}
}

func TestPrecoComDesconto(t *testing.T) {
	preco := Preco(3500.00)

	resultado := preco.ComDesconto(10)
	esperado := Preco(3150.00)

	if resultado != esperado {
		t.Errorf("ComDesconto() = %v; esperado %v", resultado, esperado)
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	resultado := CelsiusToFahrenheit(Celsius(32))
	esperado := Fahrenheit(89.6)

	if math.Abs(float64(resultado-esperado)) > 0.0001 {
		t.Errorf("CelsiusToFahrenheit() = %v; esperado %v", resultado, esperado)
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	resultado := FahrenheitToCelsius(Fahrenheit(110))
	esperado := Celsius(43.3333333333)

	if math.Abs(float64(resultado-esperado)) > 0.0001 { // math.Abs: transforma essa diferença em um valor não negativo.
		t.Errorf("FahrenheitToCelsius() = %v; esperado aproximadamente %v", resultado, esperado)
	}
}

func TestCalcularIMC(t *testing.T) {
	peso := Peso(85)
	altura := Altura(1.80)
	idade := Idade(22)

	imc, classificacao, err := CalcularIMC(peso, altura, idade)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	imcEsperado := 85.0 / (1.80 * 1.80)

	if imc != imcEsperado {
		t.Errorf("IMC esperado: %.2f, recebido: %.2f", imcEsperado, imc)
	}

	classificacaoEsperada := "Sobrepeso"

	if classificacao != classificacaoEsperada {
		t.Errorf(
			"classificação esperada: %s, recebida: %s",
			classificacaoEsperada,
			classificacao,
		)
	}
}
