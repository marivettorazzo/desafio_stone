package usecases

import (
	"errors"
	"math/rand"
	"time"

	"exemple.com/DesafioStone/desafio_stone/domain"
)

var (
	DataTran []domain.Transfer
)

func IDGeneratorTransaction() int { // Transfer entity function that generates transaction ID
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(50)
	return value
}
func CaptureId(id int) error {

	for i := 0; i < len(DataAcc); i++ {
		if DataAcc[i].ID == id {

			return nil
		}
	}

	return errors.New("invalid id")
}

func Validator(balance float64, Amount float64) error {

	switch {
	case balance < Amount:

		return errors.New("not enough balance for transaction")
	case Amount < 0:

		return errors.New("invalid transfer value")

	case Amount == 0:

		return errors.New("invalid transfer value")
	default:

		return nil
	}

}
func Transaction(ori int, dest int, Amount float64) error {
	var ret error
	ret = CaptureId(ori)
	if ret != nil {
		return errors.New(ret.Error())
	}
	ret = CaptureId(dest)
	if ret != nil {
		return errors.New(ret.Error())
	}
	if ori == dest {

		return errors.New("the ids must be different")
	}

	for i := 0; i < len(DataAcc); i++ {
		if DataAcc[i].ID == ori {
			resp := Validator(DataAcc[i].Balance, Amount)
			if resp != nil {
				return errors.New(resp.Error())
			}
			DataAcc[i].Balance -= Amount
		}
		if DataAcc[i].ID == dest {
			DataAcc[i].Balance += Amount

		}
	}
	return nil
}

func TransferringUnique(d domain.Transfer) error { // rota deve ser put

	Transf := domain.Transfer{
		ID:                     IDGeneratorTransaction(),
		Amount:                 d.Amount,
		Account_origin_id:      d.Account_origin_id,
		Account_destination_id: d.Account_destination_id,
		Created_at_:            time.Now(),
	}

	ret := Transaction(Transf.Account_origin_id, Transf.Account_destination_id, Transf.Amount)
	if ret != nil {
		return errors.New(ret.Error())
	}
	DataTran = append(DataTran, Transf)
	return ret

}
func Transfers() error {
	if len(DataTran) == 0 {
		return errors.New("there are no transactions to show")
	}
	return nil

}
