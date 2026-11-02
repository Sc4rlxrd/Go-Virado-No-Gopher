package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type RequisicaoURL struct {
	ID  int
	URL string
}

type RespostaURL struct {
	ID       int
	URL      string
	Valida   bool
	WorkerId int
}

func WorkerValidador(id int, jobs <-chan RequisicaoURL, results chan<- RespostaURL, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {

		results <- RespostaURL{
			ID:       job.ID,
			URL:      job.URL,
			Valida:   strings.Contains(job.URL, "https://"),
			WorkerId: id,
		}

	}
}

func ExecutarWorkerPool(urls []string, numWorkers int) []RespostaURL {
	totalTarefas := len(urls)
	jobs := make(chan RequisicaoURL, totalTarefas)
	results := make(chan RespostaURL, totalTarefas)

	if numWorkers == 0 {
		return []RespostaURL{}
	}

	var wg sync.WaitGroup

	// Inicializando o Pool com Workers Fixos
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go WorkerValidador(w, jobs, results, &wg)
	}

	// Envia as tarefas para o canal de jobs
	for j, url := range urls {
		jobs <- RequisicaoURL{ID: j + 1, URL: url}
	}
	close(jobs) // Fecha o canal de jobs após enviar todas as tarefas

	// Goroutine para fechar o canal 'results' assim que todos os workers terminarem
	go func() {
		wg.Wait()      // Aguarda todos os workers terminarem
		close(results) // Fecha o canal de resultados
	}()

	// Coleta os resultados
	respostas := make([]RespostaURL, 0, len(urls))
	for res := range results {
		respostas = append(respostas, res)
	}

	return respostas
}

func main() {
	fmt.Println("=== Exercício 1: Worker Pool ===")

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"http://www.youtube.com/",
	}

	fmt.Println("Antes do Executar os Workers")
	for _, url := range urls {
		fmt.Println(url)
	}

	fmt.Println("\nExecutando os Workers...")
	time.Sleep(1 * time.Second)

	for _, r := range ExecutarWorkerPool(urls, 3) {
		fmt.Printf("ID: %d, URL: %s, Válida: %t, Worker: %d\n", r.ID, r.URL, r.Valida, r.WorkerId)
	}

}
