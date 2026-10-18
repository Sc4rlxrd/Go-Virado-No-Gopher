package main

import(
	"fmt"

	// Importando o pacote 'config' usando o caminho completo do módulo no go.mod
	cfg  "Go-Virado-No-Gopher/02_gcore/07_pacotes_modulos/config"
	"Go-Virado-No-Gopher/02_gcore/07_pacotes_modulos/internal/metrics"
)

// Função init() roda automaticamente ANTES da main()

func init(){
	fmt.Println("[INIT] Inicializando o sistema de telemetria no Fedora...") 
}


func main(){

	fmt.Println("===  Pacotes, Módulos e Visibilidade  ===")

	// 1. Usando o pacote com Alias 'cfg'
	stringBanco := cfg.ObterStringConexao("localhost")
	fmt.Println("String de conexão:", stringBanco) 
	fmt.Println("Porta Padrão:", cfg.PortaPadrao)

	// 2. Usando o pacote protegido dentro de 'internal/'
	coletor:= metrics.Coletor{}
	coletor.RegistrarUsoCPU(45.99)
}
