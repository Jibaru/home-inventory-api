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

func TestAssetRepositoryCreateAsset(t *testing.T) {
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

func TestAssetRepositoryCreateAssetErrorCanNotCreateAsset(t *testing.T) {
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
	assert.ErrorIs(t, err, repositories.ErrAssetRepositoryCanNotCreateAsset)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAssetRepositoryFindByEntity(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	entity := entities.NewIdentifiableEntity(uuid.NewString())

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

func TestAssetRepositoryFindByEntityWithPageFilter(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	entity := entities.NewIdentifiableEntity(uuid.NewString())

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

func TestAssetRepositoryFindByEntityErrorCanNotGetAssets(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	entity := entities.NewIdentifiableEntity(uuid.NewString())

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
	assert.ErrorIs(t, err, repositories.ErrAssetRepositoryCanNotGetAssets)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAssetRepositoryDelete(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	id := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `assets` WHERE id = ?")).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := assetRepository.Delete(id)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAssetRepositoryGetByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	ids := []string{uuid.NewString(), uuid.NewString()}

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "entity_id",
						Operator: repositories.InComparisonOperator,
						Value:    ids,
					},
					{
						Field:    "entity_name",
						Operator: repositories.EqualComparisonOperator,
						Value:    "user",
					},
				},
			},
		},
	}

	expectedAssets := []*entities.Asset{
		{
			ID:         uuid.NewString(),
			Name:       random.String(100, random.Alphanumeric),
			Extension:  ".jpg",
			Size:       89813,
			FileID:     uuid.NewString(),
			EntityID:   ids[0],
			EntityName: "user",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         uuid.NewString(),
			Name:       random.String(100, random.Alphanumeric),
			Extension:  ".jpg",
			Size:       89813,
			FileID:     uuid.NewString(),
			EntityID:   ids[1],
			EntityName: "user",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "extension", "size", "file_id", "entity_id", "entity_name", "created_at", "updated_at"}).
		AddRow(
			expectedAssets[0].ID,
			expectedAssets[0].Name,
			expectedAssets[0].Extension,
			expectedAssets[0].Size,
			expectedAssets[0].FileID,
			expectedAssets[0].EntityID,
			expectedAssets[0].EntityName,
			expectedAssets[0].CreatedAt,
			expectedAssets[0].UpdatedAt,
		).
		AddRow(
			expectedAssets[1].ID,
			expectedAssets[1].Name,
			expectedAssets[1].Extension,
			expectedAssets[1].Size,
			expectedAssets[1].FileID,
			expectedAssets[1].EntityID,
			expectedAssets[1].EntityName,
			expectedAssets[1].CreatedAt,
			expectedAssets[1].UpdatedAt,
		)

	dbMock.ExpectQuery(
		regexp.QuoteMeta(
			"SELECT * FROM `assets` WHERE entity_id IN (?,?) AND entity_name = ?",
		),
	).
		WillReturnRows(rows).
		WithArgs(ids[0], ids[1], "user")

	assets, err := assetRepository.GetByQueryFilters(queryFilter)

	assert.NotNil(t, assets)
	assert.Len(t, assets, 2)
	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAssetRepositoryGetByQueryFiltersErrorCanNotGetByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	assetRepository := NewAssetRepository(db)

	ids := []string{uuid.NewString(), uuid.NewString()}

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "entity_id",
						Operator: repositories.InComparisonOperator,
						Value:    ids,
					},
					{
						Field:    "entity_name",
						Operator: repositories.EqualComparisonOperator,
						Value:    "user",
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(
		regexp.QuoteMeta(
			"SELECT * FROM `assets` WHERE entity_id IN (?,?) AND entity_name = ?",
		),
	).
		WillReturnError(errors.New("database error")).
		WithArgs(ids[0], ids[1], "user")

	assets, err := assetRepository.GetByQueryFilters(queryFilter)

	assert.Nil(t, assets)
	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrorAssetRepositoryCanNotGetByQueryFilters)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
