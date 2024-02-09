package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
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
		logger.LogError(err)
		return repositories.ErrAssetRepositoryCanNotCreateAsset
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

	if err := result.Error; err != nil {
		logger.LogError(err)
		return nil, repositories.ErrAssetRepositoryCanNotGetAssets
	}

	return assets, nil
}

func (r *AssetRepository) Delete(id string) error {
	if err := r.db.Delete(&entities.Asset{}, "id = ?", id).Error; err != nil {
		logger.LogError(err)
		return repositories.ErrAssetRepositoryCanNotDeleteAsset
	}

	return nil
}

func (r *AssetRepository) GetByQueryFilters(queryFilter repositories.QueryFilter) ([]*entities.Asset, error) {
	var assets []*entities.Asset
	result := applyFilters(r.db, queryFilter).Find(&assets)

	if err := result.Error; err != nil {
		logger.LogError(err)
		return nil, repositories.ErrorAssetRepositoryCanNotGetByQueryFilters
	}

	return assets, nil
}
