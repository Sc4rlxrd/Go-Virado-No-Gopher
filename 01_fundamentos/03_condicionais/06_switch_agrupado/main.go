package main

import "fmt"

func main() {

	day := 7

		fmt.Println("--- Validação de Dia Útil vs Fim de Semana ---") 
		//  Agrupamento de valores em um único 'case' separados por vírgula
		switch day {
		case 1,7:
			fmt.Println("Weekend! Time to rest or focus on personal projects.")
		case 2, 3, 4, 5, 6:
			fmt.Println("Workday. Work and study routine.")
		default:
			fmt.Println("Invalid option! Enter a day between 1 and 7.")
		}


		//  Switch Sem Expressão (Tagless Switch) 
		// Funciona como uma cadeia limpa de 'if / else if / else'
		fmt.Println("")
		fmt.Println("--- Switch Sem Expressão (Tagless Switch) ---")
		studentAverage:= 9
		switch {
		case studentAverage >= 9:
			fmt.Println("Excellent! Keep up the good work. notes: " + fmt.Sprint(studentAverage))
		case studentAverage >= 7 && studentAverage < 9:
			fmt.Println("Good! But there's room for improvement. notes: " + fmt.Sprint(studentAverage))
		case studentAverage >= 5 && studentAverage < 7:
			fmt.Println("You need to study more to improve your grades. notes: " + fmt.Sprint(studentAverage))
		default:
			fmt.Println("You need to dedicate more time to your studies.")
		}
}

