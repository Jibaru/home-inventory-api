package gorm

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func makeDBMock() (*gorm.DB, sqlmock.Sqlmock) {
	dialector, dbMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dialector,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	return db, dbMock
}

func TestCreateUser(t *testing.T) {
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

func TestCreateUserError(t *testing.T) {
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
