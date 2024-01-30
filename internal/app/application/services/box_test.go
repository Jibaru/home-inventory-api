package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/stub"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestBoxServiceCreateBox(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

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
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceCreateBoxErrorInRoomRepository(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

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
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceCreateBoxErrorInBoxRepository(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

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
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceAddItemIntoBoxWhenThereIsNoBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	quantity := 1.0
	boxID := uuid.NewString()
	itemID := uuid.NewString()

	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(nil, repositories.ErrBoxRepositoryBoxItemNotFound)
	boxRepository.On("CreateBoxItem", mock.AnythingOfType("*entities.BoxItem")).
		Return(nil)
	boxRepository.On("CreateBoxTransaction", mock.AnythingOfType("*entities.BoxTransaction")).
		Return(nil)

	boxItem, err := boxService.AddItemIntoBox(quantity, boxID, itemID)

	time.Sleep(2 * time.Second)

	assert.NoError(t, err)
	assert.NotNil(t, boxItem)
	assert.Equal(t, quantity, boxItem.Quantity)
	assert.Equal(t, boxID, boxItem.BoxID)
	assert.Equal(t, itemID, boxItem.ItemID)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceAddItemIntoBoxWhenThereIsBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	quantity := 1.0
	boxID := uuid.NewString()
	itemID := uuid.NewString()

	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(&entities.BoxItem{
			BoxID:    boxID,
			ItemID:   itemID,
			Quantity: 1.0,
		}, nil)
	boxRepository.On("UpdateBoxItem", mock.AnythingOfType("*entities.BoxItem")).
		Return(nil)
	boxRepository.On("CreateBoxTransaction", mock.AnythingOfType("*entities.BoxTransaction")).
		Return(nil)

	boxItem, err := boxService.AddItemIntoBox(quantity, boxID, itemID)

	time.Sleep(2 * time.Second)

	assert.NoError(t, err)
	assert.NotNil(t, boxItem)
	assert.Equal(t, quantity+1.0, boxItem.Quantity)
	assert.Equal(t, boxID, boxItem.BoxID)
	assert.Equal(t, itemID, boxItem.ItemID)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceAddItemIntoBoxErrorInItemRepository(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	quantity := 1.0
	boxID := uuid.NewString()
	itemID := uuid.NewString()

	mockError := errors.New("repository error")
	itemRepository.On("GetByID", itemID).
		Return(nil, mockError)

	boxItem, err := boxService.AddItemIntoBox(quantity, boxID, itemID)

	assert.Error(t, err)
	assert.Nil(t, boxItem)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceAddItemIntoBoxErrorInBoxRepositoryOnCreateBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	quantity := 1.0
	boxID := uuid.NewString()
	itemID := uuid.NewString()

	mockError := errors.New("repository error")
	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(nil, repositories.ErrBoxRepositoryBoxItemNotFound)
	boxRepository.On("CreateBoxItem", mock.AnythingOfType("*entities.BoxItem")).
		Return(mockError)

	boxItem, err := boxService.AddItemIntoBox(quantity, boxID, itemID)

	assert.Error(t, err)
	assert.Nil(t, boxItem)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceAddItemIntoBoxErrorInBoxRepositoryOnUpdateBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	quantity := 1.0
	boxID := uuid.NewString()
	itemID := uuid.NewString()

	mockError := errors.New("repository error")
	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(&entities.BoxItem{
			BoxID:    boxID,
			ItemID:   itemID,
			Quantity: 1.0,
		}, nil)
	boxRepository.On("UpdateBoxItem", mock.AnythingOfType("*entities.BoxItem")).
		Return(mockError)

	boxItem, err := boxService.AddItemIntoBox(quantity, boxID, itemID)

	assert.Error(t, err)
	assert.Nil(t, boxItem)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceAddItemIntoBoxErrorInBoxRepositoryOnGetBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)
	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	quantity := 1.0
	boxID := uuid.NewString()
	itemID := uuid.NewString()

	mockError := errors.New("repository error")
	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(nil, mockError)

	boxItem, err := boxService.AddItemIntoBox(quantity, boxID, itemID)

	assert.Error(t, err)
	assert.Nil(t, boxItem)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}
