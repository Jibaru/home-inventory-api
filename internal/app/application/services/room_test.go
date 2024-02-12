package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/stub"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRoomServiceCreateRoom(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

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
	boxRepository.AssertExpectations(t)
}

func TestRoomServiceCreateRoomErrorInRepository(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

	name := random.String(100, random.Alphanumeric)
	userID := uuid.NewString()

	mockError := errors.New("repository error")
	roomRepository.On("Create", mock.AnythingOfType("*entities.Room")).Return(mockError)

	room, err := roomService.Create(name, nil, userID)

	assert.Error(t, err)
	assert.Nil(t, room)
	assert.EqualError(t, err, mockError.Error())
	roomRepository.AssertExpectations(t)
	boxRepository.AssertExpectations(t)
}

func TestRoomServiceGetAll(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

	search := random.String(100, random.Alphanumeric)
	userID := uuid.NewString()
	pageFilter := PageFilter{
		Page: 1,
		Size: 10,
	}

	rooms := []*entities.Room{
		{
			ID:          uuid.NewString(),
			Name:        random.String(100, random.Alphanumeric),
			Description: nil,
			UserID:      userID,
		},
	}
	roomRepository.On("GetByQueryFilters", mock.AnythingOfType("repositories.QueryFilter"), mock.AnythingOfType("*repositories.PageFilter")).
		Return(rooms, nil)

	result, err := roomService.GetAll(search, userID, pageFilter)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, rooms, result)
	roomRepository.AssertExpectations(t)
	boxRepository.AssertExpectations(t)
}

func TestRoomServiceGetAllErrorInRepository(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

	search := random.String(100, random.Alphanumeric)
	userID := uuid.NewString()
	pageFilter := PageFilter{
		Page: 1,
		Size: 10,
	}

	mockError := errors.New("repository error")
	roomRepository.On(
		"GetByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
		mock.AnythingOfType("*repositories.PageFilter"),
	).Return(nil, mockError)

	result, err := roomService.GetAll(search, userID, pageFilter)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, mockError.Error())
	roomRepository.AssertExpectations(t)
	boxRepository.AssertExpectations(t)
}

func TestRoomServiceCountAll(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

	search := random.String(100, random.Alphanumeric)
	userID := uuid.NewString()

	count := int64(10)
	roomRepository.On("CountByQueryFilters", mock.AnythingOfType("repositories.QueryFilter")).Return(count, nil)

	result, err := roomService.CountAll(search, userID)

	assert.NoError(t, err)
	assert.Equal(t, count, result)
	roomRepository.AssertExpectations(t)
	boxRepository.AssertExpectations(t)
}

func TestRoomServiceCountAllErrorInRepository(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

	search := random.String(100, random.Alphanumeric)
	userID := uuid.NewString()

	mockError := errors.New("repository error")
	roomRepository.On("CountByQueryFilters", mock.AnythingOfType("repositories.QueryFilter")).Return(int64(0), mockError)

	result, err := roomService.CountAll(search, userID)

	assert.Error(t, err)
	assert.Equal(t, int64(0), result)
	assert.EqualError(t, err, mockError.Error())
	roomRepository.AssertExpectations(t)
	boxRepository.AssertExpectations(t)
}

func TestRoomServiceDelete(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

	roomID := uuid.NewString()

	boxRepository.On("CountByQueryFilters", mock.AnythingOfType("repositories.QueryFilter")).
		Return(int64(0), nil)
	roomRepository.On("Delete", roomID).
		Return(nil)

	err := roomService.Delete(roomID)

	assert.NoError(t, err)
	roomRepository.AssertExpectations(t)
	boxRepository.AssertExpectations(t)
}

func TestRoomServiceDeleteErrorInRepository(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

	roomID := uuid.NewString()

	mockError := errors.New("repository error")
	boxRepository.On("CountByQueryFilters", mock.AnythingOfType("repositories.QueryFilter")).
		Return(int64(0), nil)
	roomRepository.On("Delete", roomID).Return(mockError)

	err := roomService.Delete(roomID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	roomRepository.AssertExpectations(t)
	boxRepository.AssertExpectations(t)
}

func TestRoomServiceDeleteErrorCanNotDeleteRoomWithBoxes(t *testing.T) {
	roomRepository := new(stub.RoomRepositoryMock)
	boxRepository := new(stub.BoxRepositoryMock)
	roomService := NewRoomService(roomRepository, boxRepository)

	roomID := uuid.NewString()

	boxRepository.On("CountByQueryFilters", mock.AnythingOfType("repositories.QueryFilter")).
		Return(int64(1), nil)

	err := roomService.Delete(roomID)

	assert.Error(t, err)
	assert.EqualError(t, err, "can not delete room with boxes")
	roomRepository.AssertExpectations(t)
	boxRepository.AssertExpectations(t)
}
