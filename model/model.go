package model

import (
	"task/simpleTranfers/features/transactions"
	"task/simpleTranfers/features/users"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Account_no int       `gorm:"autoIncrement;primaryKey"`
	Name       string    `gorm:"type:char(50)"`
	Balance    float64   `gorm:"type:float"`
	Created_at time.Time `gorm:"autoCreateTime"`
}

type Transaction struct {
	Transaction_id uuid.UUID `gorm:"type:char(36);primary_key"`
	Amount         float64   `gorm:"type:float"`
	Credit_account int       `gorm:"index"`
	Dedit_account  int       `gorm:"index"`
	Created_at     time.Time `gorm:"autoCreateTime"`
}

func (data *User) ToUserCore() users.UserCore {
	return users.UserCore{
		Account_no: data.Account_no,
		Name:       data.Name,
		Balance:    data.Balance,
		Created_at: data.Created_at,
	}
}

func ToUserCoreList(data []User) []users.UserCore {

	var res []users.UserCore
	for key := range data {
		res = append(res, data[key].ToUserCore())
	}

	return res

}

func ToUserModel(data users.UserCore) User {
	return User{
		Account_no: data.Account_no,
		Name:       data.Name,
		Balance:    data.Balance,
		Created_at: data.Created_at,
	}

}

func (data *Transaction) ToTransactionCore() transactions.TransactionCore {
	return transactions.TransactionCore{
		Transaction_id: data.Transaction_id,
		Amount:         data.Amount,
		Credit_account: data.Credit_account,
		Dedit_account:  data.Dedit_account,
		Created_at:     data.Created_at,
	}
}

func ToTransactionCoreList(data []Transaction) []transactions.TransactionCore {

	var res []transactions.TransactionCore
	for key := range data {
		res = append(res, data[key].ToTransactionCore())
	}

	return res

}

func ToTransactionModel(data transactions.TransactionCore) Transaction {
	return Transaction{
		Transaction_id: data.Transaction_id,
		Amount:         data.Amount,
		Credit_account: data.Credit_account,
		Dedit_account:  data.Dedit_account,
		Created_at:     data.Created_at,
	}

}
