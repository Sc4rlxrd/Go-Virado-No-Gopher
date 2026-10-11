package main

import "fmt"

func main() {

	var service string
	var door int
	var active bool

	fmt.Println("------  Zero Values (Valores Padrão)  ------")
	fmt.Printf("service (string): %q\\n", service) // "" (string vazia)
	fmt.Printf("door (int): %d\\n", door)          // 0
	fmt.Printf("active (bool): %t\\n", active)     // false

	// Declaração com inicialização e inferência de tipo de vida curta (operador :=)
	// Disponível apenas dentro de funções

	distro := "Fedora 44"
	versaoKernel := 6.14
	fmt.Println("\\n--- Inferência de Tipo (:=) ---")
	fmt.Printf("Distro: %s (Tipo: %T)\\n", distro, distro)
	fmt.Printf("Kernel: %.2f (Tipo: %T)\\n", versaoKernel, versaoKernel)

	// Conversão explícita de tipos (Obrigatória em Go)

	var memoriaMB int32 = 1024
	var memoriaGB float64 = float64(memoriaMB) / 1024.0

	fmt.Println("\\n--- Conversão Explícita de Tipos ---")
	fmt.Printf("Memória: %d MB (Tipo: %T)\\n", memoriaMB, memoriaMB)
	fmt.Printf("Memória: %.2f GB (Tipo: %T)\\n", memoriaGB, memoriaGB)

}
