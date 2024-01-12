package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/dao"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

type UserService struct {
	userDAO dao.UserDAO
}

func NewUserService(userDAO dao.UserDAO) *UserService {
	return &UserService{userDAO}
}

func (s *UserService) CreateUser(email, password string) (*entities.User, error) {
	user, err := entities.NewUser(email, password)
	if err != nil {
		return nil, err
	}

	err = s.userDAO.Create(user)
	if err != nil {
		return nil, errors.New("cannot create user")
	}

	return user, nil
}
