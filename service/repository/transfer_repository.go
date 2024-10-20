package repository

import (
	"github.com/jinzhu/gorm"
	"mnc-test/model"
)

type TransferRepository interface {
	InsertTransfer(tf *model.Transfer) (*model.Transfer, error)
}

type transferRepository struct {
	db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) TransferRepository {
	return &transferRepository{db}
}

func (r *transferRepository) InsertTransfer(tf *model.Transfer) (*model.Transfer, error) {
	err := r.db.Create(&tf).Error
	if err != nil {
		return nil, err
	}

	return tf, nil
}
