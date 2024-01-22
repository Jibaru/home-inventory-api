package gorm

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strconv"
	"testing"
	"time"
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

func TestCreateAsset(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	asset := &entities.Asset{
		ID:         uuid.NewString(),
		Name:       random.String(100, random.Alphanumeric),
		Extension:  ".jpg",
		Size:       89813,
		FileID:     uuid.NewString(),
		EntityID:   uuid.NewString(),
		EntityName: "user",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `assets` (`id`,`name`,`extension`,`size`,`file_id`,`entity_id`,`entity_name`,`created_at`,`updated_at`) VALUES  (?,?,?,?,?,?,?,?,?)")).
		WithArgs(
			asset.ID,
			asset.Name,
			asset.Extension,
			asset.Size,
			asset.FileID,
			asset.EntityID,
			asset.EntityName,
			asset.CreatedAt,
			asset.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := assetRepository.Create(asset)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateAssetErrorCanNotCreateAsset(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	asset := &entities.Asset{
		ID:         uuid.NewString(),
		Name:       random.String(100, random.Alphanumeric),
		Extension:  ".jpg",
		Size:       89813,
		FileID:     uuid.NewString(),
		EntityID:   uuid.NewString(),
		EntityName: "user",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `assets` (`id`,`name`,`extension`,`size`,`file_id`,`entity_id`,`entity_name`,`created_at`,`updated_at`) VALUES  (?,?,?,?,?,?,?,?,?)")).
		WithArgs(
			asset.ID,
			asset.Name,
			asset.Extension,
			asset.Size,
			asset.FileID,
			asset.EntityID,
			asset.EntityName,
			asset.CreatedAt,
			asset.UpdatedAt).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := assetRepository.Create(asset)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrCanNotCreateAsset)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindByEntity(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	entity := &MockEntity{
		ID: uuid.NewString(),
	}

	expectedAsset := &entities.Asset{
		ID:         uuid.NewString(),
		Name:       random.String(100, random.Alphanumeric),
		Extension:  ".jpg",
		Size:       89813,
		FileID:     uuid.NewString(),
		EntityID:   entity.EntityID(),
		EntityName: entity.EntityName(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "extension", "size", "file_id", "entity_id", "entity_name", "created_at", "updated_at"}).
		AddRow(
			expectedAsset.ID,
			expectedAsset.Name,
			expectedAsset.Extension,
			expectedAsset.Size,
			expectedAsset.FileID,
			expectedAsset.EntityID,
			expectedAsset.EntityName,
			expectedAsset.CreatedAt,
			expectedAsset.UpdatedAt,
		)

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `assets` WHERE entity_id = ? AND entity_name = ?")).
		WillReturnRows(rows).
		WithArgs(entity.EntityID(), entity.EntityName())

	assets, err := assetRepository.FindByEntity(entity, nil)

	assert.NotNil(t, assets)
	assert.Len(t, assets, 1)

	asset := assets[0]

	assert.Equal(t, expectedAsset.ID, asset.ID)
	assert.Equal(t, expectedAsset.Name, asset.Name)
	assert.Equal(t, expectedAsset.Extension, asset.Extension)
	assert.Equal(t, expectedAsset.Size, asset.Size)
	assert.Equal(t, expectedAsset.FileID, asset.FileID)
	assert.Equal(t, expectedAsset.EntityID, asset.EntityID)
	assert.Equal(t, expectedAsset.EntityName, asset.EntityName)
	assert.Equal(t, expectedAsset.CreatedAt, asset.CreatedAt)
	assert.Equal(t, expectedAsset.UpdatedAt, asset.UpdatedAt)
	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindByEntityWithPageFilter(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	entity := &MockEntity{
		ID: uuid.NewString(),
	}

	pageFilter := &repositories.PageFilter{
		Offset: 1,
		Limit:  1,
	}

	expectedAsset := &entities.Asset{
		ID:         uuid.NewString(),
		Name:       random.String(100, random.Alphanumeric),
		Extension:  ".jpg",
		Size:       89813,
		FileID:     uuid.NewString(),
		EntityID:   entity.EntityID(),
		EntityName: entity.EntityName(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "extension", "size", "file_id", "entity_id", "entity_name", "created_at", "updated_at"}).
		AddRow(
			expectedAsset.ID,
			expectedAsset.Name,
			expectedAsset.Extension,
			expectedAsset.Size,
			expectedAsset.FileID,
			expectedAsset.EntityID,
			expectedAsset.EntityName,
			expectedAsset.CreatedAt,
			expectedAsset.UpdatedAt,
		)

	dbMock.ExpectQuery(
		regexp.QuoteMeta(
			"SELECT * FROM `assets` WHERE entity_id = ? AND entity_name = ? LIMIT "+
				strconv.Itoa(pageFilter.Limit)+
				" OFFSET "+strconv.Itoa(pageFilter.Offset),
		),
	).
		WillReturnRows(rows).
		WithArgs(entity.EntityID(), entity.EntityName())

	assets, err := assetRepository.FindByEntity(entity, pageFilter)

	assert.NotNil(t, assets)
	assert.Len(t, assets, 1)

	asset := assets[0]

	assert.Equal(t, expectedAsset.ID, asset.ID)
	assert.Equal(t, expectedAsset.Name, asset.Name)
	assert.Equal(t, expectedAsset.Extension, asset.Extension)
	assert.Equal(t, expectedAsset.Size, asset.Size)
	assert.Equal(t, expectedAsset.FileID, asset.FileID)
	assert.Equal(t, expectedAsset.EntityID, asset.EntityID)
	assert.Equal(t, expectedAsset.EntityName, asset.EntityName)
	assert.Equal(t, expectedAsset.CreatedAt, asset.CreatedAt)
	assert.Equal(t, expectedAsset.UpdatedAt, asset.UpdatedAt)
	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindByEntityErrorCanNotGetAssets(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	entity := &MockEntity{
		ID: uuid.NewString(),
	}

	pageFilter := &repositories.PageFilter{
		Offset: 1,
		Limit:  1,
	}

	dbMock.ExpectQuery(
		regexp.QuoteMeta(
			"SELECT * FROM `assets` WHERE entity_id = ? AND entity_name = ? LIMIT "+
				strconv.Itoa(pageFilter.Limit)+
				" OFFSET "+strconv.Itoa(pageFilter.Offset),
		),
	).
		WillReturnError(errors.New("database error")).
		WithArgs(entity.EntityID(), entity.EntityName())

	assets, err := assetRepository.FindByEntity(entity, pageFilter)

	assert.Nil(t, assets)
	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrCanNotGetAssets)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
