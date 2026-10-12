package main

import "fmt"

//  Struct Base

type ServidorWeb struct {
	Host  string
	Porta int
	Ativo bool
}

// Value Receiver (Apenas Leitura / Não modifica a struct original)
func (s ServidorWeb) ObterURL() string {
	protocolo := "http"
	if s.Porta == 433 {
		protocolo = "https"
	}
	return fmt.Sprintf("%s://%s:%d", protocolo, s.Host, s.Porta)
}

// Pointer Receiver (Modifica o estado da struct original)
func (s *ServidorWeb) AlternarStatus() {
	if s == nil {
		return
	}
	s.Ativo = !s.Ativo // Modifica diretamente na memória
}

// Composição de Structs (Struct Embedding)
type ServidorFedora struct {
	ServidorWeb // Promove Host, Porta, Ativo e os métodos ObterURL/AlternarStatus
	VersãoOS    string
}

func main() {
	fmt.Println("=== 🛠️ Structs Avançadas e Métodos ===")

	// Instanciando o Servidor Fedora com a struct composta
	srv := ServidorFedora{
		ServidorWeb: ServidorWeb{
			Host:  "app.fedora.local",
			Porta: 433,
			Ativo: false,
		},
		VersãoOS: "Fedora Linux 44 Gnome",
	}

	// Acesso direto aos campos promovidos da ServidorWeb
	fmt.Printf("OS: %s | Host: %s\n", srv.VersãoOS, srv.Host)

	// Chamando método de leitura (Value Receiver)
	fmt.Println("URL de Acesso:", srv.ObterURL())

	// Chamando método de mutação (Pointer Receiver)
	fmt.Printf("Status antes de alternar: Ativo = %t\n", srv.Ativo)

	// Atalho sintático do Go: mesmo 'srv' sendo valor, o Go entende (&srv).AlternarStatus()
	srv.AlternarStatus()
	fmt.Printf("Status após alternar: Ativo = %t\n", srv.Ativo)
}
