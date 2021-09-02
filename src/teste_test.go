package main

import (
	"testing"
	"time"

	"exemple.com/DesafioStone/src/controllers"
	"exemple.com/DesafioStone/src/model"
)

var (
	a []model.Account
)

func atribui() {
	ClienteTeste1 := model.Account{
		ID:        10,
		Name:      "clientet1",
		Cpf:       88855522200,
		Balance:   500,
		Secrets:   controllers.AsSha256("Olar"),
		Create_at: time.Now(),
	}

	clienteTeste2 := model.Account{
		ID:        11,
		Name:      "clientet2",
		Cpf:       88855522333,
		Balance:   500,
		Secrets:   controllers.AsSha256("Olar"),
		Create_at: time.Now(),
	}
	a = append(a, ClienteTeste1)
	a = append(a, clienteTeste2)
}

func TestTransfer(t *testing.T) {

	positive := true
	result := controllers.TransactionOccurs(10, 11, 200)
	if result != positive {
		t.Errorf("resultado esperado: %t, obtida: %t", !positive, result)
	}

}

func TestCod(t *testing.T) {
	str := "Olar"
	saida := "174256a25753e953bfa0571fc72f42efccecf85be47e2a4e92ab635a488b003b"
	result := controllers.AsSha256(str)
	if result != saida {
		t.Errorf("resultado esperado: %s, obtido: %s", saida, result)
	}

}

func TestPasswordGenerator(t *testing.T) {
	str := "Olar"
	saida := "174256a25753e953bfa0571fc72f42efccecf85be47e2a4e92ab635a488b003b"
	result := controllers.PasswordGenerator(str)
	if result != saida {
		t.Error("resultado esperado:", saida, ", obtido:", result)
	}

}
