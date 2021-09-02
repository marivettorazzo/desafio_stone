package controllers

import (
	"exemple.com/DesafioStone/src/model"
)

var DataLogin []model.Login

func ValidaLogin(l model.Login) string {
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

	x := model.Login{
		Cpf:    cpf,
		Secret: secret,
	}

	DataLogin = append(DataLogin, x)

}
