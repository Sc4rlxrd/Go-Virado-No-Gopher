package config

import (
	"fmt"
)

// constante exportada (pública)
const PortaPadrao = 8080

// VariavelInterna (não exportada)
var chaveSecreta = "fedora-secret-key"

// Função exportada
func ObterStringConexao(host string) string{
	return fmt.Sprintf("postgres://user:%s@%s:%d/db", chaveSecreta, host, PortaPadrao)
}