package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBoxItem(t *testing.T) {
	quantity := 100.0
	boxID := uuid.NewString()
	item := Item{
		ID: uuid.NewString(),
	}

	boxItem, err := NewBoxItem(quantity, boxID, item)

	assert.NoError(t, err)
	assert.NotEmpty(t, boxItem.ID)
	assert.Equal(t, quantity, boxItem.Quantity)
	assert.Equal(t, boxID, boxItem.BoxID)
	assert.Equal(t, item.ID, boxItem.ItemID)
	now := time.Now()
	assert.WithinDuration(t, now, boxItem.CreatedAt, 10*time.Second)
	assert.WithinDuration(t, now, boxItem.UpdatedAt, 10*time.Second)
}

func TestNewBoxItemErrorBoxItemQuantityShouldBePositive(t *testing.T) {
	quantity := -1.0
	boxID := uuid.NewString()
	item := Item{
		ID: uuid.NewString(),
	}

	boxItem, err := NewBoxItem(quantity, boxID, item)

	assert.Error(t, err)
	assert.Nil(t, boxItem)
	assert.ErrorIs(t, err, ErrBoxItemQuantityShouldBePositive)
}

func TestNewBoxItemErrorBoxItemBoxIDShouldNotBeEmpty(t *testing.T) {
	quantity := 100.0
	boxID := ""
	item := Item{
		ID: uuid.NewString(),
	}

	boxItem, err := NewBoxItem(quantity, boxID, item)

	assert.Error(t, err)
	assert.Nil(t, boxItem)
	assert.ErrorIs(t, err, ErrBoxItemBoxIDShouldNotBeEmpty)
}
