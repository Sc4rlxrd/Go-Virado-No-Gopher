package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Struct com Struct Tags para mapeamento JSON
type Metricaprod struct {
	Servico   string    `json:"nome_servico"`
	CPUUsage  float64   `json:"uso_cpu"`
	Timestamp time.Time `json:"data_registro"`
}

func main() {

	fmt.Println("===  Biblioteca Padrão   ===")

	// 1. Variáveis de Ambiente (os)
	envAmbiente := os.Getenv("ENV")
	if envAmbiente == "" {
		envAmbiente = "development" // valor padrão
	}

	fmt.Printf("Ambiente detectado: %s\n", envAmbiente)

	// 2. Medindo Tempo de Execução (time)
	inicio := time.Now()
	time.Sleep(50 * time.Millisecond) // simula um processamento leve
	duracao := time.Since(inicio)
	fmt.Printf("Tempo de processamento simulado: %v\n", duracao)

	// 3. Criando uma Struct e convertendo para JSON (json.Marshal)
	metrica := Metricaprod{
		Servico:   "fedora-auth-service",
		CPUUsage:  80.4,
		Timestamp: time.Now(),
	}

	dadosJson, err := json.MarshalIndent(metrica, "", " ")
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return
	}
	fmt.Println("JSON Gerado:")
	fmt.Println(string(dadosJson))

	// 4. Salvação de Arquivo em Disco (os.WriteFile)
	caminhoArquivo := "metrica_temp.json"
	err = os.WriteFile(caminhoArquivo, dadosJson, 0644)
	if err != nil {
		fmt.Println("Erro ao salver o arquivo: ", err)
		return
	}

	fmt.Printf("Arquivo: %s  salvo com sucesso em disco\n", caminhoArquivo)

	// 5. Lendo Arquivo do Disco e convertendo para Struct (os.ReadFile + json.Unmarshal)

	conteudoLido, err := os.ReadFile(caminhoArquivo)
	if err != nil {
		fmt.Println("Erro ao desserlizar JSON: ", err)
		return
	}
	var metricaRestaurada Metricaprod
	err = json.Unmarshal(conteudoLido, &metricaRestaurada)
	if err != nil {
		fmt.Println("Erro ao desserializar JSON:", err)
		return
	}

	fmt.Printf("Métrica Restaurada do Arquivo -> Serviço: %s | CPU: %.1f%%\n", metricaRestaurada.Servico, metricaRestaurada.CPUUsage)

	// Limpeza do arquivo temporário criado
	_ = os.Remove(caminhoArquivo)

}
