package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrUserRepositoryCanNotCreateUser     = errors.New("can not create user")
	ErrUserRepositoryCanNotGetUser        = errors.New("can not get user")
	ErrUserRepositoryCanNotGetUserByBoxID = errors.New("can not get user by box id")
	ErrUserRepositoryUserNotFound         = errors.New("user not found")
)

type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	GetUserByBoxID(boxID string) (*entities.User, error)
}
