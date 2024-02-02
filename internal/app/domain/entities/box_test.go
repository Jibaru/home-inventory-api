package entities

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBox(t *testing.T) {
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	roomID := uuid.NewString()

	box, err := NewBox(name, &description, roomID)

	assert.NoError(t, err)
	assert.NotEmpty(t, box.ID)
	assert.Equal(t, name, box.Name)
	assert.NotNil(t, box.Description)
	assert.Equal(t, description, *box.Description)

	now := time.Now()
	assert.WithinDuration(t, now, box.CreatedAt, 10*time.Second)
	assert.WithinDuration(t, now, box.UpdatedAt, 10*time.Second)
}

func TestNewBoxErrorNameShouldNotBeEmpty(t *testing.T) {
	name := ""
	description := random.String(255, random.Alphanumeric)
	roomID := uuid.NewString()

	box, err := NewBox(name, &description, roomID)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.ErrorIs(t, err, ErrBoxNameShouldNotBeEmpty)
}

func TestNewBoxErrorNameShouldHave100OrLessChars(t *testing.T) {
	name := random.String(101, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	roomID := uuid.NewString()

	box, err := NewBox(name, &description, roomID)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.ErrorIs(t, err, ErrBoxNameShouldHave100OrLessChars)
}

func TestNewBoxErrorDescriptionShouldNotBeEmpty(t *testing.T) {
	name := random.String(100, random.Alphanumeric)
	description := ""
	roomID := uuid.NewString()

	box, err := NewBox(name, &description, roomID)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.ErrorIs(t, err, ErrBoxDescriptionShouldNotBeEmpty)
}

func TestNewBoxErrorDescriptionShouldHave255OrLessChars(t *testing.T) {
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric) + random.String(1, random.Alphanumeric)
	roomID := uuid.NewString()

	box, err := NewBox(name, &description, roomID)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.ErrorIs(t, err, ErrBoxDescriptionShouldHave255OrLessChars)
}

func TestNewBoxErrorRoomIDShouldNotBeEmpty(t *testing.T) {
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	roomID := ""

	box, err := NewBox(name, &description, roomID)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.ErrorIs(t, err, ErrBoxRoomIDShouldNotBeEmpty)
}
