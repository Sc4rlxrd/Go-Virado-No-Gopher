package main

import (
	"encoding/json"
	"fmt"
)

type ConfiguracaoApp struct {
	NomeApp     string `json:"nome_app"`
	MaxConexoes int    `json:"max_conexoes"`
	DebugMode   bool   `json:"debug_mode"`
}

func CarregarConfiguracao(conteudoJSON []byte) (ConfiguracaoApp, error) {
	var config ConfiguracaoApp
	err := json.Unmarshal(conteudoJSON, &config)
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return ConfiguracaoApp{}, err
	}
	return config, nil

}

func SalvarConfiguracao(config ConfiguracaoApp) ([]byte, error) {
	dadosJSON, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	return dadosJSON, nil

}

func main() {


	fmt.Println(" --- Ex 02 ---")

	configApp := ConfiguracaoApp{
		NomeApp:     "Golang for beginners",
		MaxConexoes: 3,
		DebugMode:   false,
	}

	dadosJSON, err := SalvarConfiguracao(configApp)
	if err != nil {
		fmt.Println("Erro ao salvar configuração:", err)
		return
	}

	fmt.Println("JSON gerado:", string(dadosJSON))

	configCarregada, err := CarregarConfiguracao(dadosJSON)
	if err != nil {
		fmt.Println("Erro ao carregar configuração:", err)
		return
	}

	fmt.Printf("Configuração carregada: %+v\n", configCarregada)

}
