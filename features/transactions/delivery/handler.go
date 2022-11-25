package delivery

import (
	"task/simpleTranfers/features/transactions"
	"task/simpleTranfers/utils/helper"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	service transactions.UsecaseInterface
}

func New(e *echo.Echo, data transactions.UsecaseInterface) {
	handler := &Delivery{
		service: data,
	}

	e.POST("/transaction", handler.CreateTransaction)
	e.GET("/transaction/list", handler.GetListTransaction)
	e.GET("/transaction/:transaction_id", handler.GetTransaction)
	e.POST("/transaction/search", handler.GetTransactionBySearch)

}

func (delivery *Delivery) CreateTransaction(c echo.Context) error {

	var req Request
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(400, "error bind"))
	}

	if req.Credit_account == req.Debit_account {
		return c.JSON(400, helper.FailedResponseHelper(400, "Cannot transfer if the sender and recipient are the same"))
	}

	data, errInt := delivery.service.PostTransaction(req.toCore())
	if errInt == -1 {
		return c.JSON(400, helper.FailedResponseHelper(400, "please input credit, debit and amount"))
	} else if errInt == -2 {
		return c.JSON(500, helper.FailedResponseHelper(500, "error server generate transaction_id"))
	} else if errInt == -3 {
		return c.JSON(404, helper.FailedResponseHelper(404, "credit_account not found"))
	} else if errInt == -4 {
		return c.JSON(404, helper.FailedResponseHelper(404, "debit_account not found"))
	} else if errInt == -5 {
		return c.JSON(400, helper.FailedResponseHelper(400, "the balance amount is smaller than the transfer amount"))
	} else if errInt == -6 {
		return c.JSON(500, helper.FailedResponseHelper(500, "a server error occurred"))
	}

	return c.JSON(201, helper.SuccessDataResponseHelper(201, "success to create transaction", toResponse(data)))

}

func (delivery *Delivery) GetListTransaction(c echo.Context) error {

	data, errInt := delivery.service.GetListTransaction()
	if errInt == -1 {
		return c.JSON(500, helper.FailedResponseHelper(500, "a server error occurred"))
	} else if data == nil {
		return c.JSON(404, helper.FailedResponseHelper(404, "data is still empty"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(200, "success get to list transaction", toResponseList(data)))

}

func (delivery *Delivery) GetTransaction(c echo.Context) error {

	param := c.Param("transaction_id")

	data := delivery.service.GetTransaction(param)
	if data.Credit_account < 1 {
		return c.JSON(404, helper.FailedResponseHelper(404, "data is still empty"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(200, "success get to list transaction", toResponse(data)))

}

func (delivery *Delivery) GetTransactionBySearch(c echo.Context) error {

	var req FilterSearch
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(400, "error bind"))
	}

	data, errInt := delivery.service.GetTransactionBySearch(req.toFilter())
	if errInt == -1 {
		return c.JSON(400, helper.FailedResponseHelper(400, "please input credit and debit number(must be positive)"))
	} else if data == nil {
		return c.JSON(404, helper.FailedResponseHelper(404, "data not found"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(200, "success get to list transaction", toResponseList(data)))

}
