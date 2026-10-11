package main

import "fmt"


func main() {

	valueTotalCar := 25302.00

	fmt.Println("--- Opções de Parcelamento (Valor Mínimo da Parcela: R$ 1000) ---")

	for parcela := 1; parcela <= int(valueTotalCar); parcela++ {
		valueParcela := valueTotalCar / float64(parcela)
		
		if valueParcela <= 1000 {
			break //  Cancela as opções subsequentes pois o valor fica abaixo do mínimo permitido
		}
		fmt.Printf("Opção %2d parcelas de R$ %8.2f\n", parcela, valueParcela)
	}		
}
