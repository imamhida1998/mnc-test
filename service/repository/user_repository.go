package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"mnc-test/model"
	"mnc-test/model/request"
)

type UserRepository interface {
	InsertUser(user *model.User) (*model.User, error)
	FindByPhoneNumber(input request.Login) (*model.User, error)
	FindById(userId string) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByPhoneNumber(input request.Login) (*model.User, error) {
	var user model.User

	err := r.db.Where("phone_number = ? and pin = ?", input.PhoneNumber, input.Pin).Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Phone Number and PIN doesn’t match.")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindById(userId string) (*model.User, error) {
	var user model.User

	err := r.db.Where("user_id = ?", userId).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := r.db.Model(&model.User{}).Where("user_id = ?", user.UserId).Updates(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
