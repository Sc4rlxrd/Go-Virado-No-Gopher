package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Categoria string  `json:"categoria"`
	Preco     float64 `json:"preco"`
}

type respostaErro struct {
	Erro string `json:"erro"`
}

// Mock em memória. Só é lido pelo handler, então não precisa de mutex.
var produtos = []Produto{
	{ID: 1, Nome: "Placa de Video RTX 5070", Categoria: "hardware", Preco: 4999.90},
	{ID: 2, Nome: "SSD NVMe 1TB", Categoria: "hardware", Preco: 429.90},
	{ID: 3, Nome: "Mouse Gamer Logitech", Categoria: "perifericos", Preco: 199.90},
	{ID: 4, Nome: "Teclado Mecanico Redragon", Categoria: "perifericos", Preco: 259.90},
	{ID: 5, Nome: "Monitor 4K Ultrawide", Categoria: "monitores", Preco: 2899.00},
}

func escreverJSON(w http.ResponseWriter, status int, corpo any) {
	w.Header().Set("Content-Type", "application/json") // antes do WriteHeader
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(corpo); err != nil {
		log.Printf("erro ao escrever resposta: %v", err)
	}
}

func produtosBuscaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		escreverJSON(w, http.StatusMethodNotAllowed, respostaErro{Erro: "metodo nao permitido"})
		return
	}

	query := r.URL.Query()

	// Busca por ID: tem prioridade sobre os outros filtros.
	// Has() distingue "?id=" (informado, mas vazio) de "sem id".
	if query.Has("id") {
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			escreverJSON(w, http.StatusBadRequest, respostaErro{Erro: "id invalido"})
			return
		}
		for _, p := range produtos {
			if p.ID == id {
				escreverJSON(w, http.StatusOK, p)
				return
			}
		}
		escreverJSON(w, http.StatusNotFound, respostaErro{Erro: "produto nao encontrado"})
		return
	}

	// Filtro por categoria (insensível a maiúsculas/minúsculas)
	categoria := strings.ToLower(strings.TrimSpace(query.Get("categoria")))
	if categoria == "" {
		escreverJSON(w, http.StatusOK, produtos) // sem parâmetros: lista completa
		return
	}

	filtrados := make([]Produto, 0) // NÃO nil: senão o JSON sai "null" em vez de "[]"
	for _, p := range produtos {
		if strings.ToLower(p.Categoria) == categoria {
			filtrados = append(filtrados, p)
		}
	}
	escreverJSON(w, http.StatusOK, filtrados)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/produtos", produtosBuscaHandler)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("servidor escutando em :8080")
	log.Fatal(srv.ListenAndServe())
}
