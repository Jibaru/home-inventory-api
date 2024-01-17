package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/auth"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
)

type AuthService struct {
	userRepository repositories.UserRepository
	tokenGenerator auth.TokenGenerator
}

func NewAuthService(
	userRepository repositories.UserRepository,
	tokenGenerator auth.TokenGenerator,
) *AuthService {
	return &AuthService{
		userRepository,
		tokenGenerator,
	}
}

func (s *AuthService) Authenticate(email, password string) (
	*struct {
		User  *entities.User
		Token string
	},
	error,
) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !user.HasEqualPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, errors.New("can not authenticate")
	}

	return &struct {
		User  *entities.User
		Token string
	}{
		user,
		token,
	}, nil
}

func (s *AuthService) GenerateToken(user *entities.User) (string, error) {
	return s.tokenGenerator.GenerateToken(user.ID, user.Email)
}
