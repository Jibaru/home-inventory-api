package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrAssetRepositoryCanNotCreateAsset = errors.New("can not create asset")
	ErrAssetRepositoryCanNotDeleteAsset = errors.New("can not delete asset")
	ErrAssetRepositoryCanNotGetAssets   = errors.New("can not get assets")
)

type AssetRepository interface {
	Create(asset *entities.Asset) error
	FindByEntity(entity entities.Entity, page *PageFilter) ([]*entities.Asset, error)
	Delete(id string) error
}
