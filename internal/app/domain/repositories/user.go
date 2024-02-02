package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrUserRepositoryCanNotGetUser = errors.New("can not get user")
	ErrUserRepositoryUserNotFound  = errors.New("user not found")
)

type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}
