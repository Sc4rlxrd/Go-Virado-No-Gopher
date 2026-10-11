package main

import "fmt"

func main() {

	valueTotalCar := 25302.00

	fmt.Println("--- Filtrando opções de parcelas a partir da parcela 12 ---")

	for parcela := int(valueTotalCar); parcela >= 1; parcela-- {
		valueParcela := valueTotalCar / float64(parcela)

		if valueParcela <= 3000 {
			continue // Ignora parcelas com valor menor que R$ 1000 e salta para a próxima iteração
		}
		fmt.Printf("Opção %2d parcelas de R$ %8.2f\n", parcela, valueParcela)
	}
}
