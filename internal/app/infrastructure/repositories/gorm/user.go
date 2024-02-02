package gorm

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	user := &entities.User{}

	err := r.db.First(user, "email = ?", email).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repositories.ErrUserRepositoryUserNotFound
	}

	if err != nil {
		return nil, repositories.ErrUserRepositoryCanNotGetUser
	}

	return user, nil
}
