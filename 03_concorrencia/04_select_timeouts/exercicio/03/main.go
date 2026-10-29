package main

import "fmt"

// TentarEnviarMetrica tenta enviar sem bloquear.
// Retorna true se a métrica entrou no canal, false se foi descartada.
func TentarEnviarMetrica(chOut chan<- string, metrica string) bool {

	select {
	case chOut <- metrica:
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println(" --- Ex 03 --- ")

	fila := make(chan string, 3) // sem consumidor: enche na terceira métrica
	descartadas := 0
	for i := 1; i <= 5; i++ {
		m := fmt.Sprintf("cpu_uso=%d%%", 40+i)
		if TentarEnviarMetrica(fila, m) {
			fmt.Printf("enviada:    %s\n", m)
		} else {
			descartadas++
			fmt.Printf("descartada: %s\n", m)
		}
	}
	fmt.Printf("na fila: %d | descartadas: %d\n", len(fila), descartadas)
}
