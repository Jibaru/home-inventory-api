package gorm

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/jibaru/home-inventory-api/m/notifier"
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
	if err := r.db.Create(user).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrUserRepositoryCanNotCreateUser
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	user := &entities.User{}

	err := r.db.First(user, "email = ?", email).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.LogError(err)
		return nil, repositories.ErrUserRepositoryUserNotFound
	}

	if err != nil {
		return nil, repositories.ErrUserRepositoryCanNotGetUser
	}

	return user, nil
}
