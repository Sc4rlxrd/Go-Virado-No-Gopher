package main

import "fmt"

func ClassificarStatusHTTP(code int) string {

	switch {
	case code <= 299 && code >= 200:
		{
			return "Sucesso"
		}
	case code <= 499 && code >= 400:
		{
			return "Cliente Erro"
		}
	case code <= 599 && code >= 500:
		{
			return "Server error"
		}
	default:
		return "Desconhecido"

	}
}

func main() {
	fmt.Println(" --- Ex 01 ---")

	fmt.Printf("Status HTTP: %v\n",ClassificarStatusHTTP(450))
}
