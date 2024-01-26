package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrCanNotCreateAsset = errors.New("can not create asset")
	ErrCanNotGetAssets   = errors.New("can not get assets")
	ErrCanNotDeleteAsset = errors.New("can not delete asset")
)

type AssetRepository interface {
	Create(asset *entities.Asset) error
	FindByEntity(entity entities.Entity, page *PageFilter) ([]*entities.Asset, error)
	Delete(id string) error
}
