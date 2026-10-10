package main

import "fmt"

func main() {

	salary := 7000.00
	var result string

	// 1. Forma tradicional em Go (Substituindo o Ternário do Java)
	//Em Java seria: status = (salario > 5000) ? "Vou doar 500" : "Ainda não posso doar";

	if salary >= 5000 {
		result = "I'm going to donate 500 to the Go-Virado-No-Gopher project."
	} else {
		result = "I can't donate yet, but I'll keep studying!"
	}

	fmt.Println("Standard Form:", result)

	// O Idiomatismo Go: 'if' com Instrução de Inicialização (Short Statement)
	// A variável 'bonus' é declarada, inicializada e avaliada na MESMA LINHA.
	//O escopo da variável 'bonus' fica RESTRITO ao bloco do 'if/else'.

	fmt.Println("--- Idiomatismo: 'if' com Short Statement ---")

	if bonus := salary * 0.10; bonus > 500 {
		fmt.Printf("Bonus: R$%.2f  aprovado para investimento!\n", bonus)
	} else {
		fmt.Printf("Bônus de R$ %.2f mantido em reserva.\n", bonus)
	}

	// Descomentar a linha abaixo geraria ERRO de compilação:   undefined: bonus (fora de escopo!)
	// fmt.Println(bonus)
}
