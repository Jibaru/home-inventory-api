package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/stub"
	serviceStubs "github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/services/stub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"path/filepath"
	"testing"
)

type MockEntity struct {
	ID string
}

func (e *MockEntity) EntityID() string {
	return e.ID
}

func (e *MockEntity) EntityName() string {
	return "mock_entity"
}

func TestCreateFromFile(t *testing.T) {
	assetRepository := &stub.AssetRepositoryMock{}
	fileManager := &serviceStubs.FileManagerMock{}
	service := NewAssetService(fileManager, assetRepository)

	tempFile, err := os.CreateTemp("", "*_"+uuid.NewString())
	defer tempFile.Close()
	assert.NoError(t, err)

	name := filepath.Base(tempFile.Name())
	extension := filepath.Ext(tempFile.Name())
	info, err := tempFile.Stat()
	assert.NoError(t, err)

	entity := &MockEntity{uuid.NewString()}
	fileID := uuid.NewString()

	assetRepository.On("Create", mock.AnythingOfType("*entities.Asset")).
		Return(nil)
	fileManager.On("Upload", mock.AnythingOfType("*os.File")).
		Return(fileID, nil)

	asset, err := service.CreateFromFile(tempFile, entity)

	assert.NoError(t, err)
	assert.NotNil(t, asset)
	assert.NotEmpty(t, asset.ID)
	assert.Equal(t, name, asset.Name)
	assert.Equal(t, extension, asset.Extension)
	assert.Equal(t, info.Size(), asset.Size)
	assert.Equal(t, fileID, asset.FileID)
	assert.Equal(t, entity.EntityID(), asset.EntityID)
	assert.Equal(t, entity.EntityName(), asset.EntityName)
	assert.NotEmpty(t, asset.CreatedAt)
	assert.NotEmpty(t, asset.UpdatedAt)
	assetRepository.AssertExpectations(t)
	fileManager.AssertExpectations(t)
}

func TestCreateFromFileErrorFromFileManager(t *testing.T) {
	assetRepository := &stub.AssetRepositoryMock{}
	fileManager := &serviceStubs.FileManagerMock{}
	service := NewAssetService(fileManager, assetRepository)

	tempFile, err := os.CreateTemp("", "*_"+uuid.NewString())
	defer tempFile.Close()
	assert.NoError(t, err)

	entity := &MockEntity{uuid.NewString()}

	fileManager.On("Upload", mock.AnythingOfType("*os.File")).
		Return("", errors.New("file manager error"))

	asset, err := service.CreateFromFile(tempFile, entity)

	assert.Error(t, err)
	assert.Nil(t, asset)
	assetRepository.AssertExpectations(t)
	fileManager.AssertExpectations(t)
}

func TestCreateFromFileErrorFromFileAssetRepository(t *testing.T) {
	assetRepository := &stub.AssetRepositoryMock{}
	fileManager := &serviceStubs.FileManagerMock{}
	service := NewAssetService(fileManager, assetRepository)

	tempFile, err := os.CreateTemp("", "*_"+uuid.NewString())
	defer tempFile.Close()
	assert.NoError(t, err)

	entity := &MockEntity{uuid.NewString()}

	assetRepository.On("Create", mock.AnythingOfType("*entities.Asset")).
		Return(errors.New("repository error"))
	fileManager.On("Upload", mock.AnythingOfType("*os.File")).
		Return(uuid.NewString(), nil)

	asset, err := service.CreateFromFile(tempFile, entity)

	assert.Error(t, err)
	assert.Nil(t, asset)
	assetRepository.AssertExpectations(t)
	fileManager.AssertExpectations(t)
}

func TestGetUrl(t *testing.T) {
	assetRepository := &stub.AssetRepositoryMock{}
	fileManager := &serviceStubs.FileManagerMock{}
	service := NewAssetService(fileManager, assetRepository)

	asset := &entities.Asset{
		Extension: ".png",
		FileID:    uuid.NewString(),
	}

	fileManager.On("GenerateUrl", asset.FileID, asset.Extension).
		Return("valid-url")

	url := service.GetUrl(asset)

	assert.NotEmpty(t, url)
	assetRepository.AssertExpectations(t)
	fileManager.AssertExpectations(t)
}

func TestGetByEntityWithoutPageFilter(t *testing.T) {
	assetRepository := &stub.AssetRepositoryMock{}
	fileManager := &serviceStubs.FileManagerMock{}
	service := NewAssetService(fileManager, assetRepository)

	var expectedAssets []*entities.Asset
	var repositoryPageFilter *repositories.PageFilter

	entity := &MockEntity{uuid.NewString()}

	assetRepository.On("FindByEntity", entity, repositoryPageFilter).
		Return(expectedAssets, nil)

	assets, err := service.GetByEntity(entity, nil)

	assert.NoError(t, err)
	assert.Len(t, assets, len(expectedAssets))
	assetRepository.AssertExpectations(t)
}

func TestGetByEntityWithPageFilter(t *testing.T) {
	assetRepository := &stub.AssetRepositoryMock{}
	fileManager := &serviceStubs.FileManagerMock{}
	service := NewAssetService(fileManager, assetRepository)

	var expectedAssets []*entities.Asset
	pageFilter := &PageFilter{
		Page: 1,
		Size: 10,
	}
	repositoryPageFilter := &repositories.PageFilter{
		Offset: (pageFilter.Page - 1) * pageFilter.Size,
		Limit:  pageFilter.Size,
	}

	entity := &MockEntity{uuid.NewString()}

	assetRepository.On("FindByEntity", entity, repositoryPageFilter).
		Return(expectedAssets, nil)

	assets, err := service.GetByEntity(entity, pageFilter)

	assert.NoError(t, err)
	assert.Len(t, assets, len(expectedAssets))
	assetRepository.AssertExpectations(t)
}

func TestGetByEntityErrorFromAssetRepository(t *testing.T) {
	assetRepository := &stub.AssetRepositoryMock{}
	fileManager := &serviceStubs.FileManagerMock{}
	service := NewAssetService(fileManager, assetRepository)

	pageFilter := &PageFilter{
		Page: 1,
		Size: 10,
	}
	repositoryPageFilter := &repositories.PageFilter{
		Offset: (pageFilter.Page - 1) * pageFilter.Size,
		Limit:  pageFilter.Size,
	}

	entity := &MockEntity{uuid.NewString()}

	assetRepository.On("FindByEntity", entity, repositoryPageFilter).
		Return(nil, errors.New("repository error"))

	assets, err := service.GetByEntity(entity, pageFilter)

	assert.Error(t, err)
	assert.Nil(t, assets)
	assetRepository.AssertExpectations(t)
}