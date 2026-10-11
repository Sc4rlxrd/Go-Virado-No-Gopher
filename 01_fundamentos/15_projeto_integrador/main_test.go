package main

import "testing"

func TestValidarServidorInvalido(t *testing.T) {
	s := ServidorCloud{ID: "", Nome: "Servidor Sem ID", CPU: 50.0}
	err := ValidadorServidor(s)
	if err == nil {
		t.Errorf("Esperava erro de ID obrigatório, mas obteve nil")
	}
}

func TestAnalisarInfra(t *testing.T) {
	infra := []ServidorCloud{
		{ID: "s1", Nome: "VM1", CPU: 90.0, Status: "RUNNING"},
		{ID: "s2", Nome: "VM2", CPU: 10.0, Status: "STOPPED"},
	}
	relatorio, err := AnalisarInfra(infra)
	if err != nil {
		t.Fatalf("Não esperava erro na análise, mas obteve: %v", err)
	}
	if relatorio["ALERTA"] != 1 {
		t.Errorf("Esperava 1 servidor em ALERTA, obteve %d", relatorio["ALERTA"])
	}
	if relatorio["OFFLINE"] != 1 {
		t.Errorf("Esperava 1 servidor OFFLINE, obteve %d", relatorio["OFFLINE"])
	}
}
