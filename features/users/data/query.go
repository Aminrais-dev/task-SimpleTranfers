package data

import (
	"task/simpleTranfers/features/users"
	"task/simpleTranfers/model"

	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataInterface {
	return &Storage{
		db: db,
	}
}

func (storage *Storage) CheckName(name string) bool {

	var user model.User
	tx := storage.db.First(&user, "name = ?", name)
	if tx.Error != nil || user.Name == "" {
		return false
	}

	return true

}

func (storage *Storage) InsertAccount(data users.UserCore) users.UserCore {

	var dataUser = model.ToUserModel(data)
	tx := storage.db.Create(&dataUser)
	if tx.Error != nil {
		return users.UserCore{}
	}

	return dataUser.ToUserCore()

}

func (storge *Storage) SelectListAccount(queryParam string) []users.UserCore {

	if queryParam == "true" {
		var userList []model.User
		tx := storge.db.Find(&userList)
		if tx.Error != nil {
			return nil
		}
		return model.ToUserCoreList(userList)

	} else if queryParam == "false" {
		var getNumber []users.GetNumber
		tx := storge.db.Model(model.User{}).Select("account_no as account_no").Scan(&getNumber)
		if tx.Error != nil {
			return nil
		}

		return users.GetNumberToCore(getNumber)
	}

	return nil

}

func (storge *Storage) SelectAccount(param int) users.UserCore {

	var user model.User
	storge.db.First(&user, "account_no = ? ", param)

	return user.ToUserCore()

}
