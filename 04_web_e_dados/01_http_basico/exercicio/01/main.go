package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Loja struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	CNPJ  string `json:"cnpj"`
	Ativa bool   `json:"ativa"`
}

type respostaErro struct {
	Erro string `json:"erro"`
}

const tamanhoMaximoCorpo = 1 << 20 // 1 MiB

// escreverJSON centraliza o padrão: header, status e corpo, nessa ordem.
func escreverJSON(w http.ResponseWriter, status int, corpo any) {
	w.Header().Set("Content-Type", "application/json") // tem que vir ANTES do WriteHeader
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(corpo); err != nil {
		log.Printf("erro ao escrever resposta: %v", err)
	}
}

func lojasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listarLojas(w)
	case http.MethodPost:
		criarLoja(w, r)
	default:
		w.Header().Set("Allow", "GET, POST")
		escreverJSON(w, http.StatusMethodNotAllowed, respostaErro{Erro: "metodo nao permitido"})
	}
}

func listarLojas(w http.ResponseWriter) {
	lojas := []Loja{
		{ID: 1, Nome: "Amazon", CNPJ: "15436940000103", Ativa: true},
		{ID: 2, Nome: "Kabum", CNPJ: "05570714000159", Ativa: true},
	}
	escreverJSON(w, http.StatusOK, lojas)
}

func criarLoja(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, tamanhoMaximoCorpo)

	var nova Loja
	if err := json.NewDecoder(r.Body).Decode(&nova); err != nil {
		escreverJSON(w, http.StatusBadRequest, respostaErro{Erro: "json invalido"})
		return
	}

	if strings.TrimSpace(nova.Nome) == "" || len(nova.CNPJ) < 14 {
		escreverJSON(w, http.StatusUnprocessableEntity, respostaErro{Erro: "dados da loja invalidos"})
		return
	}

	nova.ID = 10      // o servidor decide, ignora o que o cliente mandou
	nova.Ativa = true // idem
	escreverJSON(w, http.StatusCreated, nova)
}

func main() {

	fmt.Println(" --- Ex 01 --- ")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/lojas", lojasHandler)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second, // evita conexão lenta segurando o servidor
	}

	log.Println("servidor escutando em :8080")
	log.Fatal(srv.ListenAndServe())
}
