package usecases

import (
	"fmt"
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

	if nome == "" {

		return false
	}
	for i := 0; i < len(DataAcc); i++ {

		if cpf == DataAcc[i].Cpf {

			if len(strconv.Itoa(cpf)) < 11 {

				return false
			}
			return false
		}
	}
	return true
}
func InsertLieDatabase(d domain.Account) bool { //function that enters account data in the "Database"

	if ValidatorAccount(d.Name, d.Cpf) {

		fmt.Print("Validou como true")
		accountUnique := domain.Account{
			ID:        IDGeneratorAcc(),
			Name:      d.Name,
			Cpf:       d.Cpf,
			Balance:   d.Balance,
			Secrets:   PasswordGenerator(d.Secrets),
			Create_at: time.Now(),
		}

		DataAcc = append(DataAcc, accountUnique)
		insertLogin(accountUnique.Cpf, accountUnique.Secrets)
		return true
	}
	return false
}
