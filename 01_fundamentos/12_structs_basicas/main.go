package main

import "fmt"

// Declaração de um tipo composto heterogêneo
type Servidor struct {
	IP       string
	Porta    int
	Ativo    bool
	Servicos []string
}

// Função que aceita struct como parâmetro (passagem por valor)

func exibirStatusServidor(s Servidor) {
	status := "INATIVO"
	if s.Ativo {
		status = "ATIVO"
	}
	fmt.Printf("[%s] Servidor %s:%d | Serviços: %v\n", status, s.IP, s.Porta, s.Servicos)
}

func main() {

	//  Instanciação com Zero Value (todos os campos iniciam com os zero values de seus tipos)
	var s1 Servidor
	fmt.Println("--- Struct com Zero Value ---")
	fmt.Printf("s1: %+v\n", s1) // IP="", Porta=0, Ativo=false, Servicos=nil

	//  Instanciação com Composite Literal (nomeando os campos explicitamente)

	s2 := Servidor{
		IP:       "192.186.1.100",
		Porta:    9090,
		Ativo:    true,
		Servicos: []string{"HTTP", "Metrics"},
	}
	fmt.Println("--- Struct Inicializada ---")
	exibirStatusServidor(s2)

	//  Leitura e alteração individual de campos
	s2.Porta = 8443
	fmt.Println("Porta atualizada:", s2.Porta)
	fmt.Println("Dados Atualizados: ", s2)

}
