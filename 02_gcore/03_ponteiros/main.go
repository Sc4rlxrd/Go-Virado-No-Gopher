package main

import (
	"fmt"
)

// Struct representando a configuração de um servidor na Oracle Cloud / Fedora
type ServidorCloud struct {
	Hostname string
	Porta    int
	Ativo    bool
}


// Tenta alterar por VALOR (Recebe uma CÓPIA do ServidorConfig)
func tentarAtivarPorValor(s ServidorCloud) {
	s.Ativo = true
}

// Altera por PONTEIRO (Recebe o endereço da struct original em *ServidorConfig)
func ativarServidorReal(s *ServidorCloud) {
	if s == nil {
		fmt.Println("Erro: Ponteiro nulo (nil) recebido!")
		return
	}
	// Atalho sintático do Go: s.Ativo é o mesmo que (*s).Ativo
	s.Ativo = true
}

func main() {

	fmt.Println("=== 📍 Ponteiros em Go  ===")

	//  Variável primitiva simples
	portaHTTP := 8080
	prtPorta := &portaHTTP // ptrPorta é do tipo *int

	fmt.Printf("Valor de portaHTTP: %d\n", portaHTTP)
	fmt.Printf("Endereço de portaHTTP: %p\n", &portaHTTP)
	fmt.Printf("Conteúdo de ptrPorta: %p\n", prtPorta)
	fmt.Printf("Valor via desreferenciação (*ptrPorta): %d\n", *prtPorta)

	// Alterando portaHTTP diretamente pelo ponteiro

	*prtPorta = 9090
	fmt.Printf("Novo valor de portaHTTP pós *ptrPorta = 9090: %d\n", portaHTTP)

	//  Modificação de Structs em funções (Valor vs Ponteiro)

	srv := ServidorCloud{
		Hostname: "fedora-server-001",
		Porta:    433,
		Ativo:    false,
	}

	fmt.Println("--- Teste 1: Passagem por Valor ---")

	tentarAtivarPorValor(srv)
	fmt.Printf("Servidor ativo após tentativa por valor? %t (Não mudou!)\n", srv.Ativo)

	fmt.Println("--- Teste 2: Passagem por Ponteiro ---")
	ativarServidorReal(&srv) // Passando o endereço de memória com &;
	fmt.Printf("Servidor ativo após atualização por ponteiro? %t (Atualizado!)\n", srv.Ativo)

	//  Checagem de segurança para ponteiros 'nil'
	var srvNulo *ServidorCloud // Zero value de um ponteiro é 'nil'
	fmt.Println("-- Teste 3: Proteção contra Nil Pointer ---")
	ativarServidorReal(srvNulo)



}
