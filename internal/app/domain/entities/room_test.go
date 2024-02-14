package entities

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRoom(t *testing.T) {
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	userID := uuid.NewString()

	room, err := NewRoom(name, &description, userID)

	assert.NoError(t, err)
	assert.NotEmpty(t, room.ID)
	assert.Equal(t, name, room.Name)
	assert.NotNil(t, room.Description)
	assert.Equal(t, description, *room.Description)

	now := time.Now()
	assert.WithinDuration(t, now, room.CreatedAt, 10*time.Second)
	assert.WithinDuration(t, now, room.UpdatedAt, 10*time.Second)
}

func TestNewRoomErrorNameShouldNotBeEmpty(t *testing.T) {
	name := ""
	description := random.String(255, random.Alphanumeric)
	userID := uuid.NewString()

	room, err := NewRoom(name, &description, userID)

	assert.Error(t, err)
	assert.Nil(t, room)
	assert.ErrorIs(t, err, ErrRoomNameShouldNotBeEmpty)
}

func TestNewRoomErrorNameShouldHave100OrLessChars(t *testing.T) {
	name := random.String(101, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	userID := uuid.NewString()

	room, err := NewRoom(name, &description, userID)

	assert.Error(t, err)
	assert.Nil(t, room)
	assert.ErrorIs(t, err, ErrRoomNameShouldHave100OrLessChars)
}

func TestNewRoomErrorDescriptionShouldNotBeEmpty(t *testing.T) {
	name := random.String(100, random.Alphanumeric)
	description := ""
	userID := uuid.NewString()

	room, err := NewRoom(name, &description, userID)

	assert.Error(t, err)
	assert.Nil(t, room)
	assert.ErrorIs(t, err, ErrRoomDescriptionShouldNotBeEmpty)
}

func TestNewRoomErrorDescriptionShouldHave255OrLessChars(t *testing.T) {
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric) + random.String(1, random.Alphanumeric)
	userID := uuid.NewString()

	room, err := NewRoom(name, &description, userID)

	assert.Error(t, err)
	assert.Nil(t, room)
	assert.ErrorIs(t, err, ErrRoomDescriptionShouldHave255OrLessChars)
}

func TestNewRoomErrorUserIDShouldNotBeEmpty(t *testing.T) {
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	userID := ""

	room, err := NewRoom(name, &description, userID)

	assert.Error(t, err)
	assert.Nil(t, room)
	assert.ErrorIs(t, err, ErrRoomUserIDShouldNotBeEmpty)
}

func TestRoomUpdate(t *testing.T) {
	description := random.String(100, random.Alphanumeric)
	room := &Room{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
	}

	name := random.String(100, random.Alphanumeric)
	description = random.String(100, random.Alphanumeric)

	err := room.Update(name, &description)

	assert.NoError(t, err)
	assert.Equal(t, name, room.Name)
	assert.Equal(t, description, *room.Description)
}

func TestRoomUpdateErrorNameShouldNotBeEmpty(t *testing.T) {
	description := random.String(100, random.Alphanumeric)
	room := &Room{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
	}

	name := ""
	description = random.String(100, random.Alphanumeric)

	err := room.Update(name, &description)

	assert.Error(t, err)
	assert.Equal(t, room.Name, room.Name)
	assert.Equal(t, *room.Description, *room.Description)
	assert.ErrorIs(t, err, ErrRoomNameShouldNotBeEmpty)
}

func TestRoomUpdateErrorNameShouldHave100OrLessChars(t *testing.T) {
	description := random.String(100, random.Alphanumeric)
	room := &Room{
		ID:          uuid.NewString(),
		Name:        random.String(20, random.Alphanumeric),
		Description: &description,
	}

	name := random.String(101, random.Alphanumeric)
	description = random.String(100, random.Alphanumeric)

	err := room.Update(name, &description)

	assert.Error(t, err)
	assert.Equal(t, room.Name, room.Name)
	assert.Equal(t, *room.Description, *room.Description)
	assert.ErrorIs(t, err, ErrRoomNameShouldHave100OrLessChars)
}

func TestRoomUpdateErrorDescriptionShouldHave255OrLessChars(t *testing.T) {
	description := random.String(100, random.Alphanumeric)
	room := &Room{
		ID:          uuid.NewString(),
		Name:        random.String(40, random.Alphanumeric),
		Description: &description,
	}

	name := random.String(100, random.Alphanumeric)
	description = random.String(255, random.Alphanumeric) + "1"

	err := room.Update(name, &description)

	assert.Error(t, err)
	assert.Equal(t, room.Name, room.Name)
	assert.Equal(t, *room.Description, *room.Description)
	assert.ErrorIs(t, err, ErrRoomDescriptionShouldHave255OrLessChars)
}
