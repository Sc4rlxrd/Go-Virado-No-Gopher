package storage

import "errors"

var ErrChaveNaoEncontrada = errors.New("chave nao encontrada no cache")

// Repositório Genérico usando [K compareble, V any]
type CacheMemoria[K comparable, V any] struct {
	dados map[K]V
}

func NovoCache[K comparable, V any]() *CacheMemoria[K, V] {
	return &CacheMemoria[K, V]{
		dados: make(map[K]V),
	}
}

func (c *CacheMemoria[K, V]) Salvar(chave K, valor V) {
	c.dados[chave] = valor
}

func (c *CacheMemoria[K, V]) Obter(chave K) (V, error) {
	valor, ok := c.dados[chave]
	if !ok {
		var zero V
		return zero, ErrChaveNaoEncontrada
	}
	return valor, nil

}

func (c *CacheMemoria[K, V]) ListarTodo() []V {
	lista := make([]V, 0, len(c.dados))
	for _, v := range c.dados {
		lista = append(lista, v)
	}
	return lista
}
