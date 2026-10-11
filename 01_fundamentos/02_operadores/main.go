package main

import "fmt"

func main() {

	// Operadores Aritméticos
	// +, -, *, /, %

	a, b := 10, 3
	fmt.Println("--- Operadores Aritméticos ---")
	fmt.Printf("Soma (%d + %d): %d\n", a, b, a+b)

	fmt.Printf("Subtração (%d - %d): %d\n", a, b, a-b)

	fmt.Printf("Multiplicação (%d * %d): %d\n", a, b, a*b)

	fmt.Printf("Divisão Inteira (%d / %d): %d\n", a, b, a/b) // Corta as casas decimais

	fmt.Printf("Resto da Divisão / Módulo (%d %% %d): %d\n", a, b, a%b)

	// Divisão com ponto flutuante requer operandos de ponto flutuante
	var x, y float64 = 10.0, 3.0
	fmt.Printf("Divisão Flutuante (%.1f / %.1f): %.2f\n", x, y, x/y)

	// Operadores Relacionais
	// ==, !=, >, <, >=, <=

	fmt.Println("--- Operadores relacionais ---")
	fmt.Printf("Igual (%d == %d): %t\n", a, b, a == b)

	fmt.Printf("Diferente (%d != %d): %t\n", a, b, a != b)

	fmt.Printf("Maior (%d > %d): %t\n", a, b, a > b)

	fmt.Printf("Menor (%d < %d): %t\n", a, b, a < b)

	fmt.Printf("Maior ou Igual (%d >= %d): %t\n", a, b, a >= b)

	fmt.Printf("Menor ou Igual (%d <= %d): %t\n", a, b, a <= b)

	// Operadores Lógicos
	// &&, ||, !

	temhasSalary, hasEconomies, hasDebts := true, true, false
	fmt.Println("--- Operadores Lógicos ---")
	fmt.Printf("Comprar o carro (Salário e Economias): %t\n", temhasSalary && hasEconomies)
	fmt.Println("-")
	fmt.Printf("Comprar  itens (Salário ou Economias) %t\n", temhasSalary || hasEconomies)
	fmt.Println("-")
	fmt.Printf("Situação Regularizada (Não Tem Dívidas): %t\n", !hasDebts)

	// Operadores de Atribuição
	// =, +=, -=, *=, /=, %=

	counter := 10
	fmt.Println("--- Operadores de Atribuição ---")
	counter += 5 // Equivalente a: counter = counter + 5
	fmt.Printf("Após += 5: %d\n", counter)

	counter %= 4
	fmt.Printf("Após %%= 4: %d\n", counter)
	fmt.Println("-")

	// Operadores Unários  incremento E decremento
	// ++, --
	// Em Go, ++ e -- são INSTRUÇÕES (statements), NÃO expressões! [1, 2]

	i := 1
	i++ // Válido
	i-- // Válido
	fmt.Println("--- Incremento / Decremento ---")
	fmt.Printf("Valor final de i: %d\n", i)
	fmt.Println("-")

	// Operadores Bitwise / Bits
	// &, |, ^, <<, >>, &^

	p, q := 0b1100, 0b1010 // 12 e 10 em notação binária
	fmt.Println("--- Operadores Bitwise ---")
	fmt.Printf("AND (&):   %04b (%d)\n", p&q, p&q)
	fmt.Printf("OR (|):    %04b (%d)\n", p|q, p|q)
	fmt.Printf("XOR (^):   %04b (%d)\n", p^q, p^q)
	fmt.Printf("BIT CLEAR (&^): %04b (%d)\n", p&^q, p&^q) // Operador exclusivo do Go (AND NOT)
	fmt.Printf("Shift Left (<<):%04b (%d)\n", p<<1, p<<1) // Desloca os bits para a esquerda (multiplica por 2)
}
