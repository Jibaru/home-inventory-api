package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) CreateUser(email, password string) (*entities.User, error) {
	user, err := entities.NewUser(email, password)
	if err != nil {
		return nil, err
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, errors.New("cannot create user")
	}

	return user, nil
}
