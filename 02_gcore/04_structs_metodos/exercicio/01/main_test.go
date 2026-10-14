package main

import "testing"

func TestStatus(t *testing.T) {

	bd := BancoDeDados{
		Host:            "Void Linux-01",
		ConexoesAbertas: 2,
	}

	esperado := "Host: Void Linux-01 | Conexões ativas: 2"
	recebido := bd.Status()

	if recebido != esperado {
		t.Errorf(
			"status esperado: %q, recebido: %q",
			esperado,
			recebido)
	}

}

func TestIncrementarConexao(t *testing.T) {
	bd := BancoDeDados{
		Host:            "Void Linux-01",
		ConexoesAbertas: 0,
	}

	for i := 0; i < 4; i++ {
		bd.IncrementarConexao()
	}

	esperado := 4

	if bd.ConexoesAbertas != esperado {
		t.Errorf(
			"conexões esperadas: %d, recebidas: %d",
			esperado,
			bd.ConexoesAbertas,
		)
	}
}

func TestDecrementoConexao(t *testing.T) {
	bd := BancoDeDados{
		Host:            "Void Linux-01",
		ConexoesAbertas: 0,
	}

	for i := 0; i < 4; i++ {
		bd.IncrementarConexao()
	}

	esperado := 4

	if bd.ConexoesAbertas != esperado {
		t.Errorf(
			"conexões esperadas: %d, recebidas: %d",
			esperado,
			bd.ConexoesAbertas,
		)
	}
}
