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

func TestRoomRepositoryCreateRoom(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	description := random.String(255, random.Alphanumeric)
	room := &entities.Room{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `rooms` (`id`,`name`,`description`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(room.ID, room.Name, *room.Description, room.UserID, room.CreatedAt, room.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := roomRepository.Create(room)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryCreateRoomErrorCanNotCreateRoom(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	description := random.String(255, random.Alphanumeric)
	room := &entities.Room{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `rooms` (`id`,`name`,`description`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(room.ID, room.Name, *room.Description, room.UserID, room.CreatedAt, room.UpdatedAt).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := roomRepository.Create(room)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrCanNotCreateRoom)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
