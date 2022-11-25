package transactions

import (
	"time"

	"github.com/google/uuid"
)

type TransactionCore struct {
	Transaction_id uuid.UUID
	Amount         float64
	Credit_account int
	Debit_account  int
	Created_at     time.Time
}

type Filter struct {
	Credit_account int
	Debit_account  int
}

type Amount struct {
	Amount_credit float64
	Amount_debit  float64
}

type DataInterface interface {
	SelectAndCheckAmount(credit_account, debit_account int) (Amount, int)
	InsertTransaction(data TransactionCore, amount Amount) TransactionCore
	SelectListTransaction() ([]TransactionCore, int)
	SelectTransaction(transaction_id string) TransactionCore
	SelectTransactionBySearch(search Filter) []TransactionCore
}

type UsecaseInterface interface {
	PostTransaction(data TransactionCore) (TransactionCore, int)
	GetListTransaction() ([]TransactionCore, int)
	GetTransaction(transaction_id string) TransactionCore
	GetTransactionBySearch(search Filter) ([]TransactionCore, int)
}
