package repository

import (
	"github.com/jinzhu/gorm"
	"mnc-test/model"
)

type TopUpRepository interface {
	InsertTopUp(topup *model.TopUp) (*model.TopUp, error)
}

type topUpRepository struct {
	db *gorm.DB
}

func NewTopUpRepository(db *gorm.DB) TopUpRepository {
	return &topUpRepository{db}
}

func (r *topUpRepository) InsertTopUp(topup *model.TopUp) (*model.TopUp, error) {
	err := r.db.Create(&topup).Error
	if err != nil {
		return nil, err
	}

	return topup, nil
}
