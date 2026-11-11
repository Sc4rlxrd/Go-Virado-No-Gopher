package main

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand/v2"
	"net/http"
	"strings"
	"time"
)

// HealthResponse DTO/Model simples para representação JSON
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Versao    string    `json:"versao"`
}

type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

// Handler para verificação de saúde do serviço (Health Check)
func healthHandler(w http.ResponseWriter, r *http.Request) {

	// 1. Validação de Método HTTP
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	// 2. Definindo o Header de Content-Type como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := HealthResponse{
		Status:    "UP",
		Timestamp: time.Now(),
		Versao:    "v1.0.0",
	}
	// 3. Serializando a struct em JSON e escrevendo na resposta
	_ = json.NewEncoder(w).Encode(res)
}

// Handler para manipulação básica de produtos
func produtosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		produtos := []Produto{
			{ID: 1, Nome: "Teclado Mecânico Fedora", Preco: 250.00},
			{ID: 2, Nome: "Mouse Vertical DevOps", Preco: 180.00},
			{ID: 3, Nome: "Monitor Samsung 25pol 75hz", Preco: 280.00},
			{ID: 4, Nome: "Mousepad ", Preco: 80.00},
			{ID: 5, Nome: "NVIDIA GeForce RTX 5070", Preco: 7200.00},
			{ID: 6, Nome: "SoundBar G1500 ", Preco: 349.00},
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(produtos)
	case http.MethodPost:
		var novoProduto Produto
		// Desserializando o JSON recebido no corpo da requisição (r.Body)
		err := json.NewDecoder(r.Body).Decode(&novoProduto)
		if err != nil {
			http.Error(w, `{"erro": "JSON inválido no corpo da requisição"}`, http.StatusBadRequest)
			return
		}

		// Validação simples do payload
		if strings.TrimSpace(novoProduto.Nome) == "" || novoProduto.Preco <= 0 {
			http.Error(w, `{"erro": "Nome e preço positivo são obrigatórios"}`, http.StatusUnprocessableEntity)
			return
		}

		novoProduto.ID = rand2.IntN(1000) + 1
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(novoProduto)
	default:
		http.Error(w, `{"erro": "Método não suportado"}`, http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	fmt.Println("=== 🌐 Starting Price Management API (net/http) na porta :8080 ===")

	// Registrando as rotas
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/v1/produtos", produtosHandler)

	// Iniciando o servidor web
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Falha crítica ao iniciar servidor HTTP: %v", err)
	}
}
