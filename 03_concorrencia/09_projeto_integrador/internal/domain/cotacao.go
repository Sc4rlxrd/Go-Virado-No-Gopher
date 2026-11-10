package domain

import "time"

type ItemCotacao struct {
	ID       int       `json:"id"`
	Produto  string    `json:"produto"`
	Loja     string    `json:"loja"`
	Preco    float64   `json:"preco"`
	CriadoEm time.Time `json:"criado_em"`
}

type SolicitacaoBusca struct {
	ID      int
	Produto string
	Lojas   []string
}

type ResultadoColeta struct {
	Cotacao ItemCotacao
	Erro    error
}
