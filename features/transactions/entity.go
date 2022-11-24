package transactions

import (
	"time"

	"github.com/google/uuid"
)

type TransactionCore struct {
	Transaction_id uuid.UUID
	Amount         float64
	Credit_account int
	Dedit_account  int
	Created_at     time.Time
}

type Filter struct {
	Credit_account int
	Dedit_account  int
}

type dataInterface interface {
	InsertTransaction(data TransactionCore) TransactionCore
	SelectListTransaction() []TransactionCore
	SelectTransaction(transaction_id string) TransactionCore
	SelectTransactionBySearch(search Filter) []TransactionCore
}

type usecaseInterface interface {
	PostTransaction(data TransactionCore) TransactionCore
	GetListTransaction() []TransactionCore
	Getransaction(transaction_id string) TransactionCore
	GetTransactionBySearch(search Filter) []TransactionCore
}
