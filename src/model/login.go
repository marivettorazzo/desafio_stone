package model

type Login struct {
	Cpf    int    `json:"Cpf,omitempty"`
	Secret string `json:"Secret,omitempty"`
}
