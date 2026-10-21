package main

import (
	"Go-Virado-No-Gopher/02_gcore/10_projeto_integrador/internal/domain"
	"Go-Virado-No-Gopher/02_gcore/10_projeto_integrador/internal/storage"
	"testing"
)

func TestFluxoIntegracaoCore(t *testing.T) {
	testes := []struct {
		nome        string
		id          string
		nomeRecurso string
		ip          string
		cpu         float64
		statusEsp   domain.StatusRecurso
		deveErro    bool
	}{
		{nome: "Recurso Normal (ONLINE)", id: "r1", nomeRecurso: "WebServer Fedora", ip: "10.0.0.1", cpu: 50.0, statusEsp: domain.StatusOnline, deveErro: false},
		{nome: "Recurso com CPU Alta (ALERTA)", id: "r2", nomeRecurso: "DB Node", ip: "10.0.0.2", cpu: 95.0, statusEsp: domain.StatusAlerta, deveErro: false},
		{nome: "Nome Vazio (Deve Falhar)", id: "r3", nomeRecurso: " ", ip: "10.0.0.3", cpu: 10.0, deveErro: true},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			rec, err := domain.NovoRecurso(tt.id, tt.nomeRecurso, tt.ip, tt.cpu)
			if (err != nil) != tt.deveErro {
				t.Fatalf("status de erro inesperado: %v", err)
			}

			if !tt.deveErro {
				if rec.Status != tt.statusEsp {
					t.Errorf("esperado status %s, obtido %s", tt.statusEsp, rec.Status)
				}
			}
			cache := storage.NovoCache[string, domain.RecursoInfra]()
			cache.Salvar(rec.ID, rec)
			recSalvo, errCache := cache.Obter(rec.ID)

			if errCache != nil || recSalvo.ID != rec.ID {
				t.Errorf("falha ao salvar/obter do cache genérico")
			}

		})
	}
}
