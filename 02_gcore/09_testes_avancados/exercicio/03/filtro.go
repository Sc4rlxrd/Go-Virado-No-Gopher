package main

import (
	"errors"
	"strings"
)

var ErrListaVazia = errors.New("lista de ips nil ou vazia")

type FiltroIp struct{}

func (f FiltroIp) FiltrarPrivados(ips []string) ([]string, error) {
	if len(ips) == 0 {
		return nil, ErrListaVazia
	}

	privados := []string{}

	for _, ip := range ips {
		if strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "192.168.") {
			privados = append(privados, ip)
		}
	}
	return privados, nil
}
