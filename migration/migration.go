package migration

import (
	model "task/simpleTranfers/model"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(model.User{})
	db.AutoMigrate(model.Transaction{})
}
