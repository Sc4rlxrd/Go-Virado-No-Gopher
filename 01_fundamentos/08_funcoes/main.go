package main

import "fmt"

//  Função simples com parâmetros e retorno único
func sun(a int, b int) int {
	return a + b
}

//  Parâmetros do mesmo tipo podem compartilhar a declaração do tipo (a, b int)
// Go suporta Múltiplos Retornos nativamente.

func divideAndCalculateRemainder(dividend, divider int) (int, int) {
	quotient := dividend / divider
	rest := dividend % divider
	return quotient, rest
}

//  Retornos Nomeados (Named Return Values): As variáveis de retorno são pré-inicializadas com zero value
func calculateServerStatistics(cpuUsage float64) (status string, alert bool) {

	if cpuUsage > 85.0 {
		status = "CRÍTICO"
		alert = true
		return // Return limpo (naked return)
	}
	status = "ESTÁVEL"
	alert = false
	return
}

//  Função Variádica: Aceita um número variável de argumentos do mesmo tipo (...int)
// Internamente, 'numeros' é tratado como um Slice ([]int)
func sumAll(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	
	fmt.Println("--- Funções Básicas e Múltiplos Retornos ---")

	sumResult := sun(10, 5)
	fmt.Println("Soma:", sumResult)

	// Capturando múltiplos retornos [1, 2]
	quotient, remainder := divideAndCalculateRemainder(10, 3)
	fmt.Printf("Divisão 10/3 -> Quociente: %d, Resto: %d\n", quotient, remainder)

	// Descartando um retorno indesejado usando o Blank Identifier (_) [3]
	_,remainderOnly:= divideAndCalculateRemainder(15, 4)
	fmt.Printf("Resto da divisão 15/4: %d\n", remainderOnly)

	// Retornos nomeados
	st, alt := calculateServerStatistics(80.5)
	fmt.Printf("Servidor Fedora -> Status: %s | Alerta: %t\n", st, alt)

	// Função Variádica
	fmt.Println("--- Funções Variádicas ---")
	all1:= sumAll(10, 20, 30, 40, 50)
	all2:= sumAll(5, 15, 25)
	all3:= sumAll() 
	fmt.Println("Total 1: ", all1)
	fmt.Println("Total 2: ", all2)
	// Chamando sem argumentos
	fmt.Println("Total 3 (sem argumentos): ", all3)

	// Expandindo um Slice existente para passar para a função variádica (...)

	myNumbers := []int{5, 10, 15,44,444,44,4,4,4,4,4,4,44,4}
	totalSlice:= sumAll(myNumbers...)
	fmt.Println("Total do Slice expandido: ", totalSlice)

}
