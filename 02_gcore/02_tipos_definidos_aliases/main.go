package main

import (
	"fmt"
	"strings"
)

// Tipos definidos a partir de tipos existentes.
type Nome string
type Peso float64
type Preco float64
type Celsius float64
type Fahrenheit float64
type Altura float64
type Idade int

func CelsiusToFahrenheit(celsius Celsius) Fahrenheit {

	return Fahrenheit(float64(celsius)*9.0/5.0 + 32)

}

func FahrenheitToCelsius(fahrenheit Fahrenheit) Celsius {

	return Celsius((float64(fahrenheit) - 32) * 5.0 / 9.0)

}

func CalcularIMC(peso Peso, altura Altura, idade Idade) (float64, string, error) {

	if peso <= 0 {
		return 0, "", fmt.Errorf("o peso deve ser maior que zero")
	}

	if altura <= 0 {
		return 0, "", fmt.Errorf("a altura deve ser maior que zero")
	}

	if idade <= 0 {
		return 0, "", fmt.Errorf("a idade deve ser maior que zero")
	}

	imc := float64(peso) / (float64(altura) * float64(altura))

	switch {
	case idade < 20:
		return imc, "A classificação precisa considerar a idade e o sexo", nil
	case imc < 18.5:
		return imc, "Abaixo do peso", nil
	case imc < 25.0:
		return imc, "Peso normal", nil
	case imc < 30.0:
		return imc, "Sobrepeso", nil
	default:
		return imc, "Obesidade", nil
	}
}

// Métodos associados aos tipos definidos.
func (n Nome) EmMaiusculas() string {
	return strings.ToUpper(string(n))
}

func (p Peso) EmGramas() float64 {
	return float64(p * 1000)
}

func (p Preco) ComDesconto(percentual float64) Preco {
	return Preco(float64(p) * (1 - percentual/100))
}

func main() {
	// Criando valores dos tipos personalizados.
	var nome Nome = "Notebook"
	var peso Peso = 1.75
	var preco Preco = 3500.00

	fmt.Println("----- Tipos definidos -----")
	fmt.Println("Nome:", nome)
	fmt.Println("Peso:", peso, "kg")
	fmt.Printf("Preço: R$ %.2f\n", preco)

	fmt.Println("----- Utilizando métodos -----")
	fmt.Println("Nome em maiúsculas:", nome.EmMaiusculas())
	fmt.Printf("Peso em gramas: %.0f g\n", peso.EmGramas())

	precoComDesconto := preco.ComDesconto(10)
	fmt.Printf("Preço com 10%% de desconto: R$ %.2f\n", precoComDesconto)

	fmt.Println("----- Conversão entre tipos -----")
	nomeOriginal := string(nome)
	fmt.Println("Nome como string:", nomeOriginal)

	pesoOriginal := float64(peso)
	fmt.Println("Peso como float64:", pesoOriginal)

	fmt.Println("----- Comparando tipos -----")
	fmt.Println("Preço original:", preco)
	fmt.Println("Preço com desconto:", precoComDesconto)

	fmt.Println("---- Conversão de Celsius para Fahrenheit ----")
	fahrenheit := CelsiusToFahrenheit(32)
	fmt.Printf("32 °C = %.1f °F\n", fahrenheit)
	fmt.Println("---- Conversão de Fahrenheit para Celsius ----")
	celsius := FahrenheitToCelsius(110)
	fmt.Printf("110 °F = %.1f °C\n", celsius)

	pesoPessoa := Peso(85)
	alturaPessoa := Altura(1.80)
	idadePessoa := Idade(22)
	nomePessoa := Nome("Scar").EmMaiusculas()

	imc, classificacao, err := CalcularIMC(pesoPessoa, alturaPessoa, idadePessoa)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("---- Calculadora de IMC ----")
	fmt.Println("Nome: ", nomePessoa)
	fmt.Printf("Idade: %d anos\n", idadePessoa)
	fmt.Printf("IMC: %.2f\n", imc)
	fmt.Println("Classificação:", classificacao)

}
