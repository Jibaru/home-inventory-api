package services

import (
	"errors"
	tokenstub "github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/auth/stub"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/stub"
	"testing"

	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	userRepositoryMock := new(stub.UserRepositoryMock)
	tokenGeneratorMock := new(tokenstub.TokenGeneratorMock)

	authService := NewAuthService(userRepositoryMock, tokenGeneratorMock)

	email := "test@example.com"
	password := "123abc"
	user := &entities.User{
		ID:       "test_user_id",
		Email:    email,
		Password: "$2a$14$9VTo1/y3dUttmnaRERp41etwpGvk4Atv8UkKWqwqU20dHlzYu/rDa",
	}

	userRepositoryMock.On("FindByEmail", email).
		Return(user, nil)
	tokenGeneratorMock.On("GenerateToken", user.ID, user.Email).
		Return("fake_token", nil)

	result, err := authService.Authenticate(email, password)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user, result.User)
	assert.Equal(t, "fake_token", result.Token)
	userRepositoryMock.AssertExpectations(t)
	tokenGeneratorMock.AssertExpectations(t)
}

func TestAuthenticateInvalidCredentials(t *testing.T) {
	userRepositoryMock := new(stub.UserRepositoryMock)
	tokenGeneratorMock := new(tokenstub.TokenGeneratorMock)

	authService := NewAuthService(userRepositoryMock, tokenGeneratorMock)

	email := "test@example.com"
	password := "InvalidPassword"
	user := &entities.User{
		ID:       "test_user_id",
		Email:    email,
		Password: "$2a$14$9VTo1/y3dUttmnaRERp41etwpGvk4Atv8UkKWqwqU20dHlzYu/rDa",
	}

	userRepositoryMock.On("FindByEmail", email).Return(user, nil)

	result, err := authService.Authenticate(email, password)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid credentials")
	userRepositoryMock.AssertExpectations(t)
	tokenGeneratorMock.AssertNotCalled(t, "GenerateToken")
}

func TestAuthenticateErrorInRepository(t *testing.T) {
	userRepositoryMock := new(stub.UserRepositoryMock)
	tokenGeneratorMock := new(tokenstub.TokenGeneratorMock)

	authService := NewAuthService(userRepositoryMock, tokenGeneratorMock)

	email := "test@example.com"
	password := "TestPassword123"

	userRepositoryMock.On("FindByEmail", email).
		Return(nil, errors.New("repository error"))

	result, err := authService.Authenticate(email, password)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")

	userRepositoryMock.AssertExpectations(t)
	tokenGeneratorMock.AssertNotCalled(t, "GenerateToken")
}

func TestAuthenticateErrorInTokenGenerator(t *testing.T) {
	userRepositoryMock := new(stub.UserRepositoryMock)
	tokenGeneratorMock := new(tokenstub.TokenGeneratorMock)

	authService := NewAuthService(userRepositoryMock, tokenGeneratorMock)

	email := "test@example.com"
	password := "123abc"
	user := &entities.User{
		ID:       "test_user_id",
		Email:    email,
		Password: "$2a$14$9VTo1/y3dUttmnaRERp41etwpGvk4Atv8UkKWqwqU20dHlzYu/rDa",
	}

	userRepositoryMock.On("FindByEmail", email).Return(user, nil)
	tokenGeneratorMock.On("GenerateToken", user.ID, user.Email).
		Return("", errors.New("token generation error"))

	result, err := authService.Authenticate(email, password)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "can not authenticate")
	userRepositoryMock.AssertExpectations(t)
	tokenGeneratorMock.AssertExpectations(t)
}