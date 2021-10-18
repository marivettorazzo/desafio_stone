package domain

import (
	"time"
)

type Account struct {
	ID        int       `json:"ID,omitempty" `
	Name      string    `json:"Name,omitempty"`
	Cpf       int       `json:"Cpf,omitempty"` // A rota para essa entidade deve retornar tudo
	Balance   float64   `json:"Balance,omitempty"`
	Secrets   string    `json:"Secret,omitempty"`
	Create_at time.Time `json:"create_at,omitempty"`
}
