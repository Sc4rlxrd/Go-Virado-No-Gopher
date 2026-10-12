package  main

import(
	"fmt"
)

func ReiniciarPorta(porta *int) {
	if porta == nil {
		fmt.Println("Erro: Ponteiro nulo (nil) recebido!")
		return
	}
	*porta = 8080

}

func main(){

	fmt.Println("---- Exercício 01 -----")
	porta := 3000
	fmt.Println("Antes:", porta)
	ReiniciarPorta(&porta)
	fmt.Println("Depois:", porta)
}
