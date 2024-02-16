package gorm

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func TestUserRepositoryCreateUser(t *testing.T) {
	db, dbMock := makeDBMock()
	userRepository := NewUserRepository(db)

	user := &entities.User{
		ID:        uuid.NewString(),
		Email:     "test@example.com",
		Password:  "3ncr1pt3dP44sw0rd",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`email`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).
		WithArgs(user.ID, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := userRepository.Create(user)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepositoryCreateUserError(t *testing.T) {
	db, dbMock := makeDBMock()
	userRepository := NewUserRepository(db)

	user := &entities.User{
		ID:        uuid.NewString(),
		Email:     "test@example.com",
		Password:  "3ncr1pt3dP44sw0rd",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`email`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).
		WithArgs(user.ID, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).
		WillReturnError(errors.New("some error"))
	dbMock.ExpectRollback()

	err := userRepository.Create(user)

	assert.Error(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	db, dbMock := makeDBMock()
	userRepository := NewUserRepository(db)

	email := "test@email.com"
	expectedUser := entities.User{
		ID:        uuid.NewString(),
		Email:     email,
		Password:  "123abc",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
		AddRow(
			expectedUser.ID,
			expectedUser.Email,
			expectedUser.Password,
			expectedUser.CreatedAt,
			expectedUser.UpdatedAt,
		)

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(rows).
		WithArgs(email)

	user, err := userRepository.FindByEmail(email)

	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Password, user.Password)
	assert.Equal(t, expectedUser.CreatedAt, user.CreatedAt)
	assert.Equal(t, expectedUser.UpdatedAt, user.UpdatedAt)
	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepositoryFindByEmailErrorUserNotFound(t *testing.T) {
	db, dbMock := makeDBMock()
	userRepository := NewUserRepository(db)

	email := "test@email.com"

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1")).
		WillReturnError(gorm.ErrRecordNotFound).
		WithArgs(email)

	user, err := userRepository.FindByEmail(email)

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.ErrorIs(t, repositories.ErrUserRepositoryUserNotFound, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepositoryGetUserByBoxID(t *testing.T) {
	db, dbMock := makeDBMock()
	userRepository := NewUserRepository(db)

	boxID := uuid.NewString()

	expectedUser := entities.User{
		ID:        uuid.NewString(),
		Email:     "email@test.com",
		Password:  "encryptedPassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
		AddRow(
			expectedUser.ID,
			expectedUser.Email,
			expectedUser.Password,
			expectedUser.CreatedAt,
			expectedUser.UpdatedAt,
		)

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`id`,`users`.`email`,`users`.`password`,`users`.`created_at`,`users`.`updated_at` FROM `users` rooms on rooms.user_id = users.id boxes on boxes.room_id = rooms.id WHERE boxes.id = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs(boxID).
		WillReturnRows(rows)

	user, err := userRepository.GetUserByBoxID(boxID)

	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Password, user.Password)
	assert.Equal(t, expectedUser.CreatedAt, user.CreatedAt)
	assert.Equal(t, expectedUser.UpdatedAt, user.UpdatedAt)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepositoryGetUserByBoxIDErrorUserNotFound(t *testing.T) {
	db, dbMock := makeDBMock()
	userRepository := NewUserRepository(db)

	boxID := uuid.NewString()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`id`,`users`.`email`,`users`.`password`,`users`.`created_at`,`users`.`updated_at` FROM `users` rooms on rooms.user_id = users.id boxes on boxes.room_id = rooms.id WHERE boxes.id = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs(boxID).
		WillReturnError(gorm.ErrRecordNotFound)

	user, err := userRepository.GetUserByBoxID(boxID)

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.ErrorIs(t, repositories.ErrUserRepositoryUserNotFound, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
