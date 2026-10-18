package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestMapear_Inteiros_Para_Strings(t *testing.T) {

	numeros := []int{1, 2, 3, 4, 5, 6, 7}

	resultado := Mapear(numeros, func(n int) string {
		return strconv.Itoa(n)
	})

	esperado := []string{"1", "2", "3", "4", "5", "6", "7"}

	if !slices.Equal(resultado, esperado) {
		t.Errorf("Mapear() = %v; esperado %v", resultado, esperado)
	}
}

func TestMapear_Decimais_Para_Strings(t *testing.T) {

	decimais:=[]float64{10.5,20.45,333.44}

	resultado := Mapear(decimais, func(n float64) string {
		return strconv.FormatFloat(n, 'f', -1, 64)  
	})

	esperado := []string{"10.5","20.45","333.44"}

	if !slices.Equal(resultado, esperado) {
		t.Errorf("Mapear() = %v; esperado %v", resultado, esperado)
	}
}

func TestMapear_Strings_Para_Inteiros(t *testing.T) {

	strings := []string{"1", "2", "3", "4", "5", "6", "7"}

	resultado := Mapear(strings, func(n string) int {
		valor, _ := strconv.Atoi(n)
		return valor
	})

	esperado := []int{1, 2, 3, 4, 5, 6, 7}

	if !slices.Equal(resultado, esperado) {
		t.Errorf("Mapear() = %v; esperado %v", resultado, esperado)
	}
}

func TestMapear_Strings_Para_Decimais(t *testing.T) {

	strings := []string{"10.5","20.45","333.44"}

	resultado:= Mapear(strings, func(n string) float64 {
		valor,_:= strconv.ParseFloat(n,64)
		return valor
	})

	esperado:=[]float64{10.5,20.45,333.44}

	if !slices.Equal(resultado, esperado) {
		t.Errorf("Mapear() = %v; esperado %v", resultado, esperado)
	}
}