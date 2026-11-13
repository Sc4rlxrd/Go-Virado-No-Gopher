package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

func chamar(metodo, alvo string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(metodo, alvo, nil)
	rec := httptest.NewRecorder()
	produtosBuscaHandler(rec, req)
	return rec
}

func decodificar[T any](t *testing.T, rec *httptest.ResponseRecorder) T {
	t.Helper()
	var v T
	if err := json.Unmarshal(rec.Body.Bytes(), &v); err != nil {
		t.Fatalf("resposta não é JSON válido: %v (corpo: %q)", err, rec.Body.String())
	}
	return v
}

func idsDe(ps []Produto) []int {
	ids := make([]int, 0, len(ps))
	for _, p := range ps {
		ids = append(ids, p.ID)
	}
	return ids
}

func TestProdutosBuscaPorID(t *testing.T) {
	testes := []struct {
		nome     string
		alvo     string
		status   int
		esperado *Produto // produto esperado no 200
		erro     string   // mensagem esperada em {"erro": ...}
	}{
		{"id existente", "/api/v1/produtos?id=1", 200, &produtos[0], ""},
		{"outro id existente", "/api/v1/produtos?id=3", 200, &produtos[2], ""},
		{"id inexistente", "/api/v1/produtos?id=999", 404, nil, "produto nao encontrado"},
		{"id zero", "/api/v1/produtos?id=0", 404, nil, "produto nao encontrado"},
		{"id negativo", "/api/v1/produtos?id=-1", 404, nil, "produto nao encontrado"},
		{"id alfanumérico", "/api/v1/produtos?id=abc", 400, nil, "id invalido"},
		{"id misto", "/api/v1/produtos?id=1abc", 400, nil, "id invalido"},
		{"id decimal", "/api/v1/produtos?id=1.5", 400, nil, "id invalido"},
		{"id vazio", "/api/v1/produtos?id=", 400, nil, "id invalido"},
		{"id vence a categoria", "/api/v1/produtos?id=3&categoria=hardware", 200, &produtos[2], ""},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			rec := chamar(http.MethodGet, tt.alvo)

			assertIgual(t, rec.Code, tt.status)
			assertIgual(t, rec.Header().Get("Content-Type"), "application/json")

			if tt.esperado != nil {
				assertIgual(t, decodificar[Produto](t, rec), *tt.esperado)
				return
			}
			assertIgual(t, decodificar[respostaErro](t, rec).Erro, tt.erro)
		})
	}
}

func TestProdutosBuscaPorCategoria(t *testing.T) {
	testes := []struct {
		nome string
		alvo string
		ids  []int
	}{
		{"hardware", "/api/v1/produtos?categoria=hardware", []int{1, 2}},
		{"perifericos", "/api/v1/produtos?categoria=perifericos", []int{3, 4}},
		{"maiúsculas", "/api/v1/produtos?categoria=HARDWARE", []int{1, 2}},
		{"misto", "/api/v1/produtos?categoria=PeRiFeRiCoS", []int{3, 4}},
		{"com espaços nas pontas", "/api/v1/produtos?categoria=%20hardware%20", []int{1, 2}},
		{"categoria vazia: sem filtro", "/api/v1/produtos?categoria=", []int{1, 2, 3, 4, 5}},
		{"sem parâmetros: lista completa", "/api/v1/produtos", []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			rec := chamar(http.MethodGet, tt.alvo)

			assertIgual(t, rec.Code, http.StatusOK)
			assertIgual(t, rec.Header().Get("Content-Type"), "application/json")

			obtidos := idsDe(decodificar[[]Produto](t, rec))
			if !slices.Equal(obtidos, tt.ids) {
				t.Errorf("ids = %v; esperado %v", obtidos, tt.ids)
			}
		})
	}
}

// Categoria sem resultado: o corpo tem que ser "[]" e não "null".
func TestProdutosBuscaCategoriaSemResultadoRetornaArrayVazio(t *testing.T) {
	rec := chamar(http.MethodGet, "/api/v1/produtos?categoria=xyz")

	assertIgual(t, rec.Code, http.StatusOK)
	assertIgual(t, strings.TrimSpace(rec.Body.String()), "[]")
}

func TestProdutosBuscaMetodoInvalido(t *testing.T) {
	for _, metodo := range []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch} {
		t.Run(metodo, func(t *testing.T) {
			rec := chamar(metodo, "/api/v1/produtos")

			assertIgual(t, rec.Code, http.StatusMethodNotAllowed)
			assertIgual(t, rec.Header().Get("Allow"), "GET")
			assertIgual(t, rec.Header().Get("Content-Type"), "application/json")
		})
	}
}
