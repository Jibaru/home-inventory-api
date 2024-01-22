package entities

import (
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	email := "test@example.com"
	password := random.String(6, random.Numeric) + random.String(6, random.Alphabetic)

	user, err := NewUser(email, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, email, user.Email)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	assert.NoError(t, err)

	now := time.Now()
	assert.WithinDuration(t, now, user.CreatedAt, 10*time.Second)
	assert.WithinDuration(t, now, user.UpdatedAt, 10*time.Second)
}

func TestNewUserErrorInvalidEmail(t *testing.T) {
	invalidEmail := "invalidemail"
	password := random.String(6, random.Numeric) + random.String(6, random.Alphabetic)

	_, err := NewUser(invalidEmail, password)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidEmailAddress)
}

func TestNewUserErrorEmailExceedsMaxLength(t *testing.T) {
	invalidEmail := random.String(101) + "@email.com"
	password := random.String(6, random.Numeric) + random.String(6, random.Alphabetic)

	_, err := NewUser(invalidEmail, password)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrEmailExceedsMaxLengthOf100Chars)
}

func TestNewUserErrorPasswordMustBeBetween6And100Chars(t *testing.T) {
	smallPassword := random.String(5)
	giantPassword := random.String(101)

	_, err1 := NewUser("test@example.com", smallPassword)
	_, err2 := NewUser("test@example.com", giantPassword)

	assert.Error(t, err1)
	assert.ErrorIs(t, err1, ErrPasswordMustBeBetween6And100Chars)
	assert.Error(t, err2)
	assert.ErrorIs(t, err1, ErrPasswordMustBeBetween6And100Chars)
}

func TestNewUserErrorPasswordMustContainAtLeastOneNumber(t *testing.T) {
	invalidPassword := random.String(6, random.Alphabetic)

	_, err := NewUser("test@example.com", invalidPassword)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrPasswordMustContainAtLeastOneLetterAndOneNumber)
}

func TestNewUserErrorPasswordMustContainAtLeastOneLetter(t *testing.T) {
	invalidPassword := random.String(6, random.Numeric)

	_, err := NewUser("test@example.com", invalidPassword)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrPasswordMustContainAtLeastOneLetterAndOneNumber)
}
