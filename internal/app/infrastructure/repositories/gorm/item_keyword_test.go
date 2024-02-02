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
	"testing"
	"time"
)

func TestItemKeywordRepositoryCreateMany(t *testing.T) {
	db, dbMock := makeDBMock()
	itemKeywordRepository := NewItemKeywordRepository(db)

	itemKeywords := []*entities.ItemKeyword{
		{
			ID:        uuid.NewString(),
			ItemID:    uuid.NewString(),
			Value:     random.String(10, random.Alphanumeric),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.NewString(),
			ItemID:    uuid.NewString(),
			Value:     random.String(10, random.Alphanumeric),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `item_keywords` (`id`,`value`,`item_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?),(?,?,?,?,?)")).
		WithArgs(
			itemKeywords[0].ID,
			itemKeywords[0].Value,
			itemKeywords[0].ItemID,
			itemKeywords[0].CreatedAt,
			itemKeywords[0].UpdatedAt,
			itemKeywords[1].ID,
			itemKeywords[1].Value,
			itemKeywords[1].ItemID,
			itemKeywords[1].CreatedAt,
			itemKeywords[1].UpdatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := itemKeywordRepository.CreateMany(itemKeywords)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemKeywordRepositoryCreateManyErrorCanNotCreateItemKeywords(t *testing.T) {
	db, dbMock := makeDBMock()
	itemKeywordRepository := NewItemKeywordRepository(db)

	itemKeywords := []*entities.ItemKeyword{
		{
			ID:        uuid.NewString(),
			ItemID:    uuid.NewString(),
			Value:     random.String(10, random.Alphanumeric),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.NewString(),
			ItemID:    uuid.NewString(),
			Value:     random.String(10, random.Alphanumeric),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `item_keywords` (`id`,`value`,`item_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?),(?,?,?,?,?)")).
		WithArgs(
			itemKeywords[0].ID,
			itemKeywords[0].Value,
			itemKeywords[0].ItemID,
			itemKeywords[0].CreatedAt,
			itemKeywords[0].UpdatedAt,
			itemKeywords[1].ID,
			itemKeywords[1].Value,
			itemKeywords[1].ItemID,
			itemKeywords[1].CreatedAt,
			itemKeywords[1].UpdatedAt,
		).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := itemKeywordRepository.CreateMany(itemKeywords)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrItemKeywordRepositoryCanNotCreateItemKeywords)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
