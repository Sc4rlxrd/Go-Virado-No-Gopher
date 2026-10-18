package main

import (
	"fmt"

	val "Go-Virado-No-Gopher/02_gcore/07_pacotes_modulos/exercicio/02/validator"
)

func main() {

	fmt.Println("--- Ex02 ---")

	// IPs Válidos (devem retornar true / passar)
	val.ValidarIP("47.70.50.105")
	val.ValidarIP("223.159.1.225")

	// IPs Inválidos (devem retornar false / cair no net.ParseIP == nil)
	val.ValidarIP("999.999.999.999") // Números fora do limite
	val.ValidarIP("abc.def.ghi.jkl") // Letras
	val.ValidarIP("192.168.1")       // Incompleto

}
