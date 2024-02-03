package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
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
		logger.LogError(err)
		return repositories.ErrItemRepositoryCanNotCreateItem
	}

	return nil
}

func (r *ItemRepository) GetByID(id string) (*entities.Item, error) {
	var item entities.Item

	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		logger.LogError(err)
		return nil, repositories.ErrItemRepositoryItemNotFound
	}

	return &item, nil
}
