package main

import (
	"errors"
	"fmt"
)

type Armazenamento interface {
	Buscar(chave string) (any, bool)
}

type MemoriaCache struct {
	dados map[string]any
}

func (m *MemoriaCache) Buscar(chave string) (any, bool) {
	valor, ok := m.dados[chave]

	return valor, ok
}

func BuscarComoString(a Armazenamento, chave string) (string, any) {

	valor, encontrado := a.Buscar(chave)
	if !encontrado {
		return "", errors.New("chave não encontrada no cache")
	}

	str, ok := valor.(string)
	if !ok {
		return "", errors.New("o valor armazenado não é uma string")
	}

	return str, nil

}

func main() {

	fmt.Println("---- Exercício 02 -----")

	cache:= &MemoriaCache{
		dados: map[string]any{
			"nome":"Golang",
			"idade":10,
		},
	}

	resultado,err := BuscarComoString(cache,"nome")
	if err != nil{
		fmt.Println("Erro:", err)
	}else{
		fmt.Println("Sucesso: ", resultado)
	}


	resultado2, err := BuscarComoString(cache, "idade")
	if err != nil {
		fmt.Println("Erro esperado:", err)
	} else {
		fmt.Println("Sucesso:", resultado2)
	}

}
