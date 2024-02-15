package listeners

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	domain "github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"github.com/jibaru/home-inventory-api/m/logger"
)

type RollbackAssetListener struct {
	assetService *services.AssetService
}

func NewRollbackAssetListener(
	assetService *services.AssetService,
) *RollbackAssetListener {
	return &RollbackAssetListener{
		assetService: assetService,
	}
}

func (l *RollbackAssetListener) Handle(event domain.Event) {
	var asset *entities.Asset
	if e, ok := event.(domain.ItemNotCreatedEvent); ok {
		asset = &e.Asset
	} else if e, ok := event.(domain.ItemKeywordsNotCreatedEvent); ok {
		asset = &e.Asset
	}

	if asset != nil {
		err := l.assetService.Delete(asset)
		if err != nil {
			logger.LogError(err)
		}
	}
}
