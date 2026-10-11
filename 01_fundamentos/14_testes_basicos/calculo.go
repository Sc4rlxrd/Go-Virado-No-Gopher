package calculo

import "errors"

func Sun(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {

	return a - b
}

func Multiplication(a, b int) int {
	return a * b
}

func Division(a, b float64) float64 {
	return a / b
}

func ValidarIP(ip string) error {
	if ip == "" {
		return errors.New("endereço IP não pode ser vazio")
	}
	return nil
}
