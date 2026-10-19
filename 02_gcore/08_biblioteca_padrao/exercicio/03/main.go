package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	CaminhoArquivo string
}

func (l *Logger) RegistrarLog(mensagem string) error {
	timestamp := time.Now().Format(time.RFC3339)
	linha := fmt.Sprintf("[%s] %s\n", timestamp, mensagem)

	arquivo, err := os.OpenFile(l.CaminhoArquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	_, err = arquivo.WriteString(linha)
	return err
}

func main() {

	fmt.Println(" --- Ex 03 ---")
	meuLogger := &Logger{
		CaminhoArquivo: "app.log",
	}
	err := meuLogger.RegistrarLog("Aplicação iniciada com sucesso...")
	if err != nil {
		log.Fatalf("Erro ao registrar log: %v", err)
	}

	time.Sleep( 1  * time.Minute)

	err = meuLogger.RegistrarLog("Usuário realizou login no sistema.")
	if err != nil {
		log.Fatalf("Erro ao registrar log: %v", err)
	}

	fmt.Println("Logs gravados com sucesso no arquivo app.log!")

	_ = os.Remove(meuLogger.CaminhoArquivo)

}
