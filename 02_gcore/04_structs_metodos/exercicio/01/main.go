package main

import (
	"fmt"
)

type BancoDeDados struct {
	Host            string
	ConexoesAbertas int
}

func (bd BancoDeDados) Status() string {
	return fmt.Sprintf(
		"Host: %s | Conexões ativas: %d",
		bd.Host,
		bd.ConexoesAbertas,
	)
}

func (bd *BancoDeDados) IncrementarConexao() {

	bd.ConexoesAbertas++
	fmt.Printf("Poll de conexões: %d\n", bd.ConexoesAbertas)

}

func (bd *BancoDeDados) DecrementoConexao() {

	bd.ConexoesAbertas--
	fmt.Printf("Poll de conexões: %d\n", bd.ConexoesAbertas)
}

func main() {

	fmt.Println("---- Exercício 01 -----")
	srv := BancoDeDados{
		Host:            "Void Linux-01",
		ConexoesAbertas: 0,
	}
	fmt.Println("Servidor base: ", srv)
	for i := 0; i < 4; i++ {
		srv.IncrementarConexao()
	}
	fmt.Println(" --- Incrementando conexões ---")
	fmt.Println("Status: ", srv.Status())
	fmt.Println(" --- Decrementando conexões ---")
	for i := 0; i < 2; i++ {
		srv.DecrementoConexao()
	}
	fmt.Println("Status: ", srv.Status())

}
