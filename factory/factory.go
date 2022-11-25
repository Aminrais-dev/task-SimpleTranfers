package factory

import (
	userData "task/simpleTranfers/features/users/data"
	userDelivery "task/simpleTranfers/features/users/delivery"
	userUsecase "task/simpleTranfers/features/users/usecase"

	transactionData "task/simpleTranfers/features/transactions/data"
	transactionDelivery "task/simpleTranfers/features/transactions/delivery"
	transactionUsecase "task/simpleTranfers/features/transactions/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	transactionDataFactory := transactionData.New(db)
	transactionUsecaseFactory := transactionUsecase.New(transactionDataFactory)
	transactionDelivery.New(e, transactionUsecaseFactory)

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)
}
