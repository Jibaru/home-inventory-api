package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
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
		logger.LogError(err)
		return repositories.ErrItemKeywordRepositoryCanNotCreateItemKeywords
	}

	return nil
}
