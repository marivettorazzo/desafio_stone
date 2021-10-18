package usecases

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"exemple.com/DesafioStone/desafio_stone/domain"
)

func IDGeneratorAcc() int { //account type method that generates a non-repeating id
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10000)
	return value
}
func ValidatorAccount(nome string, cpf int) bool {
	for i := 0; i < len(DataAcc); i++ {

		if DataAcc[i].Name == "" {
			return false
		}
		if DataAcc[i].Cpf == cpf {
			if len(strconv.Itoa(cpf)) < 11 {
				return false
			}
		}
	}
	return true
}
func InsertLieDatabase(d domain.Account) (domain.Account, error) { //function that enters account data in the "Database"
	var a domain.Account
	a.Create_at = time.Now()
	accountUnique := domain.Account{
		ID:        IDGeneratorAcc(),
		Name:      d.Name,
		Cpf:       d.Cpf,
		Balance:   d.Balance,
		Secrets:   PasswordGenerator(d.Secrets),
		Create_at: time.Now(),
	}
	if !ValidatorAccount(accountUnique.Name, accountUnique.Cpf) {

		err := errors.New("404 Not Found")
		return a, err
	}

	DataAcc = append(DataAcc, accountUnique)
	insertLogin(accountUnique.Cpf, accountUnique.Secrets)
	return accountUnique, nil
}
