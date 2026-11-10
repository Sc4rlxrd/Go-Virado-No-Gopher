package main

import (
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/collector"
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/domain"
	"Go-Virado-No-Gopher/03_concorrencia/09_projeto_integrador/internal/storage"
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 🚀 price-agent-core v1.0 — Motor Concorrente de Coleta ===")

	cache := storage.NovoCacheCotacoes()
	engine := collector.NovaEngineColeta(cache, 4, 500*time.Millisecond)

	solicitacoes := []domain.SolicitacaoBusca{
		{ID: 1, Produto: "Teclado Mecânico Fedora", Lojas: []string{"Amazon", "Kabum"}},
		{ID: 2, Produto: "Mouse Vertical DevOps", Lojas: []string{"MercadoLivre", "Shopee"}},
		{ID: 3, Produto: "Monitor 4K Ultrawide", Lojas: []string{"Magalu"}},
		{ID: 4, Produto: "Poco f7 512gb de rom 12ram flash", Lojas: []string{"Boa dica"}},
		{ID: 5, Produto: "NVIDIA GeForce RTX 5070", Lojas: []string{"TeraByte"}},
	}

	ctxPai := context.Background()
	inicio := time.Now()
	cotacoes, erros := engine.ExecutarLote(ctxPai, solicitacoes)

	fmt.Printf("⏱️ Tempo total de execução concorrente: %v\n", time.Since(inicio))
	fmt.Printf("📦 Total de cotações obtidas: %d\n", len(cotacoes))
	fmt.Printf("⚠️ Total de erros capturados: %d\n", len(erros))

	fmt.Println("📋 Cotações em Cache:")
	for _, req := range cache.ListarTodos() {
		fmt.Printf(" • [%-12s] %-25s | R$ %.2f\n", req.Loja, req.Produto, req.Preco)
	}
}
