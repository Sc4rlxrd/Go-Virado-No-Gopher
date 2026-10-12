package main

import "fmt"

func LimparCacheIPs(cache *[]string) {

	if cache == nil {
		fmt.Println("Erro: ponteiro para cache é nil")
		return
	}

	*cache = nil // Altera o slice original recebido na chamada da função.


	//  cache = nil -> Altera somente o ponteiro local dentro da própria função

}

func main() {

	IPs := []string{
		"214.11.156.108",
		"197.222.252.246",
		"52.5.113.164",
	}
	fmt.Println("---- Exercício 03 -----")
	fmt.Println("Todos IPs cadastrados: ", IPs)

	LimparCacheIPs(&IPs)
	fmt.Println("Cache limpo com sucesso:", IPs)
	fmt.Println("Cache é nil?", IPs == nil)

}
