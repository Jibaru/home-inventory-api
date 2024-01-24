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

func TestBoxServiceCreateBox(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	boxService := NewBoxService(boxRepository, roomRepository)

	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	roomID := uuid.NewString()

	roomRepository.On("ExistsByID", roomID).
		Return(true, nil)
	boxRepository.On("Create", mock.AnythingOfType("*entities.Box")).
		Return(nil)

	box, err := boxService.Create(name, &description, roomID)

	assert.NoError(t, err)
	assert.NotNil(t, box)
	assert.Equal(t, name, box.Name)
	assert.Equal(t, description, *box.Description)
	assert.Equal(t, roomID, box.RoomID)
	boxRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceCreateBoxErrorInRoomRepository(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	boxService := NewBoxService(boxRepository, roomRepository)

	name := random.String(100, random.Alphanumeric)
	roomID := uuid.NewString()

	mockError := errors.New("repository error")
	roomRepository.On("ExistsByID", roomID).
		Return(false, mockError)

	box, err := boxService.Create(name, nil, roomID)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceCreateBoxErrorInBoxRepository(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	boxService := NewBoxService(boxRepository, roomRepository)

	name := random.String(100, random.Alphanumeric)
	roomID := uuid.NewString()

	mockError := errors.New("repository error")
	roomRepository.On("ExistsByID", roomID).
		Return(true, nil)
	boxRepository.On("Create", mock.AnythingOfType("*entities.Box")).
		Return(mockError)

	box, err := boxService.Create(name, nil, roomID)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}
