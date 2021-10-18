package usecases

import (
	"math/rand"
	"time"

	"exemple.com/DesafioStone/desafio_stone/domain"
)

var (
	DataAcc  []domain.Account
	DataTran []domain.Transfer
)

func IDGeneratorTransaction() int { // Transfer entity function that generates transaction ID
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(50)
	return value
}
func CaptureId(id int) (string, bool) {
	var ret string
	for i := 0; i < len(DataAcc); i++ {
		if DataAcc[i].ID == id {
			ret = "Sucess"
			return ret, true
		}
	}
	ret = "invalid id"
	return ret, false
}

func Validator(balance float64, Amount float64) (string, bool) {
	var response string
	switch {
	case balance < Amount:
		response = "Not enough balance for transaction"
		return response, false
	case Amount < 0:
		response = "Invalid transfer value"
		return response, false
	case Amount == 0:
		response = "Invalid transfer value"
		return response, false
	default:
		response = "Sucess"
		return response, true
	}

}
func Transaction(ori int, dest int, Amount float64) (string, bool) {
	msg := "Sucess"
	str, ret := CaptureId(ori)
	if !ret {
		return str, false
	}
	str, result := CaptureId(dest)
	if !result {
		return str, false
	}
	if ori == dest {
		msg = "The ids must be different"
		return msg, false
	}

	for i := 0; i < len(DataAcc); i++ {
		str, resp := Validator(DataAcc[i].Balance, Amount)
		if DataAcc[i].ID == ori {
			if !resp {
				return str, false
			}
			DataAcc[i].Balance -= Amount
		}
		if DataAcc[i].ID == dest {
			DataAcc[i].Balance += Amount

		}
	}
	return msg, true
}

func TransferringUnique(d domain.Transfer) (string, bool) { // rota deve ser put

	Transf := domain.Transfer{
		ID:                     IDGeneratorTransaction(),
		Amount:                 d.Amount,
		Account_origin_id:      d.Account_origin_id,
		Account_destination_id: d.Account_destination_id,
		Created_at_:            time.Now(),
	}

	str, ret := Transaction(Transf.Account_origin_id, Transf.Account_destination_id, Transf.Amount)
	return str, ret

}
