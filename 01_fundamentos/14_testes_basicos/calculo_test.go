package calculo

import "testing"

// O nome do teste deve começar com Test (com T maiúsculo)
func TestSun(t *testing.T) {
	result := Sun(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Somar(2, 3) = %d; esperado %d", result, expected)
	}
}

func TestValidarIPVazio(t *testing.T) {
	err := ValidarIP("")
	if err == nil {
		t.Errorf("Esperava erro para IP vazio, mas obteve nil")
	}
}

func TestSubtraction(t *testing.T) {

	result := Subtraction(5, 2)
	expected := 3

	if result != expected {
		t.Errorf("Subtração(5,2) = %d; esperado %d", result, expected)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(5, 2)
	expected := 10

	if result != expected {
		t.Errorf("Multiplication(5,2) = %d; esperado %d", result, expected)
	}
}

func TestDivision(t *testing.T) {

	result := Division(5, 2)
	expected := 2.5

	if result != expected {
		t.Errorf("Division(5,2) = %f; esperado %f", result, expected)
	}
}
