package usecases

import (
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
func ValidatorAccount(nome string, cpf int) (string, bool) {
	var response string
	if nome == "" {
		response = "Preencha um nome válido"
		return response, false
	}
	for i := 0; i < len(DataAcc); i++ {

		if cpf == DataAcc[i].Cpf {
			response = "Preencha um CPF válido"
			if len(strconv.Itoa(cpf)) < 11 {
				response = "Preencha um CPF válido"
				return response, false
			}
			return response, false
		}
	}
	response = "Sucesso"
	return response, true
}
func InsertLieDatabase(d domain.Account) (string, bool) { //function that enters account data in the "Database"
	var response string
	str, ret := ValidatorAccount(d.Name, d.Cpf)
	if ret {

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
		response = "Sucesso"
		return response, true
	}
	return str, ret
}
