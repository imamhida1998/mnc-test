package repository

import (
	"github.com/jinzhu/gorm"
	"mnc-test/model"
)

type TransactionRepository interface {
	InsertTransaction(topup *model.Transaction) (*model.Transaction, error)
	FindByUserId(UserId string) (*[]model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) InsertTransaction(topup *model.Transaction) (*model.Transaction, error) {
	err := r.db.Create(&topup).Error
	if err != nil {
		return nil, err
	}

	return topup, nil
}

func (r *transactionRepository) FindByUserId(UserId string) (*[]model.Transaction, error) {
	var tx []model.Transaction

	err := r.db.Where("user_id = ?", UserId).Order("created_at DESC").Find(&tx).Error
	if err != nil {
		return nil, err
	}

	return &tx, nil
}
