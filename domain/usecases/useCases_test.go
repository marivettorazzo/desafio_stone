package usecases_test

import (
	"errors"
	"testing"

	"exemple.com/DesafioStone/desafio_stone/domain"
	"exemple.com/DesafioStone/desafio_stone/domain/usecases"
)

func TestValidatorTransfer(t *testing.T) {
	var (
		balance float64
		amount  float64
	)
	balance = 100
	amount = 200

	result := usecases.Validator(balance, amount)
	if result == nil {
		t.Error("resultado : esperado err obtido nil")
	}

	balance = 100
	amount = 50
	result1 := usecases.Validator(balance, amount)
	if result1 != nil {
		t.Errorf("resultado : esperado nil obtido err")
	}
	balance = 100
	amount = 0

	result2 := usecases.Validator(balance, amount)
	if result2 == nil {
		t.Errorf("resultado : esperado err obtido nil")
	}
	balance = 100
	amount = -10

	result3 := usecases.Validator(balance, amount)
	if result3 == nil {
		t.Errorf("resultado : esperado err obtido nil")
	}
}
func TestPasswordGenerator(t *testing.T) {
	str := "Olar"
	saida := "174256a25753e953bfa0571fc72f42efccecf85be47e2a4e92ab635a488b003b"
	result := usecases.PasswordGenerator(str)
	if result != saida {
		t.Error("resultado esperado:", saida, ", obtido:", result)
	}
	t.Logf("Teste senha ok")

}
func TestValidatorAccount(t *testing.T) {
	type args struct {
		nome string
		cpf  int
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{

		{
			name:    "teste1",
			args:    args{nome: "", cpf: 12345678912},
			wantErr: errors.New("invalid name"),
		},

		{
			name:    "teste2",
			args:    args{nome: "Fulano2", cpf: 12345678912},
			wantErr: nil,
		},
	}
	for i := 0; i < len(tests); i++ {
		result := usecases.ValidatorAccount(tests[i].args.nome, tests[i].args.cpf)
		if tests[i].name == "teste1" {

			if result == nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido", result)
			}

		}
		if tests[i].name == "teste2" {

			if result != nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido ", result)
			}

		}

	}

}
func TestValidaLogin(t *testing.T) {

	var Args []domain.Login
	z := domain.Login{Cpf: 123456789124, Secret: usecases.PasswordGenerator("Amora")}
	a := domain.Login{Cpf: 12358252556, Secret: usecases.PasswordGenerator("oi")}
	Args = append(Args, z, a)
	tests := []struct {
		name    string
		args    []domain.Login
		wantErr error
	}{

		{
			name:    "teste1",
			args:    Args,
			wantErr: nil,
		},
		{
			name:    "teste",
			args:    Args,
			wantErr: errors.New("invalido"),
		},
	}
	for i := 0; i < len(tests); i++ {
		_, result := usecases.ValidaLogin(Args[i])
		if tests[i].name == "teste1" {

			if result == nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido", result.Error())
			}

		}
		if tests[i].name == "teste2" {

			if result == nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido ", result.Error())
			}

		}

	}

}

func TestTransaction(t *testing.T) {
	type args struct {
		ori    int
		dest   int
		Amount float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "teste1",
			args:    args{ori: 1, dest: 2, Amount: 100},
			wantErr: errors.New(""),
		},
		{
			name:    "teste3",
			args:    args{ori: 1, dest: 2, Amount: -100},
			wantErr: errors.New(""),
		},
		{
			name:    "teste4",
			args:    args{ori: 1, dest: 2, Amount: 1000},
			wantErr: errors.New(""),
		},
		{
			name:    "teste5",
			args:    args{ori: 1, dest: 1, Amount: 100},
			wantErr: errors.New(""),
		},
	}
	for i := 0; i < len(tests); i++ {

		ori := tests[i].args.ori
		dest := tests[i].args.dest
		amount := tests[i].args.Amount

		result := usecases.Transaction(ori, dest, amount)
		if tests[i].name == "teste1" {

			if result == nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido", result.Error())
			}

		}
		if tests[i].name == "teste2" {

			if result == nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido ", result.Error())
			}

		}
		if tests[i].name == "teste3" {

			if result == nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido ", result.Error())
			}

		}
		if tests[i].name == "teste4" {

			if result == nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido ", result.Error())
			}

		}
		if tests[i].name == "teste5" {

			if result == nil {
				t.Error("resultado : esperado", tests[i].wantErr, "obtido ", result.Error())
			}

		}
	}
}
