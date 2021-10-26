package usecases

import (
	"errors"

	"exemple.com/DesafioStone/desafio_stone/domain"
)

var DataLogin []domain.Login

func ValidaLogin(l domain.Login) (domain.Account, error) {
	for i := 0; i < len(DataLogin); i++ {

		if DataLogin[i].Cpf == l.Cpf && DataLogin[i].Secret == PasswordGenerator(l.Secret) {

			x, _ := Dates(DataAcc[i].Cpf)
			return x, nil
		}

	}

	return domain.Account{}, errors.New("Access denied")
}
func insertLogin(cpf int, secret string) error {

	x := domain.Login{
		Cpf:    cpf,
		Secret: secret,
	}

	DataLogin = append(DataLogin, x)
	return nil

}

func Dates(cpf int) (domain.Account, error) {
	for i := 0; i < len(DataAcc); i++ {
		if cpf == DataAcc[i].Cpf {
			return DataAcc[i], nil
		}
	}
	return domain.Account{}, errors.New(" 404 Not found")
}
