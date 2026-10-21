package main

import (
	"Go-Virado-No-Gopher/02_gcore/10_projeto_integrador/internal/domain"
	"Go-Virado-No-Gopher/02_gcore/10_projeto_integrador/internal/exporter"
	"Go-Virado-No-Gopher/02_gcore/10_projeto_integrador/internal/storage"
	"fmt"
	"os"
)

func main() {

	fmt.Println("=== 🚀 gtols-core v1.0 — Gestor de Infraestrutura (Go Core Integrado) ===")

	cache := storage.NovoCache[string, domain.RecursoInfra]()

	r1, _ := domain.NovoRecurso("srv-01", "fedora Server App", "192.168.1.10", 42.5)
	r2, _ := domain.NovoRecurso("srv-02", "Oracle DB Primary", "192.168.1.20", 91.2)
	r3, _ := domain.NovoRecurso("srv-03", "Debian Web Proxy", "192.168.1.30", 18.7)
	r4, _ := domain.NovoRecurso("srv-04", "PostgreSQL Replica", "10.0.0.15", 99.3)

	cache.Salvar(r1.ID, r1)
	cache.Salvar(r2.ID, r2)
	cache.Salvar(r3.ID, r3)
	cache.Salvar(r4.ID, r4)

	recurso := cache.ListarTodo()
	fmt.Printf("Recursos armazenados no cache genérico: %d\n", len(recurso))

	arquivoSaida := "relatorio_infra.json"
	var exp exporter.Exportador = exporter.ExportadorJSON{}
	err := exp.Exportar(arquivoSaida, recurso)
	if err != nil {
		fmt.Println("Erro ao exportar relatório:", err)
		return
	}
	fmt.Printf("Relatório exportado com sucesso para: %s\n", arquivoSaida)

	_ = os.Remove(arquivoSaida) // Limpeza pós execução
}
