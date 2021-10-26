package usecases_test

// import (
// 	"testing"

// 	"exemple.com/DesafioStone/desafio_stone/domain/usecases"
// )

// type TestAccount struct {
// 	nome string
// 	cpf  int
// }

// func TestValidatorAccount(t *testing.T) {

// 	ClienteTeste1 := TestAccount{nome: "José", cpf: 58258263998}
// 	result := usecases.ValidatorAccount(ClienteTeste1.nome, ClienteTeste1.cpf)
// 	if result != nil {
// 		t.Error("resultado : esperado nil obtido err")
// 	}

// 	ClienteTeste2 := TestAccount{nome: "", cpf: 12345678996}

// 	result2 := usecases.ValidatorAccount(ClienteTeste2.nome, ClienteTeste2.cpf)
// 	if result2 == nil {
// 		t.Error("resultado : esperado err obtido nil")
// 	}

// }
// func TestValidatorTransfer(t *testing.T) {
// 	var (
// 		balance float64
// 		amount  float64
// 	)
// 	balance = 100
// 	amount = 200

// 	result := usecases.Validator(balance, amount)
// 	if result == nil {
// 		t.Error("resultado : esperado err obtido nil")
// 	}

// 	balance = 100
// 	amount = 50
// 	result1 := usecases.Validator(balance, amount)
// 	if result1 != nil {
// 		t.Errorf("resultado : esperado nil obtido err")
// 	}
// 	balance = 100
// 	amount = 0

// 	result2 := usecases.Validator(balance, amount)
// 	if result2 == nil {
// 		t.Errorf("resultado : esperado err obtido nil")
// 	}
// 	balance = 100
// 	amount = -10

// 	result3 := usecases.Validator(balance, amount)
// 	if result3 == nil {
// 		t.Errorf("resultado : esperado err obtido nil")
// 	}
// }
// func TestPasswordGenerator(t *testing.T) {
// 	str := "Olar"
// 	saida := "174256a25753e953bfa0571fc72f42efccecf85be47e2a4e92ab635a488b003b"
// 	result := usecases.PasswordGenerator(str)
// 	if result != saida {
// 		t.Error("resultado esperado:", saida, ", obtido:", result)
// 	}
// 	t.Logf("Teste senha ok")

// }
// func TestValidatorAccount(t *testing.T) {
// 	type args struct {
// 		nome string
// 		cpf  int
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr error
// 	}{

// 		{
// 			name:    "teste1",
// 			args:    args{nome: "", cpf: 12345678912},
// 			wantErr: errors.New("invalid name"),
// 		},

// 		{
// 			name:    "teste2",
// 			args:    args{nome: "Fulano2", cpf: 12345678912},
// 			wantErr: nil,
// 		},
// 	}
// 	for i := 0; i < len(tests); i++ {
// 		result := ValidatorAccount(tests[i].args.nome, tests[i].args.cpf)
// 		if tests[i].name == "teste1" {

// 			if result == nil {
// 				t.Error("resultado : esperado", tests[i].wantErr, "obtido", result)
// 			}

// 		}
// 		if tests[i].name == "teste2" {

// 			if result != nil {
// 				t.Error("resultado : esperado", tests[i].wantErr, "obtido ", result)
// 			}

// 		}

// 	}

// }
// func TestValidaLogin(t *testing.T) {

// 	type args struct {
// 		cpf    int
// 		secret string
// 	}
// 	secret1:= usecases.AsSha256("Amora")
// 	secret2:= usecases.AsSha256("oi")
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr error
// 	}	{
// 			{
// 			name:    "teste1",
// 			args:    args{
// 				Cpf: 12345678333,
// 				Secret: "Amora"
// 			},
// 			wantErr: nil,
// 			},
// 			{
// 			name:    "teste2",
// 			args:    args{
// 				cpf: 10000000000,
// 				 secret: "oi"
// 				},
// 			wantErr: errors.New("Access denied"),
// 			},
// 		}

// 	for i := 0; i < len(tests); i++ {
// 		_, x := ValidaLogin(tests[i].args.cpf, tests[i].args.secret)
// 		if tests[i].name == "teste1" {
// 			if x != nil {
// 				t.Errorf("ValidaLogin() = %v, want %v", x, tests.wantErr)
// 			}
// 		}
// 		if tests[i].name == "teste2" {
// 			if x == nil {
// 				t.Errorf("ValidaLogin() = %v, want %v", x, tests.wantErr)
// 			}
// 		}
// 	}

// }
// func TestTransaction(t *testing.T) {
// 	type args struct {
// 		ori    int
// 		dest   int
// 		Amount float64
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := Transaction(tt.args.ori, tt.args.dest, tt.args.Amount); (err != nil) != tt.wantErr {
// 				t.Errorf("Transaction() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
