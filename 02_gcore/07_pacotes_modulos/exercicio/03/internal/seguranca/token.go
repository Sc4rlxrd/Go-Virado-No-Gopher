package seguranca

import (
	"crypto/rand"
)

type TokenJWT struct {
	GerarHash string
}

func GerarHash() string {

	hash := rand.Text() + rand.Text()
	return hash
}
