package delivery

import (
	"task/simpleTranfers/features/transactions"

	"github.com/google/uuid"
)

type Request struct {
	Credit_account int     `json:"credit_account" form:"credit_account"`
	Debit_account  int     `json:"debit_account" form:"debit_account"`
	Amount         float64 `json:"amount" form:"amount"`
}

type FilterSearch struct {
	Filter []Search `json:"filter" form:"filter"`
}

type Search struct {
	Credit_account int `json:"credit_account" form:"credit_account"`
	Debit_account  int `json:"debit_account" form:"debit_account"`
}

func (req *Request) toCore() transactions.TransactionCore {
	transaction_id := uuid.New()

	return transactions.TransactionCore{
		Transaction_id: transaction_id,
		Credit_account: req.Credit_account,
		Debit_account:  req.Debit_account,
		Amount:         req.Amount,
	}
}

func (req *FilterSearch) toFilter() transactions.Filter {

	if len(req.Filter) == 1 {
		return transactions.Filter{
			Credit_account: req.Filter[0].Credit_account,
			Debit_account:  req.Filter[0].Debit_account,
		}
	}

	if len(req.Filter) == 2 {
		if req.Filter[0].Credit_account > 0 {
			return transactions.Filter{
				Credit_account: req.Filter[0].Credit_account,
				Debit_account:  req.Filter[1].Debit_account,
			}

		} else {
			return transactions.Filter{
				Credit_account: req.Filter[1].Credit_account,
				Debit_account:  req.Filter[0].Debit_account,
			}
		}
	}

	return transactions.Filter{
		Credit_account: 0,
		Debit_account:  0,
	}

}
