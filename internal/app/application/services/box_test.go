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

func TestBoxServiceRemoveItemFromBoxDeleteBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	itemID := uuid.NewString()
	quantity := 10.0

	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(&entities.BoxItem{
			BoxID:    boxID,
			ItemID:   itemID,
			Quantity: 10.0,
		}, nil)
	boxRepository.On("DeleteBoxItem", boxID, itemID).
		Return(nil)
	boxRepository.On("CreateBoxTransaction", mock.AnythingOfType("*entities.BoxTransaction")).
		Return(nil)

	err := boxService.RemoveItemFromBox(quantity, boxID, itemID)

	time.Sleep(2 * time.Second)

	assert.NoError(t, err)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceRemoveItemFromBoxUpdateBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	itemID := uuid.NewString()
	quantity := 5.0

	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(&entities.BoxItem{
			BoxID:    boxID,
			ItemID:   itemID,
			Quantity: 10.0,
		}, nil)
	boxRepository.On("UpdateBoxItem", mock.AnythingOfType("*entities.BoxItem")).
		Return(nil)
	boxRepository.On("CreateBoxTransaction", mock.AnythingOfType("*entities.BoxTransaction")).
		Return(nil)

	err := boxService.RemoveItemFromBox(quantity, boxID, itemID)

	time.Sleep(2 * time.Second)

	assert.NoError(t, err)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceRemoveItemFromBoxErrorInItemRepository(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	itemID := uuid.NewString()
	quantity := 5.0

	mockError := errors.New("repository error")
	itemRepository.On("GetByID", itemID).
		Return(nil, mockError)

	err := boxService.RemoveItemFromBox(quantity, boxID, itemID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceRemoveItemFromBoxErrorInBoxRepositoryOnGetBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	itemID := uuid.NewString()
	quantity := 5.0

	mockError := errors.New("repository error")
	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(nil, mockError)

	err := boxService.RemoveItemFromBox(quantity, boxID, itemID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceRemoveItemFromBoxErrorInBoxRepositoryOnDeleteBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	itemID := uuid.NewString()
	quantity := 10.0

	mockError := errors.New("repository error")
	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(&entities.BoxItem{
			BoxID:    boxID,
			ItemID:   itemID,
			Quantity: 10.0,
		}, nil)
	boxRepository.On("DeleteBoxItem", boxID, itemID).
		Return(mockError)

	err := boxService.RemoveItemFromBox(quantity, boxID, itemID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceRemoveItemFromBoxErrorInBoxRepositoryOnUpdateBoxItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	itemID := uuid.NewString()
	quantity := 5.0

	mockError := errors.New("repository error")
	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", boxID, itemID).
		Return(&entities.BoxItem{
			BoxID:    boxID,
			ItemID:   itemID,
			Quantity: 10.0,
		}, nil)
	boxRepository.On("UpdateBoxItem", mock.AnythingOfType("*entities.BoxItem")).
		Return(mockError)

	err := boxService.RemoveItemFromBox(quantity, boxID, itemID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceGetAll(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	roomID := uuid.NewString()
	userID := uuid.NewString()
	search := "search"

	pageFilter := PageFilter{
		Page: 1,
		Size: 10,
	}

	boxRepository.On(
		"GetByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
		mock.AnythingOfType("*repositories.PageFilter"),
	).
		Return([]*entities.Box{
			{
				ID:          uuid.NewString(),
				Name:        "box",
				Description: nil,
				RoomID:      roomID,
			},
		}, nil)

	boxes, err := boxService.GetAll(roomID, userID, search, pageFilter)

	assert.NoError(t, err)
	assert.NotNil(t, boxes)
	assert.Len(t, boxes, 1)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceGetAllErrorInBoxRepository(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	roomID := uuid.NewString()
	userID := uuid.NewString()
	search := "search"

	pageFilter := PageFilter{
		Page: 1,
		Size: 10,
	}

	mockError := errors.New("repository error")
	boxRepository.On(
		"GetByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
		mock.AnythingOfType("*repositories.PageFilter"),
	).
		Return(nil, mockError)

	boxes, err := boxService.GetAll(roomID, userID, search, pageFilter)

	assert.Error(t, err)
	assert.Nil(t, boxes)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceCountAll(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	roomID := uuid.NewString()
	userID := uuid.NewString()
	search := "search"
	expectedCount := 10

	boxRepository.On(
		"CountByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
	).
		Return(int64(expectedCount), nil)

	count, err := boxService.CountAll(userID, search, roomID)

	assert.NoError(t, err)
	assert.Equal(t, int64(expectedCount), count)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceCountAllErrorInBoxRepository(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	roomID := uuid.NewString()
	userID := uuid.NewString()
	search := "search"

	mockError := errors.New("repository error")
	boxRepository.On(
		"CountByQueryFilters",
		mock.AnythingOfType("repositories.QueryFilter"),
	).
		Return(int64(0), mockError)

	count, err := boxService.CountAll(userID, search, roomID)

	assert.Error(t, err)
	assert.Equal(t, int64(0), count)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceTransferItem(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	originBoxID := uuid.NewString()
	destinationBoxID := uuid.NewString()
	itemID := uuid.NewString()

	itemRepository.On("GetByID", itemID).
		Return(&entities.Item{
			ID: itemID,
		}, nil)
	boxRepository.On("GetBoxItem", originBoxID, itemID).
		Return(&entities.BoxItem{
			BoxID:    originBoxID,
			ItemID:   itemID,
			Quantity: 10.0,
		}, nil)
	boxRepository.On("GetBoxItem", destinationBoxID, itemID).
		Return(&entities.BoxItem{
			BoxID:    originBoxID,
			ItemID:   itemID,
			Quantity: 10.0,
		}, nil)
	boxRepository.On("DeleteBoxItem", originBoxID, itemID).
		Return(nil)
	boxRepository.On("CreateBoxTransaction", mock.AnythingOfType("*entities.BoxTransaction")).
		Return(nil)
	boxRepository.On("UpdateBoxItem", mock.AnythingOfType("*entities.BoxItem")).
		Return(nil)

	err := boxService.TransferItem(originBoxID, destinationBoxID, itemID)

	time.Sleep(2 * time.Second)

	assert.NoError(t, err)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceDeleteWithTransactionsAndItemQuantities(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()

	boxRepository.On("DeleteBoxTransactionsByBoxID", boxID).
		Return(nil)
	boxRepository.On("DeleteBoxItemsByBoxID", boxID).
		Return(nil)
	boxRepository.On("Delete", boxID).
		Return(nil)

	err := boxService.DeleteWithTransactionsAndItemQuantities(boxID)

	assert.NoError(t, err)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceDeleteWithTransactionsAndItemQuantitiesErrorInBoxRepositoryOnDeleteBoxTransactionsByBoxID(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()

	mockError := errors.New("repository error")
	boxRepository.On("DeleteBoxTransactionsByBoxID", boxID).
		Return(mockError)

	err := boxService.DeleteWithTransactionsAndItemQuantities(boxID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceDeleteWithTransactionsAndItemQuantitiesErrorInBoxRepositoryOnDeleteBoxItemsByBoxID(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()

	mockError := errors.New("repository error")
	boxRepository.On("DeleteBoxTransactionsByBoxID", boxID).
		Return(nil)
	boxRepository.On("DeleteBoxItemsByBoxID", boxID).
		Return(mockError)

	err := boxService.DeleteWithTransactionsAndItemQuantities(boxID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceDeleteWithTransactionsAndItemQuantitiesErrorInBoxRepositoryOnDeleteBox(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()

	mockError := errors.New("repository error")
	boxRepository.On("DeleteBoxItemsByBoxID", boxID).
		Return(nil)
	boxRepository.On("DeleteBoxTransactionsByBoxID", boxID).
		Return(nil)
	boxRepository.On("Delete", boxID).
		Return(mockError)

	err := boxService.DeleteWithTransactionsAndItemQuantities(boxID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceUpdate(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	name := "box"
	description := "description"

	boxRepository.On("GetByID", boxID).
		Return(&entities.Box{
			ID:          boxID,
			Name:        name,
			Description: &description,
		}, nil)
	boxRepository.On("Update", mock.AnythingOfType("*entities.Box")).
		Return(nil)

	box, err := boxService.Update(boxID, name, &description)

	assert.NoError(t, err)
	assert.NotNil(t, box)
	assert.Equal(t, boxID, box.ID)
	assert.Equal(t, name, box.Name)
	assert.Equal(t, description, *box.Description)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceUpdateErrorInBoxRepositoryOnGetByID(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	name := "box"
	description := "description"

	mockError := errors.New("repository error")
	boxRepository.On("GetByID", boxID).
		Return(nil, mockError)

	box, err := boxService.Update(boxID, name, &description)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceUpdateErrorInBoxRepositoryOnUpdate(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	name := "box"
	description := "description"

	mockError := errors.New("repository error")
	boxRepository.On("GetByID", boxID).
		Return(&entities.Box{
			ID:          boxID,
			Name:        name,
			Description: &description,
		}, nil)
	boxRepository.On("Update", mock.AnythingOfType("*entities.Box")).
		Return(mockError)

	box, err := boxService.Update(boxID, name, &description)

	assert.Error(t, err)
	assert.Nil(t, box)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceTransferToRoom(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	roomID := uuid.NewString()

	boxRepository.On("GetByID", boxID).
		Return(&entities.Box{
			ID: boxID,
		}, nil)
	boxRepository.On("Update", mock.AnythingOfType("*entities.Box")).
		Return(nil)

	err := boxService.TransferToRoom(boxID, roomID)

	assert.NoError(t, err)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceTransferToRoomErrorInBoxRepositoryOnGetByID(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	roomID := uuid.NewString()

	mockError := errors.New("repository error")
	boxRepository.On("GetByID", boxID).
		Return(nil, mockError)

	err := boxService.TransferToRoom(boxID, roomID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceTransferToRoomErrorInBoxRepositoryOnUpdate(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()
	roomID := uuid.NewString()

	mockError := errors.New("repository error")
	boxRepository.On("GetByID", boxID).
		Return(&entities.Box{}, nil)
	boxRepository.On("Update", mock.AnythingOfType("*entities.Box")).
		Return(mockError)

	err := boxService.TransferToRoom(boxID, roomID)

	assert.Error(t, err)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceGetBoxTransactions(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()

	pageFilter := PageFilter{
		Page: 1,
		Size: 1,
	}

	boxRepository.On("GetBoxTransactionsByQueryFilters", mock.AnythingOfType("repositories.QueryFilter"), mock.AnythingOfType("*repositories.PageFilter")).
		Return([]*entities.BoxTransaction{}, nil)

	transactions, err := boxService.GetBoxTransactions(boxID, pageFilter)

	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceGetBoxTransactionsErrorInBoxRepositoryOnGetBoxTransactionsByQueryFilters(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()

	pageFilter := PageFilter{
		Page: 1,
		Size: 1,
	}

	mockError := errors.New("repository error")
	boxRepository.On("GetBoxTransactionsByQueryFilters", mock.AnythingOfType("repositories.QueryFilter"), mock.AnythingOfType("*repositories.PageFilter")).
		Return(nil, mockError)

	transactions, err := boxService.GetBoxTransactions(boxID, pageFilter)

	assert.Error(t, err)
	assert.Nil(t, transactions)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceCountBoxTransactions(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()

	boxRepository.On("CountBoxTransactionsByQueryFilters", mock.AnythingOfType("repositories.QueryFilter")).
		Return(int64(1), nil)

	count, err := boxService.CountBoxTransactions(boxID)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}

func TestBoxServiceCountBoxTransactionsErrorInBoxRepositoryOnCountBoxTransactionsByQueryFilters(t *testing.T) {
	boxRepository := new(stub.BoxRepositoryMock)
	roomRepository := new(stub.RoomRepositoryMock)
	itemRepository := new(stub.ItemRepositoryMock)

	boxService := NewBoxService(boxRepository, itemRepository, roomRepository)

	boxID := uuid.NewString()

	mockError := errors.New("repository error")
	boxRepository.On("CountBoxTransactionsByQueryFilters", mock.AnythingOfType("repositories.QueryFilter")).
		Return(int64(0), mockError)

	count, err := boxService.CountBoxTransactions(boxID)

	assert.Error(t, err)
	assert.Equal(t, int64(0), count)
	assert.EqualError(t, err, mockError.Error())
	boxRepository.AssertExpectations(t)
	itemRepository.AssertExpectations(t)
	roomRepository.AssertExpectations(t)
}
