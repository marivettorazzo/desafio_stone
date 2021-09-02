package controllers

import (
	"math/rand"
	"time"

	"exemple.com/DesafioStone/src/model"
)

var (
	DataAcc  []model.Account
	DataTran []model.Transfer
)

func IDGeneratorTransaction() int { // Transfer entity function that generates transaction ID
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(50)
	return value
}
func CaptureId(id int) bool {
	x := true
	for i := 0; i < len(DataAcc); i++ {

		if DataAcc[i].ID == id {
			return x
		}
	}
	return !x

}
func TransactionOccurs(ori, dest int, Amount float64) bool {

	for i := 0; i < len(DataAcc); i++ {

		if DataAcc[i].ID == ori {
			valorNaconta := DataAcc[i].Balance

			if valorNaconta < Amount {

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

func TransferringUnique(d model.Transfer) string { // rota deve ser put

	positive := true
	ori := CaptureId(d.Account_origin_id)
	dest := CaptureId(d.Account_destination_id)

	if ori != positive || dest != positive {
		err := "id does not exist"
		return err
	}

	Transf := model.Transfer{
		ID:                     IDGeneratorTransaction(),
		Amount:                 d.Amount,
		Account_origin_id:      d.Account_origin_id,
		Account_destination_id: d.Account_destination_id,
		Created_at_:            time.Now(),
	}

	ret := ValidateTransf(Transf.Account_origin_id, Transf.Account_destination_id, Transf.Amount, Transf)

	return ret

}
func ValidateTransf(ori, dest int, Amount float64, t model.Transfer) string {

	positive := true
	x := TransactionOccurs(ori, dest, Amount)
	if x == positive {
		ret := "Transaction performed successfully"
		DataTran = append(DataTran, t)
		return ret
	}
	ret := "Insufficient balance of source account"
	return ret
}
