package repository

import (
	"github.com/jinzhu/gorm"
	"mnc-test/model"
)

type PaymentRepository interface {
	InsertPayment(pay *model.Payment) (*model.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) InsertPayment(pay *model.Payment) (*model.Payment, error) {
	err := r.db.Create(&pay).Error
	if err != nil {
		return nil, err
	}

	return pay, nil
}
