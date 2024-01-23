package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"gorm.io/gorm"
)

type BoxRepository struct {
	db *gorm.DB
}

func NewBoxRepository(db *gorm.DB) *BoxRepository {
	return &BoxRepository{db}
}

func (r *BoxRepository) Create(box *entities.Box) error {
	if err := r.db.Create(box).Error; err != nil {
		return repositories.ErrCanNotCreateBox
	}

	return nil
}
