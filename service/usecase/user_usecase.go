package usecase

import (
	"github.com/google/uuid"
	"mnc-test/model"
	"mnc-test/model/request"
	"mnc-test/model/response"
	"mnc-test/service/repository"
)

type UserUsecase interface {
	Register(input request.Register) (*response.Registration, error)
	GetAcoountByPhoneNumber(input request.Login) (*model.User, error)
	UpdateProfiles(data, user *model.User) (*response.User, error)
}
type userUsercase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsercase{repository}
}

func (s *userUsercase) Register(input request.Register) (*response.Registration, error) {
	id, _ := uuid.NewV7()

	users := &model.User{
		UserId:      id.String(),
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		Pin:         input.Pin,
	}

	newUser, err := s.repository.InsertUser(users)
	if err != nil {
		return nil, err
	}

	return &response.Registration{
		UserID:      newUser.UserId,
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		PhoneNumber: newUser.PhoneNumber,
		Address:     newUser.Address,
		CreatedDate: newUser.CreatedAt.Format("2006-1-2 15:04:05"),
	}, nil
}

func (s *userUsercase) GetAcoountByPhoneNumber(input request.Login) (*model.User, error) {
	return s.repository.FindByPhoneNumber(input)
}

func (s *userUsercase) UpdateProfiles(data, user *model.User) (*response.User, error) {

	if data.FirstName != "" {
		user.FirstName = data.FirstName
	}

	if data.LastName != "" {
		user.LastName = data.LastName
	}

	if data.Address != "" {
		user.Address = data.Address
	}
	res, err := s.repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &response.User{
		UserId:    res.UserId,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Address:   res.Address,
		UpdatedAt: res.UpdatedAt,
	}, nil

}
