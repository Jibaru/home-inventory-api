package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/stub"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
	"time"
)

func TestItemServiceCreate(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	itemRepository.On("Create", mock.AnythingOfType("*entities.Item")).
		Return(nil)
	itemKeywordRepository.On("CreateMany", mock.AnythingOfType("[]*entities.ItemKeyword")).
		Return(nil)
	assetService.On("CreateFromFile", mock.AnythingOfType("*os.File"), mock.AnythingOfType("*entities.Item")).
		Return(&entities.Asset{
			ID:         uuid.NewString(),
			Name:       random.String(10, random.Alphanumeric),
			Extension:  ".png",
			Size:       12314,
			FileID:     uuid.NewString(),
			EntityID:   uuid.NewString(),
			EntityName: "item",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, nil)

	sku := random.String(10, random.Alphanumeric)
	name := random.String(10, random.Alphanumeric)
	description := random.String(10, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()
	keywords := []string{
		random.String(10, random.Alphanumeric),
		random.String(10, random.Alphanumeric),
	}
	file := &os.File{}

	item, err := itemService.Create(
		sku,
		name,
		&description,
		unit,
		userID,
		keywords,
		file,
	)

	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, sku, item.Sku)
	assert.Equal(t, name, item.Name)
	assert.Equal(t, description, *item.Description)
	assert.Equal(t, unit, item.Unit)
	assert.Equal(t, userID, item.UserID)
	assert.NotEmpty(t, item.CreatedAt)
	assert.NotEmpty(t, item.UpdatedAt)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}

func TestItemServiceCreateErrorOnAssetService(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	assetService.On("CreateFromFile", mock.AnythingOfType("*os.File"), mock.AnythingOfType("*entities.Item")).
		Return(nil, errors.New("asset service error"))

	sku := random.String(10, random.Alphanumeric)
	name := random.String(10, random.Alphanumeric)
	description := random.String(10, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()
	keywords := []string{
		random.String(10, random.Alphanumeric),
		random.String(10, random.Alphanumeric),
	}
	file := &os.File{}

	item, err := itemService.Create(
		sku,
		name,
		&description,
		unit,
		userID,
		keywords,
		file,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}

func TestItemServiceCreateErrorOnItemRepository(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	assetService.On("CreateFromFile", mock.AnythingOfType("*os.File"), mock.AnythingOfType("*entities.Item")).
		Return(&entities.Asset{
			ID:         uuid.NewString(),
			Name:       random.String(10, random.Alphanumeric),
			Extension:  ".png",
			Size:       12314,
			FileID:     uuid.NewString(),
			EntityID:   uuid.NewString(),
			EntityName: "item",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, nil)
	assetService.On("Delete", mock.AnythingOfType("*entities.Asset")).
		Return(nil)
	itemRepository.On("Create", mock.AnythingOfType("*entities.Item")).
		Return(errors.New("item repository error"))

	sku := random.String(10, random.Alphanumeric)
	name := random.String(10, random.Alphanumeric)
	description := random.String(10, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()
	keywords := []string{
		random.String(10, random.Alphanumeric),
		random.String(10, random.Alphanumeric),
	}
	file := &os.File{}

	item, err := itemService.Create(
		sku,
		name,
		&description,
		unit,
		userID,
		keywords,
		file,
	)

	time.Sleep(2 * time.Second) // Needed to wait for the goroutine to finish

	assert.Error(t, err)
	assert.Nil(t, item)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}

func TestItemServiceCreateErrorOnItemKeywordRepository(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	assetService.On("CreateFromFile", mock.AnythingOfType("*os.File"), mock.AnythingOfType("*entities.Item")).
		Return(&entities.Asset{
			ID:         uuid.NewString(),
			Name:       random.String(10, random.Alphanumeric),
			Extension:  ".png",
			Size:       12314,
			FileID:     uuid.NewString(),
			EntityID:   uuid.NewString(),
			EntityName: "item",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, nil)
	assetService.On("Delete", mock.AnythingOfType("*entities.Asset")).
		Return(nil)
	itemRepository.On("Create", mock.AnythingOfType("*entities.Item")).
		Return(nil)
	itemKeywordRepository.On("CreateMany", mock.AnythingOfType("[]*entities.ItemKeyword")).
		Return(errors.New("item keyword repository error"))

	sku := random.String(10, random.Alphanumeric)
	name := random.String(10, random.Alphanumeric)
	description := random.String(10, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()
	keywords := []string{
		random.String(10, random.Alphanumeric),
		random.String(10, random.Alphanumeric),
	}
	file := &os.File{}

	item, err := itemService.Create(
		sku,
		name,
		&description,
		unit,
		userID,
		keywords,
		file,
	)

	time.Sleep(2 * time.Second) // Needed to wait for the goroutine to finish

	assert.Error(t, err)
	assert.Nil(t, item)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}

func TestItemServiceGetAll(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	itemRepository.On(
		"GetByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
		mock.AnythingOfType("*repositories.PageFilter"),
	).
		Return([]*entities.Item{
			{
				ID:          uuid.NewString(),
				Sku:         random.String(10, random.Alphanumeric),
				Name:        random.String(10, random.Alphanumeric),
				Description: nil,
				Unit:        "unit",
				UserID:      uuid.NewString(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}, nil)
	assetService.On("GetByEntities", mock.AnythingOfType("[]entities.Entity")).
		Return([]*entities.Asset{
			{
				ID:         uuid.NewString(),
				Name:       random.String(10, random.Alphanumeric),
				Extension:  ".png",
				Size:       12314,
				FileID:     uuid.NewString(),
				EntityID:   uuid.NewString(),
				EntityName: "item",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
		}, nil)

	items, err := itemService.GetAll("search", uuid.NewString(), PageFilter{
		Page: 1,
		Size: 1,
	})

	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.NotEmpty(t, items)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}

func TestItemServiceGetAllErrorOnItemRepository(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	itemRepository.On(
		"GetByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
		mock.AnythingOfType("*repositories.PageFilter"),
	).
		Return(nil, errors.New("item repository error"))

	items, err := itemService.GetAll("search", uuid.NewString(), PageFilter{
		Page: 1,
		Size: 1,
	})

	assert.Error(t, err)
	assert.Nil(t, items)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}

func TestItemServiceGetAllErrorOnAssetService(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	itemRepository.On(
		"GetByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
		mock.AnythingOfType("*repositories.PageFilter"),
	).
		Return([]*entities.Item{
			{
				ID:          uuid.NewString(),
				Sku:         random.String(10, random.Alphanumeric),
				Name:        random.String(10, random.Alphanumeric),
				Description: nil,
				Unit:        "unit",
				UserID:      uuid.NewString(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}, nil)
	assetService.On("GetByEntities", mock.AnythingOfType("[]entities.Entity")).
		Return(nil, errors.New("asset service error"))

	items, err := itemService.GetAll("search", uuid.NewString(), PageFilter{
		Page: 1,
		Size: 1,
	})

	assert.Error(t, err)
	assert.Nil(t, items)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}

func TestItemServiceCountAll(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	itemRepository.On(
		"CountByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
	).
		Return(int64(1), nil)

	count, err := itemService.CountAll("search", uuid.NewString())

	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}

func TestItemServiceCountAllErrorOnItemRepository(t *testing.T) {
	itemRepository := &stub.ItemRepositoryMock{}
	itemKeywordRepository := &stub.ItemKeywordRepositoryMock{}
	assetService := &AssetServiceMock{}

	itemService := NewItemService(
		itemRepository,
		itemKeywordRepository,
		assetService,
	)

	itemRepository.On(
		"CountByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
	).
		Return(int64(0), errors.New("item repository error"))

	count, err := itemService.CountAll("search", uuid.NewString())

	assert.Error(t, err)
	assert.Equal(t, int64(0), count)
	itemRepository.AssertExpectations(t)
	itemKeywordRepository.AssertExpectations(t)
	assetService.AssertExpectations(t)
}
