package main

import (
	"errors"
	"strings"
)

var (
	ErrHostInvalido = errors.New("hostname inválido ou vazio")
	ErrHostDominio  = errors.New("hostname deve terminar com .fedora.local")
)

const sufixoFedora = ".fedora.local"

// ValidarHostFedora verifica se o hostname segue o padrão do cluster
func ValidarHostFedora(hostname string) (string, error) {
	host := strings.ToLower(strings.TrimSpace(hostname))
	if host == "" {
		return "", ErrHostInvalido
	}

	nome, ok:= strings.CutSuffix(host,sufixoFedora)
	if !ok{
		return "", ErrHostDominio
	}
	if nome == ""{
		return  "", ErrHostInvalido
	}
	return host, nil
}
