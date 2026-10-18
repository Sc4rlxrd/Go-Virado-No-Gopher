package main

import (
	"fmt"
	"os"
	"strconv"
)

func ObterPortaServidor() (int, error) {
	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		return 8080, nil
	}

	return strconv.Atoi(appPort)
}

func main() {
	fmt.Println("--- Ex01 ---")

	porta, err := ObterPortaServidor()
	if err != nil {
		fmt.Println("Erro ao obter a porta:", err)
		return
	}

	fmt.Printf("Porta do servidor: %d\n", porta)
}
