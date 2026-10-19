package main

import (
	"encoding/json"
	"testing"
)

func TestSalvarConfiguracao(t *testing.T) {
	config := ConfiguracaoApp{
		NomeApp:     "Golang for beginners",
		MaxConexoes: 3,
		DebugMode:   false,
	}

	dadosJSON, err := SalvarConfiguracao(config)
	if err != nil {
		t.Fatalf("erro ao salvar configuração: %v", err)
	}

	var resultado map[string]interface{}
	if err := json.Unmarshal(dadosJSON, &resultado); err != nil {
		t.Fatalf("erro ao interpretar JSON gerado: %v", err)
	}

	if resultado["nome_app"] != "Golang for beginners" {
		t.Errorf("nome_app incorreto: %v", resultado["nome_app"])
	}

	if resultado["max_conexoes"] != float64(3) {
		t.Errorf("max_conexoes incorreto: %v", resultado["max_conexoes"])
	}

	if resultado["debug_mode"] != false {
		t.Errorf("debug_mode incorreto: %v", resultado["debug_mode"])
	}
}

func TestCarregarConfiguracao(t *testing.T) {
	conteudoJSON := []byte(`{
		"nome_app": "Golang for beginners",
		"max_conexoes": 3,
		"debug_mode": false
	}`)

	config, err := CarregarConfiguracao(conteudoJSON)
	if err != nil {
		t.Fatalf("erro ao carregar configuração: %v", err)
	}

	if config.NomeApp != "Golang for beginners" {
		t.Errorf("NomeApp incorreto: %v", config.NomeApp)
	}

	if config.MaxConexoes != 3 {
		t.Errorf("MaxConexoes incorreto: %v", config.MaxConexoes)
	}

	if config.DebugMode != false {
		t.Errorf("DebugMode incorreto: %v", config.DebugMode)
	}
}

func TestCarregarConfiguracaoJSONInvalido(t *testing.T) {
	conteudoJSON := []byte(`{"nome_app":`)

	_, err := CarregarConfiguracao(conteudoJSON)
	if err == nil {
		t.Fatal("esperava um erro para JSON inválido, mas nenhum erro foi retornado")
	}
}
