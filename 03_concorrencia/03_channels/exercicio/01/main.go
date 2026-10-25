package main

import (
	"fmt"
)

func GerarLogs(mensagens []string, chOut chan<- string) {
	defer close(chOut)

	for _, msg := range mensagens {
		chOut <- msg
	}
}

func ConsumirLogs(chIn <-chan string) []string {

	lidas := []string{}
	for msg := range chIn {
		lidas = append(lidas, "[LOG_PROCESSADO] "+msg)
	}
	return lidas

}

func main() {
	fmt.Println(" --- Ex 01 --- ")

	mensagens := []string{"Amazon Sucesso", "MercadoLivre Sucesso ", "Kabum Falha", "Magalu Sucesso", "Boa Dica Sucesso"}

	canal := make(chan string, 3)
	go GerarLogs(mensagens, canal)

	for _, log := range ConsumirLogs(canal) {
		fmt.Println(log)
	}

}
