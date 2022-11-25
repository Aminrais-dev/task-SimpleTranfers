package data

import (
	"task/simpleTranfers/features/transactions"
	"task/simpleTranfers/model"

	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) transactions.DataInterface {
	return &Storage{
		db: db,
	}
}

func (storage *Storage) SelectAndCheckAmount(credit_account, debit_account int) (transactions.Amount, int) {

	var dataCredit model.User
	txCredit := storage.db.First(&dataCredit, "account_no = ?", credit_account)
	if txCredit.Error != nil {
		return transactions.Amount{}, -1
	}

	var dataDebit model.User
	txDebit := storage.db.First(&dataDebit, "account_no = ?", debit_account)
	if txDebit.Error != nil {
		return transactions.Amount{}, -2
	}

	return transactions.Amount{Amount_credit: dataCredit.Balance, Amount_debit: dataDebit.Balance}, 1

}

func (storage *Storage) InsertTransaction(data transactions.TransactionCore, amount transactions.Amount) transactions.TransactionCore {

	tx := storage.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return transactions.TransactionCore{}
	}

	errSpend := tx.Model(&model.User{}).Where("account_no = ?", data.Debit_account).Update("balance", amount.Amount_debit-data.Amount)
	if errSpend.Error != nil {
		tx.Rollback()
		return transactions.TransactionCore{}
	}

	errReceive := tx.Model(&model.User{}).Where("account_no = ?", data.Credit_account).Update("balance", amount.Amount_credit+data.Amount)
	if errReceive.Error != nil {
		tx.Rollback()
		return transactions.TransactionCore{}
	}

	var dataCreate = model.ToTransactionModel(data)
	errInsert := tx.Create(&dataCreate)
	if errInsert.Error != nil {
		tx.Rollback()
		return transactions.TransactionCore{}
	}

	tx.Commit()

	return dataCreate.ToTransactionCore()

}

func (storege *Storage) SelectListTransaction() ([]transactions.TransactionCore, int) {

	var data []model.Transaction
	tx := storege.db.Find(&data).Order("created_at DESC")
	if tx.Error != nil {
		return nil, -1
	}

	return model.ToTransactionCoreList(data), 1

}

func (storage *Storage) SelectTransaction(transaction_id string) transactions.TransactionCore {

	var data model.Transaction
	tx := storage.db.First(&data, "transaction_id = ?", transaction_id)
	if tx.Error != nil {
		return transactions.TransactionCore{}
	}

	return data.ToTransactionCore()

}

func (storage *Storage) SelectTransactionBySearch(search transactions.Filter) []transactions.TransactionCore {

	var data []model.Transaction
	if search.Credit_account != 0 && search.Debit_account != 0 {
		tx := storage.db.Find(&data, "credit_account = ? AND debit_account = ?", search.Credit_account, search.Debit_account)
		if tx.Error != nil {
			return nil
		}
		return model.ToTransactionCoreList(data)
	} else if search.Credit_account != 0 && search.Debit_account == 0 {
		tx := storage.db.Find(&data, "credit_account = ? ", search.Credit_account)
		if tx.Error != nil {
			return nil
		}
		return model.ToTransactionCoreList(data)
	} else if search.Credit_account == 0 && search.Debit_account != 0 {
		tx := storage.db.Find(&data, "debit_account = ?", search.Debit_account)
		if tx.Error != nil {
			return nil
		}
		return model.ToTransactionCoreList(data)
	}

	return nil

}
