package main

import (
	"fmt"
	"sync"
	"time"
)

type TarefaColeta struct {
	ID      int
	Produto string
}

type ResultadoColeta struct {
	Tarefa TarefaColeta
	Preco  float64
	Worker int
}

// Worker: Função executada por N goroutines em paralelo
func workerColetor(idWorker int, jobs <-chan TarefaColeta, results chan<- ResultadoColeta, wg *sync.WaitGroup) {
	defer wg.Done() // Garante que o WaitGroup seja decrementado ao final da execução da função

	for job := range jobs {
		// Simula o trabalho de scaping e coleta de preço do produto
		time.Sleep(15 * time.Millisecond)
		results <- ResultadoColeta{
			Tarefa: job,
			Preco:  99.90 + float64(job.ID),
			Worker: idWorker}
	}

}

func main() {
	fmt.Println("=== 🛠️ Aula 06: Padrões Avançados (Worker Pool & Pipelines) ===")

	totalTarefas := 444
	numWorkers := 5

	jobs := make(chan TarefaColeta, totalTarefas)
	results := make(chan ResultadoColeta, totalTarefas)

	var wg sync.WaitGroup

	// Inicializando o Pool com  Workers Fixos
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go workerColetor(w, jobs, results, &wg)
	}

	// Envia as tarefas para o canal de jobs
	for j := 1; j <= totalTarefas; j++ {
		jobs <- TarefaColeta{ID: j, Produto: fmt.Sprintf("Produto %d", j)}
	}
	close(jobs) // Fecha o canal de jobs após enviar todas as tarefas

	// Goroutine para fechar o canal 'results' assim que todos os workers terminarem
	go func() {
		wg.Wait()      // Aguarda todos os workers terminarem
		close(results) // Fecha o canal de resultados após todos os workers terminarem
	}()

	fmt.Println("📦 Processando respostas do Worker Pool:")
	// Coleta os resultados
	for result := range results {
		fmt.Printf("Worker %d coletou preço do %s: R$ %.2f\n", result.Worker, result.Tarefa.Produto, result.Preco)
	}

	fmt.Println("✅ Todas as tarefas foram concluídas pelo Worker Pool com limite de concorrência!")
}
