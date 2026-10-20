package domain

import (
	"errors"
	"strings"
	"time"
)

var ErrNomeInvalido = errors.New("nome do recurso nao pode ser vazio")

type StatusRecurso string

const (
	StatusOnline  StatusRecurso = "ONLINE"
	StatusOffline StatusRecurso = "OFFILINE"
	StatusAlerta  StatusRecurso = "ALERTA"
)

type RecursoInfra struct {
	ID       string        `json:"id"`
	Nome     string        `json:"nome"`
	IP       string        `json:"ip"`
	Status   StatusRecurso `json:"status"`
	UsoCPU   float64       `json:"uso_cpu"`
	UpdateAt time.Time     `json:"updated_at"`
}

func NovoRecurso(id, nome, ip string, cpu float64) (RecursoInfra, error) {
	if strings.TrimSpace(nome) == "" {
		return RecursoInfra{}, ErrNomeInvalido
	}
	status := StatusOnline
	if cpu > 85.0 {
		status = StatusAlerta
	}
	return RecursoInfra{
		ID:       id,
		Nome:     nome,
		IP:       ip,
		Status:   status,
		UsoCPU:   cpu,
		UpdateAt: time.Now(),
	}, nil

}
