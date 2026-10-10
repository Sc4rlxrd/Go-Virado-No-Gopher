package main

import "fmt"

func main() {


	day := 9

	fmt.Println("--- Dia da Semana com Switch ---") 
	// Em Go, o 'break' é AUTOMÁTICO no final de cada 'case'. 
	// Não há vazamento de código (fallthrough) como no Java!

	switch day {
	case 1:
		println("Today is Sunday")
	case 2:
		println("Today is Monday")
	case 3:
		println("Today is Tuesday")
	case 4:
		println("Today is Wednesday")
	case 5:
		println("Today is Thursday")
	case 6:
		println("Today is Friday")
	case 7:
		println("Today is Saturday")
	default:
		println("Invalid option! Enter a day between 1 and 7.")
	}			
}
