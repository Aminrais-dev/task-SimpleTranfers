package usecase

import (
	"errors"
	"strings"
	"task/simpleTranfers/features/users"
)

type Service struct {
	storage users.DataInterface
}

func New(data users.DataInterface) users.UsecaseInterface {
	return &Service{
		storage: data,
	}
}

func (service *Service) PostAccount(data users.UserCore) (users.UserCore, int) {

	if data.Name == "" || data.Balance < 0 {
		return users.UserCore{}, -1
	}

	checkName := service.storage.CheckName(data.Name)
	if checkName == true {
		return users.UserCore{}, -2
	}

	dataRes := service.storage.InsertAccount(data)
	if dataRes.Name == "" {
		return users.UserCore{}, -3
	}

	return dataRes, 0

}

func (service *Service) GetListAccount(queryParam string) ([]users.UserCore, error) {

	var valid = strings.ToLower(queryParam)
	data := service.storage.SelectListAccount(valid)
	if data == nil {
		return nil, errors.New("error query")
	}

	return data, nil

}

func (service *Service) GetAccount(param int) users.UserCore {

	data := service.storage.SelectAccount(param)

	return data

}
