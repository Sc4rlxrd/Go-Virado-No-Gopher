package main

import (
	"fmt"
)

// Numero 1. Constraint Customization usando União de Tipos (| e \~)
type Numero interface {
	~int | ~int64 | ~float32 | ~float64
}

// SomarValores 2. Função Genérica com Restrição Numérica
func SomarValores[T Numero](valores []T) T {
	var total T
	for _, v := range valores {
		total += v
	}
	return total
}

// RespostaAPI 3. Struct Genérica (Buffer/Cache Genérico)
type RespostaAPI[T any] struct {
	StatusCode int    `json: "status"`
	Mensagem   string `json: "mensagem"`
	Dados      T      `json: "dados"` // T pode ser uma Struct, Slice, Map, etc.
}

// 4. Função Genérica usando 'comparable' (necessário para mapas/chaves)

func RemoverDuplicados[T comparable](slice []T) []T {
	vistos := make(map[T]bool)
	resultados := []T{}

	for _, item := range slice {
		if !vistos[item] {
			vistos[item] = true
			resultados = append(resultados, item)
		}
	}
	return resultados
}

type Usuario struct {
	ID    int
	Nome  string
	Idade int
}

func main() {

	fmt.Println("===  Generics em Go  ===")

	// 1. Testando Soma Genérica com diferentes tipos numéricos
	inteiros := []int{10, 20, 30, 40}
	decimais := []float64{10.5, 20.25, 30.10, 45.10}

	fmt.Println("Soma dos inteiros: ", SomarValores(inteiros)) // Go infere o tipo [int] automaticamente!
	fmt.Println("Soma dos valores decimais com vírgulas: ", SomarValores(decimais))

	// 2. Testando Struct Genérica RespostaAPI[T]

	respUsuario := RespostaAPI[Usuario]{
		StatusCode: 200,
		Mensagem:   "Usuário encontrado no banco",
		Dados:      Usuario{ID: 1, Nome: "Scarlxrd", Idade: 22},
	}

	respLogs := RespostaAPI[[]string]{
		StatusCode: 500,
		Mensagem:   "Logs do contêiner docker",
		Dados:      []string{"err: Out of Memory", "err: Connection Refused"},
	}

	fmt.Println("Resposta 1: (Dados = Struct) ", respUsuario)
	fmt.Println("Resposta 2: (Dados = Slice) ", respLogs)

	// 3. Testando Remoção de Duplicados com 'comparable'

	ipDuplicados := []string{"192.168.1.1", "10.0.0.1", "192.168.1.1", "172.16.0.1"}
	ipUnicos := RemoverDuplicados(ipDuplicados)
	fmt.Println("Slice com todos os IPs duplicados ou não: ", ipDuplicados)
	fmt.Println("IPs Únicos sem duplicidades:  ", ipUnicos)

}
