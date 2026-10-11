package main

import (
	"fmt"
	"math"
)

// REGRA DE VISIBILIDADE EM GO (Exported vs. Unexported)
//  Identificador que começa com Letra MAIÚSCULA = Exportado (Público / Visível fora do pacote)
//  Identificador que começa com Letra MINÚSCULA = Não-Exportado (Privado ao pacote)

// SystemName é uma constante EXPORTADA (pública)
const SystemName = "Go-Virado-No-Gopher v1.0"

// versaoInterna é uma constante NÃO-EXPORTADA (privada ao pacote main)
const versaoInterna = "1.0.0-build.44"

// ComputarAreaCirculo é uma função EXPORTADA (pública)
func ComputarAreaCirculo(raio float64) float64 {
	return math.Pi * math.Pow(raio, 2)
}

// formatarSaida é uma função NÃO-EXPORTADA (privada)
func formatarSaida(mensagem string) string {
	return fmt.Sprintf("[LOG INTERNO]: %s", mensagem)
}
func main() {

	fmt.Println("--- Regras de Visibilidade e Pacotes ---")
	fmt.Println("Constante Pública (Exported):", SystemName)
	fmt.Println("Constante Privada (Unexported):", versaoInterna)

	area := ComputarAreaCirculo(5.0)
	fmt.Printf("Área do Círculo (via math.Pi): %.2f\n", area)

	msgFormatada := formatarSaida("Ambiente Fedora Linux 44 pronto!")
	fmt.Println(msgFormatada)

}
