package delivery

import (
	"strconv"
	"task/simpleTranfers/features/users"
	"task/simpleTranfers/utils/helper"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	service users.UsecaseInterface
}

func New(e *echo.Echo, data users.UsecaseInterface) {
	handler := &Delivery{
		service: data,
	}

	e.POST("/account", handler.CreateAccount)
	e.GET("/account/list", handler.GetAccountList)
	e.GET("/account/:account_no", handler.GetAccount)

}

func (delivery *Delivery) CreateAccount(c echo.Context) error {

	var req Request
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(400, "error bind"))
	}

	data, errorInt := delivery.service.PostAccount(req.toCore())
	if errorInt == -1 {
		return c.JSON(400, helper.FailedResponseHelper(400, "please input name and balance"))
	} else if errorInt == -2 {
		return c.JSON(400, helper.FailedResponseHelper(400, "name already exists"))
	} else if errorInt == -3 {
		return c.JSON(500, helper.FailedResponseHelper(500, "a server error occurred"))
	}

	return c.JSON(201, helper.SuccessDataResponseHelper(201, "success to create account", toResponse(data)))

}

func (delivery *Delivery) GetAccountList(c echo.Context) error {

	queryParam := c.QueryParam("details")
	if queryParam != "true" && queryParam != "false" {
		return c.JSON(400, helper.FailedResponseHelper(400, "please input query params details=true or false"))
	}

	data, err := delivery.service.GetListAccount(queryParam)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper(500, "a server error occurred"))
	}

	if data[0].Name == "" {
		return c.JSON(200, helper.SuccessDataResponseHelper(200, "success to list account", toResponseNumber(data)))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(200, "success to list account", toResponseList(data)))

}

func (delivery *Delivery) GetAccount(c echo.Context) error {

	param, err := strconv.Atoi(c.Param("account_no"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(400, "param must be number"))
	}

	data := delivery.service.GetAccount(param)
	if data.Name == "" {
		return c.JSON(404, helper.FailedResponseHelper(404, "account not found"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(200, "success to list account", toResponse(data)))

}
