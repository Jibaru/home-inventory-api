package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/jibaru/home-inventory-api/m/notifier"
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
		logger.LogError(err)
		notifier.NotifyError(err)
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
		notifier.NotifyError(err)
		return repositories.ErrBoxRepositoryCanBotCreateBoxItem
	}

	return nil
}

func (r *BoxRepository) UpdateBoxItem(boxItem *entities.BoxItem) error {
	if err := r.db.Save(boxItem).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrBoxRepositoryCanNotUpdateBoxItem
	}

	return nil
}

func (r *BoxRepository) CreateBoxTransaction(boxTransaction *entities.BoxTransaction) error {
	if err := r.db.Create(boxTransaction).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrBoxRepositoryCanNotCreateBoxTransaction
	}

	return nil
}

func (r *BoxRepository) DeleteBoxItem(boxID string, itemID string) error {
	if err := r.db.Where("box_id = ? AND item_id = ?", boxID, itemID).Delete(&entities.BoxItem{}).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
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
		return nil, repositories.ErrBoxRepositoryCanNotGetByQueryFilters
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
		return 0, repositories.ErrBoxRepositoryCanNotCountByQueryFilters
	}

	return count, nil
}

func (r *BoxRepository) DeleteBoxItemsByBoxID(boxID string) error {
	if err := r.db.Where("box_id = ?", boxID).Delete(&entities.BoxItem{}).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrBoxRepositoryCanNotDeleteBoxItemsByBoxID
	}

	return nil
}

func (r *BoxRepository) DeleteBoxTransactionsByBoxID(boxID string) error {
	if err := r.db.Where("box_id = ?", boxID).Delete(&entities.BoxTransaction{}).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrBoxRepositoryCanNotDeleteBoxTransactionsByBoxID
	}

	return nil
}

func (r *BoxRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Box{}).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrBoxRepositoryCanNotDeleteBox
	}

	return nil
}

func (r *BoxRepository) GetByID(id string) (*entities.Box, error) {
	var box entities.Box
	if err := r.db.Where("id = ?", id).First(&box).Error; err != nil {
		logger.LogError(err)
		return nil, repositories.ErrBoxRepositoryCanNotGetByID
	}

	return &box, nil
}

func (r *BoxRepository) Update(box *entities.Box) error {
	if err := r.db.Save(box).Error; err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrBoxRepositoryCanNotUpdate
	}

	return nil
}
