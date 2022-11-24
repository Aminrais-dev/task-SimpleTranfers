package factory

import (
	userData "task/simpleTranfers/features/users/data"
	userDelivery "task/simpleTranfers/features/users/delivery"
	userUsecase "task/simpleTranfers/features/users/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)
}
