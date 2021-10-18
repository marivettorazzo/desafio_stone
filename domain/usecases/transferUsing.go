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
func CaptureId(id int) bool {
	for i := 0; i < len(DataAcc); i++ {
		if DataAcc[i].ID == id {
			return true
		}
	}
	return false
}

func Validator(balance float64, Amount float64) bool {

	switch {
	case balance < Amount:
		return false
	case Amount < 0:
		return false
	case Amount == 0:
		return false
	default:
		return true
	}

}
func Transaction(ori int, dest int, Amount float64) bool {
	if !CaptureId(ori) {
		return false
	}
	if !CaptureId(dest) {
		return false
	}
	for i := 0; i < len(DataAcc); i++ {
		if DataAcc[i].ID == ori {
			if !Validator(DataAcc[i].Balance, Amount) {
				return false
			}
			DataAcc[i].Balance -= Amount
		}
		if DataAcc[i].ID == dest {
			DataAcc[i].Balance += Amount
		}
	}
	return true
}

func TransferringUnique(d domain.Transfer) bool { // rota deve ser put

	Transf := domain.Transfer{
		ID:                     IDGeneratorTransaction(),
		Amount:                 d.Amount,
		Account_origin_id:      d.Account_origin_id,
		Account_destination_id: d.Account_destination_id,
		Created_at_:            time.Now(),
	}
	return Transaction(Transf.Account_origin_id, Transf.Account_destination_id, Transf.Amount)

}
