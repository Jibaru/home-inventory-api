package entities

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewAddBoxTransaction(t *testing.T) {
	quantity := 100.0
	boxID := uuid.NewString()
	itemDescription := random.String(100, random.Alphanumeric)
	item := Item{
		ID:          uuid.NewString(),
		Sku:         random.String(4, random.Alphanumeric),
		Name:        random.String(10, random.Alphanumeric),
		Description: &itemDescription,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	happenedAt := time.Now()

	boxItem, err := NewAddBoxTransaction(quantity, boxID, item, happenedAt)

	assert.NoError(t, err)
	assert.NotEmpty(t, boxItem.ID)
	assert.Equal(t, quantity, boxItem.Quantity)
	assert.Equal(t, boxID, boxItem.BoxID)
	assert.Equal(t, item.ID, boxItem.ItemID)
	assert.Equal(t, item.Sku, boxItem.ItemSku)
	assert.Equal(t, item.Name, boxItem.ItemName)
	assert.Equal(t, item.Unit, boxItem.ItemUnit)
	assert.Equal(t, happenedAt, boxItem.HappenedAt)
	now := time.Now()
	assert.WithinDuration(t, now, boxItem.CreatedAt, 10*time.Second)
	assert.WithinDuration(t, now, boxItem.UpdatedAt, 10*time.Second)
}

func TestNewAddBoxTransactionErrorBoxTransactionQuantityShouldBePositive(t *testing.T) {
	quantity := -1.0
	boxID := uuid.NewString()
	itemDescription := random.String(100, random.Alphanumeric)
	item := Item{
		ID:          uuid.NewString(),
		Sku:         random.String(4, random.Alphanumeric),
		Name:        random.String(10, random.Alphanumeric),
		Description: &itemDescription,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	happenedAt := time.Now()

	boxItem, err := NewAddBoxTransaction(quantity, boxID, item, happenedAt)

	assert.Error(t, err)
	assert.Nil(t, boxItem)
	assert.ErrorIs(t, err, ErrBoxTransactionQuantityShouldBePositive)
}

func TestNewAddBoxTransactionErrorBoxTransactionBoxIDShouldNotBeEmpty(t *testing.T) {
	quantity := 100.0
	boxID := ""
	itemDescription := random.String(100, random.Alphanumeric)
	item := Item{
		ID:          uuid.NewString(),
		Sku:         random.String(4, random.Alphanumeric),
		Name:        random.String(10, random.Alphanumeric),
		Description: &itemDescription,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	happenedAt := time.Now()

	boxItem, err := NewAddBoxTransaction(quantity, boxID, item, happenedAt)

	assert.Error(t, err)
	assert.Nil(t, boxItem)
	assert.ErrorIs(t, err, ErrBoxTransactionBoxIDShouldNotBeEmpty)
}
