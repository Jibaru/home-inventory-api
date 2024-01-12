package dao

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"gorm.io/gorm"
)

type UserDAO struct {
	DB *gorm.DB
}

func (d *UserDAO) Create(user *entities.User) error {
	return d.DB.Create(user).Error
}
