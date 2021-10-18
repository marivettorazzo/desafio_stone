package main

import (
	"testing"
	"time"

	"exemple.com/DesafioStone/desafio_stone/domain"
	"exemple.com/DesafioStone/desafio_stone/domain/usecases"
)

var (
	a []domain.Account
)

func atribui() {
	ClienteTeste1 := domain.Account{
		ID:        10,
		Name:      "clientet1",
		Cpf:       88855522200,
		Balance:   500,
		Secrets:   usecases.AsSha256("Olar"),
		Create_at: time.Now(),
	}

	clienteTeste2 := domain.Account{
		ID:        11,
		Name:      "clientet2",
		Cpf:       88855522333,
		Balance:   500,
		Secrets:   usecases.AsSha256("Olar"),
		Create_at: time.Now(),
	}
	a = append(a, ClienteTeste1)
	a = append(a, clienteTeste2)
}

func TestTransfer(t *testing.T) {

	positive := true
	result := usecases.Transaction(10, 11, 200)
	if result != positive {
		t.Errorf("resultado esperado: %t, obtida: %t", !positive, result)
	}

}

func TestCod(t *testing.T) {
	str := "Olar"
	saida := "174256a25753e953bfa0571fc72f42efccecf85be47e2a4e92ab635a488b003b"
	result := usecases.AsSha256(str)
	if result != saida {
		t.Errorf("resultado esperado: %s, obtido: %s", saida, result)
	}

}

func TestPasswordGenerator(t *testing.T) {
	str := "Olar"
	saida := "174256a25753e953bfa0571fc72f42efccecf85be47e2a4e92ab635a488b003b"
	result := usecases.PasswordGenerator(str)
	if result != saida {
		t.Error("resultado esperado:", saida, ", obtido:", result)
	}

}
