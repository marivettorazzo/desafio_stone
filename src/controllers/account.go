package controllers

import (
	"errors"
	"math/rand"
	"time"

	"exemple.com/DesafioStone/src/model"
)

func IDGeneratorAcc() int { //account type method that generates a non-repeating id
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10000)
	return value
}
func InsertLieDatabase(d model.Account) (model.Account, error) { //function that enters account data in the "Database"
	var a model.Account
	a.Create_at = time.Now()
	accountUnique := model.Account{
		ID:        IDGeneratorAcc(),
		Name:      d.Name,
		Cpf:       d.Cpf,
		Balance:   d.Balance,
		Secrets:   PasswordGenerator(d.Secrets),
		Create_at: time.Now(),
	}
	if accountUnique.Cpf == 0 || accountUnique.Name == "" {
		err := errors.New("404 Not Found")
		return a, err
	}

	DataAcc = append(DataAcc, accountUnique)
	insertLogin(accountUnique.Cpf, accountUnique.Secrets)
	return accountUnique, nil
}
