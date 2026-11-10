package storage

import (
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/domain"
	"sync"
)

type CacheCotacoes struct {
	mu    sync.RWMutex
	itens map[string]domain.ItemCotacao
}

func NovoCacheCotacoes() *CacheCotacoes {
	return &CacheCotacoes{
		itens: make(map[string]domain.ItemCotacao),
	}
}

func (c *CacheCotacoes) Salvar(item domain.ItemCotacao) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.itens[item.Produto+":"+item.Loja] = item

}

func (c *CacheCotacoes) Obter(produto, loja string) (domain.ItemCotacao, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.itens[produto+":"+loja]
	return item, ok
}

func (c *CacheCotacoes) ListarTodos() []domain.ItemCotacao {
	c.mu.Lock()
	defer c.mu.Unlock()

	lista := make([]domain.ItemCotacao, 0, len(c.itens))
	for _, item := range c.itens {
		lista = append(lista, item)
	}
	return lista
}
