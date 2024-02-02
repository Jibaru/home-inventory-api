package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"gorm.io/gorm"
)

type ItemKeywordRepository struct {
	db *gorm.DB
}

func NewItemKeywordRepository(db *gorm.DB) *ItemKeywordRepository {
	return &ItemKeywordRepository{
		db,
	}
}

func (r *ItemKeywordRepository) CreateMany(itemKeywords []*entities.ItemKeyword) error {
	if err := r.db.Create(&itemKeywords).Error; err != nil {
		return repositories.ErrItemKeywordRepositoryCanNotCreateItemKeywords
	}

	return nil
}
