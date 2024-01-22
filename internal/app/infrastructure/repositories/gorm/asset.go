package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"gorm.io/gorm"
)

type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{
		db,
	}
}

func (r *AssetRepository) Create(asset *entities.Asset) error {
	if err := r.db.Create(asset).Error; err != nil {
		return repositories.ErrCanNotCreateAsset
	}

	return nil
}

func (r *AssetRepository) FindByEntity(
	entity entities.Entity,
	page *repositories.PageFilter,
) ([]*entities.Asset, error) {
	var assets []*entities.Asset
	query := r.db.
		Where("entity_id = ?", entity.EntityID()).
		Where("entity_name = ?", entity.EntityName())

	if page != nil {
		query.Offset(page.Offset).Limit(page.Limit)
	}

	result := query.Find(&assets)

	if result.Error != nil {
		return nil, repositories.ErrCanNotGetAssets
	}

	return assets, nil
}
