package validator

import (
	"fmt"
	"net"
)

func ValidarIP(ip string) bool {

	if ip == "" {
		return false
	}

	iprs := net.ParseIP(ip)

	if iprs != nil {
		fmt.Println("Validado atráves da biblioteca net: ", iprs)
		return true
	}
	fmt.Println("IP Inválido detectado:", ip)
	return false
}
