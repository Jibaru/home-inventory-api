package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/stub"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	roomService := NewRoomService(roomRepository)

	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	userID := uuid.NewString()

	roomRepository.On("Create", mock.AnythingOfType("*entities.Room")).
		Return(nil)

	room, err := roomService.Create(name, &description, userID)

	assert.NoError(t, err)
	assert.NotNil(t, room)
	assert.Equal(t, name, room.Name)
	assert.Equal(t, description, *room.Description)
	assert.Equal(t, userID, room.UserID)
	roomRepository.AssertExpectations(t)
}

func TestCreateRoomError(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	roomService := NewRoomService(roomRepository)

	name := random.String(100, random.Alphanumeric)
	userID := uuid.NewString()

	mockError := errors.New("repository error")
	roomRepository.On("Create", mock.AnythingOfType("*entities.Room")).Return(mockError)

	room, err := roomService.Create(name, nil, userID)

	assert.Error(t, err)
	assert.Nil(t, room)
	assert.EqualError(t, err, mockError.Error())
	roomRepository.AssertExpectations(t)
}
