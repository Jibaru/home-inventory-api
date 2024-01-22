package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
)

type AuthService struct {
	userRepository repositories.UserRepository
	tokenGenerator services.TokenGenerator
}

func NewAuthService(
	userRepository repositories.UserRepository,
	tokenGenerator services.TokenGenerator,
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

func (s *AuthService) ParseAuthentication(token string) (
	*struct {
		ID    string
		Email string
	},
	error,
) {
	data, err := s.tokenGenerator.ParseToken(token)
	if err != nil {
		return nil, err
	}

	return data, nil
}
