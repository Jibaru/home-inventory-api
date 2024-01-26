package services

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"github.com/stretchr/testify/mock"
	"os"
)

type AssetServiceInterface interface {
	CreateFromFile(file *os.File, entity entities.Entity) (*entities.Asset, error)
	GetUrl(asset *entities.Asset) string
	GetByEntity(entity entities.Entity, pageFilter *PageFilter) ([]*entities.Asset, error)
	Delete(asset *entities.Asset) error
}

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

func (s *AssetService) Delete(asset *entities.Asset) error {
	err := s.fileManager.Delete(asset.FileID, asset.Extension)
	if err != nil {
		return err
	}

	err = s.assetRepository.Delete(asset.ID)
	if err != nil {
		return err
	}

	return nil
}

type AssetServiceMock struct {
	mock.Mock
}

func (s *AssetServiceMock) CreateFromFile(file *os.File, entity entities.Entity) (*entities.Asset, error) {
	args := s.Called(file, entity)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Asset), args.Error(1)
}

func (s *AssetServiceMock) GetUrl(asset *entities.Asset) string {
	args := s.Called(asset)
	return args.String(0)
}

func (s *AssetServiceMock) GetByEntity(entity entities.Entity, pageFilter *PageFilter) ([]*entities.Asset, error) {
	args := s.Called(entity, pageFilter)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*entities.Asset), args.Error(1)
}

func (s *AssetServiceMock) Delete(asset *entities.Asset) error {
	args := s.Called(asset)
	return args.Error(0)
}
