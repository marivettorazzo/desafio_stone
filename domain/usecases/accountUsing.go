package usecases

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"exemple.com/DesafioStone/desafio_stone/domain"
)

var DataAcc []domain.Account

func Dadosfake() {
	x := domain.Account{ID: 1, Name: "Odair", Cpf: 123456789124, Balance: 500.00, Secrets: PasswordGenerator("Amora"), Create_at: time.Now()}
	y := domain.Account{ID: 2, Name: "JOsé", Cpf: 123456789321, Balance: 500.00, Secrets: PasswordGenerator("Olar"), Create_at: time.Now()}
	z := domain.Login{Cpf: 123456789124, Secret: PasswordGenerator("Amora")}
	a := domain.Login{Cpf: 123456789321, Secret: PasswordGenerator("Olar")}
	DataAcc = append(DataAcc, x, y)
	DataLogin = append(DataLogin, z, a)
}
func IDGeneratorAcc() int { //account type method that generates a non-repeating id
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10000)
	for i := 0; i < len(DataAcc); i++ {
		if DataAcc[i].ID == value {

			value = rand.Intn(10000)
		}
	}

	return value
}
func ValidatorAccount(nome string, cpf int) error {

	if len(nome) == 0 {
		return errors.New("invalid name")
	}

	for i := 0; i < len(DataAcc); i++ {

		if cpf == DataAcc[i].Cpf {

			return errors.New("invalid cpf")
		}
		if len(strconv.Itoa(cpf)) < 11 {

			return errors.New("invalid cpf")
		}
	}
	return nil

}
func InsertLieDatabase(d domain.Account) error { //function that enters account data in the "Database"

	ret := ValidatorAccount(d.Name, d.Cpf)
	if ret != nil {

		return errors.New(ret.Error())

	}
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
	return nil

}
func Accounts() error {
	if len(DataAcc) == 0 {
		return errors.New("there are no accounts to show")
	}
	return nil

}
func GetBalance(id string) (string, float64, error) {
	var y float64
	for i := 0; i < len(DataAcc); i++ {

		x, _ := strconv.Atoi(id)
		if x == DataAcc[i].ID { //makes a range traversing the dataset passed in the slice comparing if item is equal id passed in route
			y = DataAcc[i].Balance
			response := "Balance"
			return response, y, nil
		}
	}
	response := ""
	return response, y, errors.New("id not found")

}
