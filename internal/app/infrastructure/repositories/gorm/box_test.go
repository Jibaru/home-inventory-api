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

func TestBoxRepositoryCreateBox(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	description := random.String(255, random.Alphanumeric)
	box := &entities.Box{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		RoomID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `boxes` (`id`,`name`,`description`,`room_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(box.ID, box.Name, *box.Description, box.RoomID, box.CreatedAt, box.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.Create(box)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCreateBoxErrorCanNotCreateBox(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	description := random.String(255, random.Alphanumeric)
	box := &entities.Box{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		RoomID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `boxes` (`id`,`name`,`description`,`room_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(box.ID, box.Name, *box.Description, box.RoomID, box.CreatedAt, box.UpdatedAt).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.Create(box)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrCanNotCreateBox)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
