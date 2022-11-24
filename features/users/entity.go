package users

import (
	"time"
)

type UserCore struct {
	Account_no int
	Name       string
	Balance    float64
	Created_at time.Time
}

type GetNumber struct {
	Account_no int
}

type DataInterface interface {
	CheckName(name string) bool
	InsertAccount(data UserCore) UserCore
	SelectListAccount(queryParam string) []UserCore
	SelectAccount(param int) UserCore
}

type UsecaseInterface interface {
	PostAccount(data UserCore) (UserCore, int)
	GetListAccount(queryParam string) ([]UserCore, error)
	GetAccount(param int) UserCore
}

func GetNumberToCore(data []GetNumber) []UserCore {

	var res []UserCore
	for key := range data {
		res = append(res, UserCore{
			Account_no: data[key].Account_no,
		})
	}

	return res

}
