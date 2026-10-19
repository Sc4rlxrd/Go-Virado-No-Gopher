package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRegistrarLog(t *testing.T) {

	diretorio := t.TempDir() // Cria um diretório temporário para o teste.
	caminho := filepath.Join(diretorio, "app.log")

	logger := &Logger{
		CaminhoArquivo: caminho,
	}

	mensagem1 := "Aplicação iniciada com sucesso"
	if err := logger.RegistrarLog(mensagem1); err != nil {
		t.Fatalf("erro ao gravar primeira mensagem: %v", err)
	}

	mensagem2 := "Usuário realizou login no sistema"
	if err := logger.RegistrarLog(mensagem2); err != nil {
		t.Fatalf("erro ao gravar segunda mensagem: %v", err)
	}

	conteudo, err := os.ReadFile(caminho)
	if err != nil {
		t.Fatalf("erro ao ler arquivo de log: %v", err)
	}

	texto := string(conteudo)

	if !strings.Contains(texto, mensagem1) {
		t.Errorf("arquivo não contém a primeira mensagem: %s", mensagem1)
	}

	if !strings.Contains(texto, mensagem2) {
		t.Errorf("arquivo não contém a segunda mensagem: %s", mensagem2)
	}
}
