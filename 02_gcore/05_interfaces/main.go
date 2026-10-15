package main

import "fmt"

// 1 Definição da Interface (Contrato)
type NotificadorAlerta interface {
	Enviar(mensagem string) error
}

// 2 Implementação 1: Notificador via Slack
type SlackService struct {
	WebhookURL string
}

func (s SlackService) Enviar(mensagem string) error {

	fmt.Printf("[SLACK -> %s]: %s\n", s.WebhookURL, mensagem)
	return nil
}

// 3 Implementação 2: Notificador via PagerDuty (DevOps)

type PagerDutyService struct {
	IntegrationKey string
}

func (p PagerDutyService) Enviar(mensagem string) error {
	fmt.Printf("[PAGERDUTY -> Key: %s]: INCIDENTE DISPARADO: %s\n", p.IntegrationKey, mensagem)
	return nil
}

// 4 Função Polimórfica: Aceita QUALQUER tipo que satisfaça a interface NotificadorAlerta

func ProcessarFalhasSistema(n NotificadorAlerta, erroDetectado string) {
	fmt.Println("--- Processando Alerta de Infraestrutura ---")
	err := n.Enviar(erroDetectado)
	if err != nil {
		fmt.Println("Erro ao enviar alerta:", err)
	}
}

// Exemplo de Type Switch com a interface vazia 'any'
func InspecionarRecurso(recurso any) {
	switch v := recurso.(type) {
	case string:
		fmt.Printf("Recurso é uma String de texto: %q\n", v)
	case int:
		fmt.Printf("Recurso é um número Inteiro (ID/Porta): %d\n", v)
	case SlackService:
		fmt.Printf("Recurso é um serviço Slack configurado para: %s\n", v.WebhookURL)
	default:
		fmt.Printf("Tipo desconhecido: %T\n", v)
	}
}

func main() {

	fmt.Println("=== 🔌 Interfaces e Polimorfismo  ===")

	slack := SlackService{WebhookURL: "https://hooks.slack.com/services/123"}
	pager := PagerDutyService{IntegrationKey: "pd-key-999"}

	// ambas as stucts funcionam perfeitamente na mesma função sem declaração de herança
	ProcessarFalhasSistema(slack, "Uso de CPU acima dos 90% no Fedora Server")
	ProcessarFalhasSistema(pager, "Cluster K8s inalcançável")

	fmt.Println("--- Testando Type Switch com 'any' ---")
	InspecionarRecurso("Servidor Ativo")
	InspecionarRecurso(8080)
	InspecionarRecurso(slack)

}
