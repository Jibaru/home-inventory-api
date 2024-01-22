package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/stub"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestUserServiceCreateUser(t *testing.T) {
	mockRepo := new(stub.UserRepositoryMock)
	userService := NewUserService(mockRepo)

	email := "test@example.com"
	password := random.String(5, random.Numeric) + random.String(5, random.Alphabetic)
	mockRepo.On("Create", mock.Anything).Return(nil)

	user, err := userService.CreateUser(email, password)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Create", mock.Anything)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, email, user.Email)

	now := time.Now()
	assert.WithinDuration(t, now, user.CreatedAt, 10*time.Second)
	assert.WithinDuration(t, now, user.UpdatedAt, 10*time.Second)
}

func TestUserServiceCreateUserErrorInRepository(t *testing.T) {
	mockRepo := new(stub.UserRepositoryMock)
	userService := NewUserService(mockRepo)

	email := "test@example.com"
	password := random.String(5, random.Numeric) + random.String(5, random.Alphabetic)

	expectedError := errors.New("repository error")
	mockRepo.On("Create", mock.Anything).Return(expectedError)

	user, err := userService.CreateUser(email, password)

	assert.Error(t, err)
	assert.EqualError(t, err, "cannot create user")
	assert.Nil(t, user)
	mockRepo.AssertCalled(t, "Create", mock.Anything)
}
