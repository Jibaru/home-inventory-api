package entities

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewItem(t *testing.T) {
	sku := random.String(60, random.Alphanumeric)
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, sku, item.Sku)
	assert.Equal(t, name, item.Name)
	assert.Equal(t, description, *item.Description)
	assert.Equal(t, unit, item.Unit)
	assert.Equal(t, userID, item.UserID)

	now := time.Now()
	assert.WithinDuration(t, now, item.CreatedAt, 10*time.Second)
	assert.WithinDuration(t, now, item.UpdatedAt, 10*time.Second)
}

func TestNewItemErrorSkuShouldNotBeEmpty(t *testing.T) {
	sku := ""
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.ErrorIs(t, err, ErrItemSkuShouldNotBeEmpty)
}

func TestNewItemErrorSkuShouldHaveLessThan100Characters(t *testing.T) {
	sku := random.String(101, random.Alphanumeric)
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.ErrorIs(t, err, ErrItemSkuShouldHaveLessThan100Characters)
}

func TestNewItemErrorNameShouldNotBeEmpty(t *testing.T) {
	sku := random.String(60, random.Alphanumeric)
	name := ""
	description := random.String(255, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.ErrorIs(t, err, ErrItemNameShouldNotBeEmpty)
}

func TestNewItemErrorNameShouldHaveLessThan255Characters(t *testing.T) {
	sku := random.String(60, random.Alphanumeric)
	name := random.String(255, random.Alphanumeric) + random.String(1, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	unit := "unit"
	userID := uuid.NewString()

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.ErrorIs(t, err, ErrItemNameShouldHaveLessThan255Characters)
}

func TestNewItemErrorDescriptionShouldNotBeEmpty(t *testing.T) {
	sku := random.String(60, random.Alphanumeric)
	name := random.String(100, random.Alphanumeric)
	description := ""
	unit := "unit"
	userID := uuid.NewString()

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.ErrorIs(t, err, ErrItemDescriptionShouldNotBeEmpty)
}

func TestNewItemErrorDescriptionShouldHaveLessThan65535Characters(t *testing.T) {
	sku := random.String(60, random.Alphanumeric)
	name := random.String(100, random.Alphanumeric)
	description := ""
	for i := 0; i <= 65535; i++ {
		description += random.String(1, random.Alphanumeric)
	}
	unit := "unit"
	userID := uuid.NewString()

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.ErrorIs(t, err, ErrItemDescriptionShouldHaveLessThan65535Characters)
}

func TestNewItemErrorUserIDShouldNotBeEmpty(t *testing.T) {
	sku := random.String(60, random.Alphanumeric)
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	unit := "unit"
	userID := ""

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.ErrorIs(t, err, ErrItemUserIDShouldNotBeEmpty)
}

func TestNewItemErrorUnitShouldBeValid(t *testing.T) {
	sku := random.String(60, random.Alphanumeric)
	name := random.String(100, random.Alphanumeric)
	description := random.String(255, random.Alphanumeric)
	unit := "invalid-unit"
	userID := uuid.NewString()

	item, err := NewItem(
		sku,
		name,
		&description,
		unit,
		userID,
	)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.ErrorIs(t, err, ErrItemUnitShouldBeValid)
}

func TestItemEntityID(t *testing.T) {
	item := &Item{
		ID:          uuid.NewString(),
		Sku:         random.String(60, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: nil,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	assert.Equal(t, item.ID, item.EntityID())
}

func TestItemEntityName(t *testing.T) {
	item := &Item{
		ID:          uuid.NewString(),
		Sku:         random.String(60, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: nil,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	assert.Equal(t, "item", item.EntityName())
}
