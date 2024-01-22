package services

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"os"
)

type AssetService struct {
	fileManager     services.FileManager
	assetRepository repositories.AssetRepository
}

func NewAssetService(
	fileManager services.FileManager,
	assetRepository repositories.AssetRepository,
) *AssetService {
	return &AssetService{
		fileManager,
		assetRepository,
	}
}

func (s *AssetService) CreateFromFile(
	file *os.File,
	entity entities.Entity,
) (*entities.Asset, error) {
	fileID, err := s.fileManager.Upload(file)
	if err != nil {
		return nil, err
	}

	asset, err := entities.NewAssetFromFile(file, fileID, entity)
	if err != nil {
		return nil, err
	}

	err = s.assetRepository.Create(asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (s *AssetService) GetUrl(asset *entities.Asset) string {
	return s.fileManager.GenerateUrl(asset.FileID, asset.Extension)
}

func (s *AssetService) GetByEntity(entity entities.Entity, pageFilter *PageFilter) ([]*entities.Asset, error) {
	var repositoryPageFilter *repositories.PageFilter

	if pageFilter != nil {
		repositoryPageFilter = &repositories.PageFilter{
			Offset: (pageFilter.Page - 1) * pageFilter.Size,
			Limit:  pageFilter.Size,
		}
	}

	assets, err := s.assetRepository.FindByEntity(entity, repositoryPageFilter)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
