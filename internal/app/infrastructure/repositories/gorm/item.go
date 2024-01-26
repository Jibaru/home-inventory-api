package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db,
	}
}

func (r *ItemRepository) Create(item *entities.Item) error {
	if err := r.db.Create(item).Error; err != nil {
		return repositories.ErrCanNotCreateItem
	}

	return nil
}
