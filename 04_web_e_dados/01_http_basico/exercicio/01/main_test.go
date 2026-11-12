package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func assertIgual[T comparable](t *testing.T, obtido, esperado T) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("falha na asserção: obtido %v, esperado %v", obtido, esperado)
	}
}

func chamar(metodo, corpo string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(metodo, "/api/v1/lojas", strings.NewReader(corpo))
	rec := httptest.NewRecorder()
	lojasHandler(rec, req)
	return rec
}

func TestLojasHandlerGET(t *testing.T) {
	rec := chamar(http.MethodGet, "")

	assertIgual(t, rec.Code, http.StatusOK)
	assertIgual(t, rec.Header().Get("Content-Type"), "application/json")

	var lojas []Loja
	if err := json.NewDecoder(rec.Body).Decode(&lojas); err != nil {
		t.Fatalf("resposta não é JSON válido: %v", err)
	}
	if len(lojas) < 2 {
		t.Errorf("len = %d; esperado pelo menos 2 lojas", len(lojas))
	}
}

func TestLojasHandlerPOST(t *testing.T) {
	testes := []struct {
		nome     string
		corpo    string
		status   int
		erro     string // mensagem esperada em {"erro": ...}; vazio = sem erro
		esperada *Loja  // loja esperada no 201
	}{
		{"válido", `{"nome":"Kabum Tech","cnpj":"12345678000199"}`, 201, "",
			&Loja{ID: 10, Nome: "Kabum Tech", CNPJ: "12345678000199", Ativa: true}},
		{"cnpj com exatamente 14: aceita", `{"nome":"X","cnpj":"12345678901234"}`, 201, "",
			&Loja{ID: 10, Nome: "X", CNPJ: "12345678901234", Ativa: true}},
		{"cliente não escolhe id nem ativa", `{"id":99,"nome":"X","cnpj":"12345678901234","ativa":false}`, 201, "",
			&Loja{ID: 10, Nome: "X", CNPJ: "12345678901234", Ativa: true}},
		{"nome vazio", `{"nome":"","cnpj":"12345678000199"}`, 422, "dados da loja invalidos", nil},
		{"nome só com espaços", `{"nome":"   ","cnpj":"12345678000199"}`, 422, "dados da loja invalidos", nil},
		{"cnpj com 13: rejeita", `{"nome":"X","cnpj":"1234567890123"}`, 422, "dados da loja invalidos", nil},
		{"cnpj ausente", `{"nome":"X"}`, 422, "dados da loja invalidos", nil},
		{"nome e cnpj inválidos", `{"nome":"","cnpj":"123"}`, 422, "dados da loja invalidos", nil},
		{"json quebrado", `{nome: kabum`, 400, "json invalido", nil},
		{"corpo vazio", ``, 400, "json invalido", nil},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			rec := chamar(http.MethodPost, tt.corpo)

			assertIgual(t, rec.Code, tt.status)
			assertIgual(t, rec.Header().Get("Content-Type"), "application/json")

			if tt.esperada != nil {
				var obtida Loja
				if err := json.NewDecoder(rec.Body).Decode(&obtida); err != nil {
					t.Fatalf("resposta não é JSON válido: %v", err)
				}
				assertIgual(t, obtida, *tt.esperada) // Loja é comparável: todos os campos são simples
				return
			}

			var e respostaErro
			if err := json.NewDecoder(rec.Body).Decode(&e); err != nil {
				t.Fatalf("resposta não é JSON válido: %v", err)
			}
			assertIgual(t, e.Erro, tt.erro)
		})
	}
}

func TestLojasHandlerMetodoNaoPermitido(t *testing.T) {
	for _, metodo := range []string{http.MethodPut, http.MethodDelete, http.MethodPatch} {
		t.Run(metodo, func(t *testing.T) {
			rec := chamar(metodo, "")

			assertIgual(t, rec.Code, http.StatusMethodNotAllowed)
			assertIgual(t, rec.Header().Get("Allow"), "GET, POST")
			assertIgual(t, rec.Header().Get("Content-Type"), "application/json")
		})
	}
}
