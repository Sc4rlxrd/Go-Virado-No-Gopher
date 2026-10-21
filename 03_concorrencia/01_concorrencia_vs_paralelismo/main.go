package main

import (
	"fmt"
	"runtime"
	"time"
)

// função simples que simula uma tarefa em segundo plano
func tarefaSegundoPlano(nome string, interacoes int) {
	for i := 0; i <= interacoes; i++ {
		fmt.Printf("[%s] Executando passo %d/%d\n", nome, i, interacoes)
		// Simula um pequeno tempo de I/O ou processamento
		time.Sleep(10 * time.Second)
	}
	fmt.Printf("✅ [%s] Finalizada!\n", nome)
}

func main() {

	fmt.Println("=== ⚡ Aula 01: Concorrência vs. Paralelismo (Go Runtime) ===")
	// 1. Inspecionando o ambiente de execução do Go Runtime no Fedora
	cpus := runtime.NumCPU()
	maxProcs := runtime.GOMAXPROCS(0) // 0 apenas consulta o valor atual
	goroutinesIniciais := runtime.NumGoroutine()

	fmt.Printf("Núcleos de CPU Lógicos : %v\n", cpus)
	fmt.Printf("GOMAXPROCS Ativo : %d (Processadores lógicos P)\n", maxProcs)
	fmt.Printf("Goroutines Ativas : %d (Apenas a main)\n", goroutinesIniciais)

	// 2. Disparando uma Goroutine com a palavra-chave 'go'
	// O prefixo 'go' instrui a runtime a executar a função de forma assíncrona
	go tarefaSegundoPlano("Coletor-Background", 3)

	// Inspecionando as Goroutines vivas logo após o disparo
	fmt.Printf("Goroutines ativas após o 'go': %d\n", runtime.NumGoroutine())

	// Executando a tarefa principal na goroutine main
	tarefaSegundoPlano("Main-Thread", 3)

	// ⚠️ ATENÇÃO E ARMADILHA COMUM:
	//  Se a função main() encerrar, o processo inteiro morre imediatamente,
	//  abortando qualquer goroutine filha que ainda esteja rodando!
	//  Nas próximas aulas, usaremos canais e WaitGroups em vez de time.Sleep.

	time.Sleep(50 * time.Second)
	fmt.Println("Programa principal finalizado!")
}
