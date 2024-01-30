package gorm

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestItemRepositoryCreate(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	description := random.String(255, random.Alphanumeric)
	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `items` (`id`,`sku`,`name`,`description`,`unit`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(item.ID, item.Sku, item.Name, *item.Description, item.Unit, item.UserID, item.CreatedAt, item.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := itemRepository.Create(item)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryCreateErrorCanNotCreateItem(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	description := random.String(255, random.Alphanumeric)
	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `items` (`id`,`sku`,`name`,`description`,`unit`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(item.ID, item.Sku, item.Name, *item.Description, item.Unit, item.UserID, item.CreatedAt, item.UpdatedAt).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := itemRepository.Create(item)

	assert.Error(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryGetByID(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	description := random.String(255, random.Alphanumeric)
	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id", "sku", "name", "description", "unit", "user_id", "created_at", "updated_at"}).
		AddRow(item.ID, item.Sku, item.Name, item.Description, item.Unit, item.UserID, item.CreatedAt, item.UpdatedAt)
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `items` WHERE id = ? ORDER BY `items`.`id` LIMIT 1")).
		WithArgs(item.ID).
		WillReturnRows(rows)

	itemResult, err := itemRepository.GetByID(item.ID)

	assert.NoError(t, err)
	assert.Equal(t, item.ID, itemResult.ID)
	assert.Equal(t, item.Sku, itemResult.Sku)
	assert.Equal(t, item.Name, itemResult.Name)
	assert.Equal(t, item.Description, itemResult.Description)
	assert.Equal(t, item.Unit, itemResult.Unit)
	assert.Equal(t, item.UserID, itemResult.UserID)
	assert.Equal(t, item.CreatedAt, itemResult.CreatedAt)
	assert.Equal(t, item.UpdatedAt, itemResult.UpdatedAt)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryGetByIDErrorCanNotGetItem(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	itemID := uuid.NewString()
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `items` WHERE id = ? ORDER BY `items`.`id` LIMIT 1")).
		WithArgs(itemID).
		WillReturnError(errors.New("database error"))

	itemResult, err := itemRepository.GetByID(itemID)

	assert.Error(t, err)
	assert.Nil(t, itemResult)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
