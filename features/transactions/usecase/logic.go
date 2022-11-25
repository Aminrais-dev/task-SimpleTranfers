package usecase

import (
	"fmt"
	"task/simpleTranfers/features/transactions"

	"github.com/google/uuid"
)

type Service struct {
	storage transactions.DataInterface
}

func New(data transactions.DataInterface) transactions.UsecaseInterface {
	return &Service{
		storage: data,
	}
}

func (service *Service) PostTransaction(data transactions.TransactionCore) (transactions.TransactionCore, int) {

	if data.Debit_account < 1 || data.Credit_account < 1 || data.Amount <= 0 {
		return transactions.TransactionCore{}, -1
	} else if data.Transaction_id == uuid.Nil {
		return transactions.TransactionCore{}, -2
	}

	amount, errInt := service.storage.SelectAndCheckAmount(data.Credit_account, data.Debit_account)
	if errInt == -1 {
		return transactions.TransactionCore{}, -3
	} else if errInt == -2 {
		return transactions.TransactionCore{}, -4
	} else if amount.Amount_debit < data.Amount {
		return transactions.TransactionCore{}, -5
	}

	fmt.Println(amount.Amount_debit, data.Amount)

	dataRes := service.storage.InsertTransaction(data, amount)
	if dataRes.Credit_account < 1 {
		return transactions.TransactionCore{}, -6
	}

	return dataRes, 0

}

func (service *Service) GetListTransaction() ([]transactions.TransactionCore, int) {

	data, errInt := service.storage.SelectListTransaction()

	return data, errInt

}

func (service *Service) GetTransaction(transaction_id string) transactions.TransactionCore {

	data := service.storage.SelectTransaction(transaction_id)

	return data

}

func (service *Service) GetTransactionBySearch(search transactions.Filter) ([]transactions.TransactionCore, int) {

	if search.Credit_account < 1 && search.Debit_account < 1 {
		return nil, -1
	}

	data := service.storage.SelectTransactionBySearch(search)

	return data, 1

}
