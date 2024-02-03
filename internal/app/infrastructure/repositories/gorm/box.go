package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
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
		return repositories.ErrBoxRepositoryCanNotCreateBox
	}

	return nil
}

func (r *BoxRepository) GetBoxItem(boxID string, itemID string) (*entities.BoxItem, error) {
	var boxItem entities.BoxItem

	if err := r.db.Where("box_id = ? AND item_id = ?", boxID, itemID).First(&boxItem).Error; err != nil {
		logger.LogError(err)
		return nil, repositories.ErrBoxRepositoryBoxItemNotFound
	}

	return &boxItem, nil
}

func (r *BoxRepository) CreateBoxItem(boxItem *entities.BoxItem) error {
	if err := r.db.Create(boxItem).Error; err != nil {
		logger.LogError(err)
		return repositories.ErrBoxRepositoryCanBotCreateBoxItem
	}

	return nil
}

func (r *BoxRepository) UpdateBoxItem(boxItem *entities.BoxItem) error {
	if err := r.db.Save(boxItem).Error; err != nil {
		logger.LogError(err)
		return repositories.ErrBoxRepositoryCanNotUpdateBoxItem
	}

	return nil
}

func (r *BoxRepository) CreateBoxTransaction(boxTransaction *entities.BoxTransaction) error {
	if err := r.db.Create(boxTransaction).Error; err != nil {
		logger.LogError(err)
		return repositories.ErrBoxRepositoryCanNotCreateBoxTransaction
	}

	return nil
}

func (r *BoxRepository) DeleteBoxItem(boxID string, itemID string) error {
	if err := r.db.Where("box_id = ? AND item_id = ?", boxID, itemID).Delete(&entities.BoxItem{}).Error; err != nil {
		logger.LogError(err)
		return repositories.ErrBoxRepositoryCanNotDeleteBoxItem
	}

	return nil
}

func (r *BoxRepository) GetByQueryFilters(queryFilter repositories.QueryFilter, pageFilter *repositories.PageFilter) ([]*entities.Box, error) {
	var boxes []*entities.Box
	err := applyFilters(r.db, queryFilter).
		Joins("inner join rooms on boxes.room_id = rooms.id").
		Offset(pageFilter.Offset).
		Limit(pageFilter.Limit).
		Find(&boxes).
		Error

	if err != nil {
		logger.LogError(err)
		return nil, repositories.ErrorBoxRepositoryCanNotGetByQueryFilters
	}

	return boxes, nil
}

func (r *BoxRepository) CountByQueryFilters(queryFilter repositories.QueryFilter) (int64, error) {
	var count int64
	err := applyFilters(r.db.Model(&entities.Box{}), queryFilter).
		Joins("inner join rooms on boxes.room_id = rooms.id").
		Count(&count).
		Error

	if err != nil {
		logger.LogError(err)
		return 0, repositories.ErrorBoxRepositoryCanNotCountByQueryFilters
	}

	return count, nil
}
