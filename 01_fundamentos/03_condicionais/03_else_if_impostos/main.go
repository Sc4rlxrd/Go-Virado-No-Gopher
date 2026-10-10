package main

import "fmt"

func main() {

	annualSalary := 45000.00
	var firstTaxRate float64 = 9.70 / 100
	var secondTaxRate float64 = 37.35 / 100
	var thirdTaxRate float64 = 49.50 / 100

	var tax float64

	fmt.Println("--- Cálculo de Imposto de Renda (Alíquota Progressiva) ---")

	if annualSalary <= 34712 {
		tax = annualSalary * firstTaxRate
	} else if annualSalary > 34712 && annualSalary <= 68507 {
		tax = annualSalary * secondTaxRate
	} else {
		tax = annualSalary * thirdTaxRate
	}

	netSalary := annualSalary - tax

	fmt.Printf("Salário Bruto Anual: R$ %.2f\n", annualSalary)
	fmt.Printf("Valor do Imposto: R$ %.2f\n", tax)
	fmt.Printf("Salário Líquido Anual: R$ %.2f\n", netSalary)
}
