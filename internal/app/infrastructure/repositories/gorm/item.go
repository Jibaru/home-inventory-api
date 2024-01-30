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

func (r *ItemRepository) GetByID(id string) (*entities.Item, error) {
	var item entities.Item

	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, repositories.ErrItemRepositoryItemNotFound
	}

	return &item, nil
}
