package main

import (
	"fmt"
	"strconv"
)

func Mapear[T any, U any](slice []T, transformacao func(T) U) []U {
	resultado := make([]U, len(slice))

	for i, item := range slice {
		resultado[i] = transformacao(item) // Converte int para string
	}

	return resultado
}

func main() {

	fmt.Println(" === Ex 03 ===")
	numeros := []int{1, 2, 3, 4, 5, 6, 7} // origem 
	decimais:=[]float64{10.5,20.45,333.44}


	strings := Mapear(numeros, func(n int) string {
		// retornar uma string a partir de n
		return strconv.Itoa(n)
	})

	// TRANSFORMA STRINGS EM INTEIROS 
	inters := Mapear(strings, func(n string) int {
		valor, _ := strconv.Atoi(n)
		return valor
	})


	// TRANSFORMA FLOAT64 PARA STRINGS 
	stringsDecimais := Mapear(decimais, func(n float64) string {
		return strconv.FormatFloat(n, 'f', -1, 64)  // formata para manter a vŕrgula
	})


	// TRANSFORMA STRING PARA  FLOAT64
	textToDecimais:= Mapear(stringsDecimais, func(n string) float64 {
		valor,_:= strconv.ParseFloat(n,64)
		return valor
	})
	fmt.Printf("Tipo: %T | Valores: %v\n", strings, strings)
	fmt.Printf("Tipo: %T | Valores: %v\n", inters, inters)
	fmt.Printf("Tipo: %T | Valores: %v\n", stringsDecimais, stringsDecimais)
	fmt.Printf("Tipo: %T | Valores: %v\n", textToDecimais, textToDecimais)
	

}
