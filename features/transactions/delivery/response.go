package delivery

import (
	"strings"
	"task/simpleTranfers/features/transactions"

	"github.com/google/uuid"
)

type Response struct {
	Transaction_id uuid.UUID `json:"transaction_id" form:"transaction_id"`
	Amount         float64   `json:"amount" form:"amount"`
	Credit_account int       `json:"credit_account" form:"credit_account"`
	Debit_account  int       `json:"debit_account" form:"debit_account"`
	Created_at     string    `json:"created_at" form:"created_at"`
}

func toResponse(data transactions.TransactionCore) Response {

	dateString := data.Created_at.String()
	var timeSplit = strings.Split(dateString, ".")

	return Response{
		Transaction_id: data.Transaction_id,
		Amount:         data.Amount,
		Credit_account: data.Credit_account,
		Debit_account:  data.Debit_account,
		Created_at:     timeSplit[0],
	}

}

func toResponseList(data []transactions.TransactionCore) []Response {

	var dataResponse []Response
	for key := range data {

		dateString := data[key].Created_at.String()
		var timeSplit = strings.Split(dateString, ".")
		var dataAppend = toResponse(data[key])

		dataAppend.Created_at = timeSplit[0]
		dataResponse = append(dataResponse, dataAppend)
	}

	return dataResponse

}
