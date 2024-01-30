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

func (r *BoxRepository) GetBoxItem(boxID string, itemID string) (*entities.BoxItem, error) {
	var boxItem entities.BoxItem

	if err := r.db.Where("box_id = ? AND item_id = ?", boxID, itemID).First(&boxItem).Error; err != nil {
		return nil, repositories.ErrBoxRepositoryBoxItemNotFound
	}

	return &boxItem, nil
}

func (r *BoxRepository) CreateBoxItem(boxItem *entities.BoxItem) error {
	if err := r.db.Create(boxItem).Error; err != nil {
		return repositories.ErrBoxRepositoryCanBotCreateBoxItem
	}

	return nil
}

func (r *BoxRepository) UpdateBoxItem(boxItem *entities.BoxItem) error {
	if err := r.db.Save(boxItem).Error; err != nil {
		return repositories.ErrBoxRepositoryCanNotUpdateBoxItem
	}

	return nil
}

func (r *BoxRepository) CreateBoxTransaction(boxTransaction *entities.BoxTransaction) error {
	if err := r.db.Create(boxTransaction).Error; err != nil {
		return repositories.ErrBoxRepositoryCanNotCreateBoxTransaction
	}

	return nil
}
