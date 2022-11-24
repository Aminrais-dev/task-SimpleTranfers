package delivery

import (
	"strings"
	"task/simpleTranfers/features/users"
)

type Response struct {
	Account_no int     `json:"account_no" form:"account_no"`
	Name       string  `json:"name" form:"name"`
	Balance    float64 `json:"balance" form:"balance"`
	Created_at string  `json:"created_at" form:"created_at"`
}

type ResponseAccountNumber struct {
	Account_no int `json:"account_no" form:"account_no"`
}

func toResponse(data users.UserCore) Response {

	dateString := data.Created_at.String()
	var timeSplit = strings.Split(dateString, ".")

	return Response{
		Account_no: data.Account_no,
		Name:       data.Name,
		Balance:    data.Balance,
		Created_at: timeSplit[0],
	}
}

func toResponseList(data []users.UserCore) []Response {

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

func toResponseNumber(data []users.UserCore) []ResponseAccountNumber {

	var dataResponseNumber []ResponseAccountNumber
	for key := range data {
		dataResponseNumber = append(dataResponseNumber, ResponseAccountNumber{
			Account_no: data[key].Account_no,
		})
	}

	return dataResponseNumber

}
