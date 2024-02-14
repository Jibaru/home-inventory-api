package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/jibaru/home-inventory-api/m/notifier"
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
		notifier.NotifyError(err)
		return repositories.ErrItemKeywordRepositoryCanNotCreateItemKeywords
	}

	return nil
}

func (r *ItemKeywordRepository) DeleteByItemID(itemID string) error {
	if err := r.db.Where("item_id = ?", itemID).Delete(&entities.ItemKeyword{}).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrItemKeywordRepositoryCanNotDeleteByItemID
	}

	return nil
}
