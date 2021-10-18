package usecases

import (
	"exemple.com/DesafioStone/desafio_stone/domain"
)

var DataLogin []domain.Login

func ValidaLogin(l domain.Login) string {
	for i := 0; i < len(DataLogin); i++ {

		if DataLogin[i].Cpf == l.Cpf && DataLogin[i].Secret == PasswordGenerator(l.Secret) {
			ret := "Access allowed"
			return ret
		}

	}
	ret := "Access denied"

	return ret
}
func insertLogin(cpf int, secret string) {

	x := domain.Login{
		Cpf:    cpf,
		Secret: secret,
	}

	DataLogin = append(DataLogin, x)

}
