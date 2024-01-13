package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
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
