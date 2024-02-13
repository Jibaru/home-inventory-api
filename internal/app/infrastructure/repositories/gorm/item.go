package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/jibaru/home-inventory-api/m/notifier"
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
		notifier.NotifyError(err)
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

func (r *ItemRepository) GetByQueryFilters(queryFilter repositories.QueryFilter, pageFilter *repositories.PageFilter) ([]*entities.Item, error) {
	var items []*entities.Item
	err := applyFilters(r.db, queryFilter).
		Joins("inner join item_keywords on item_keywords.item_id = items.id").
		Offset(pageFilter.Offset).
		Limit(pageFilter.Limit).
		Preload("Keywords").
		Find(&items).
		Error

	if err != nil {
		logger.LogError(err)
		return nil, repositories.ErrItemRepositoryCanNotGetByQueryFilters
	}

	return items, nil
}

func (r *ItemRepository) CountByQueryFilters(queryFilter repositories.QueryFilter) (int64, error) {
	var count int64
	err := applyFilters(r.db, queryFilter).
		Joins("inner join item_keywords on item_keywords.item_id = items.id").
		Model(&entities.Item{}).
		Count(&count).
		Error

	if err != nil {
		logger.LogError(err)
		return 0, repositories.ErrItemRepositoryCanNotCountByQueryFilters
	}

	return count, nil
}
