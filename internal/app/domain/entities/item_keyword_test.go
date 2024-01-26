package entities

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewItemKeyword(t *testing.T) {
	value := random.String(20, random.Alphanumeric)
	itemID := uuid.NewString()

	itemKeyword, err := NewItemKeyword(itemID, value)

	assert.NoError(t, err)
	assert.NotNil(t, itemKeyword)
	assert.Equal(t, itemID, itemKeyword.ItemID)
	assert.Equal(t, value, itemKeyword.Value)
}

func TestNewItemKeywordErrorItemKeywordItemIDShouldNotBeEmpty(t *testing.T) {
	value := random.String(100, random.Alphanumeric)
	itemID := ""

	itemKeyword, err := NewItemKeyword(itemID, value)

	assert.Error(t, err)
	assert.Nil(t, itemKeyword)
	assert.ErrorIs(t, err, ErrItemKeywordItemIDShouldNotBeEmpty)
}

func TestNewItemKeywordErrorItemKeywordValueShouldNotBeEmpty(t *testing.T) {
	value := ""
	itemID := uuid.NewString()

	itemKeyword, err := NewItemKeyword(itemID, value)

	assert.Error(t, err)
	assert.Nil(t, itemKeyword)
	assert.ErrorIs(t, err, ErrItemKeywordValueShouldNotBeEmpty)
}

func TestNewItemKeywordErrorItemKeywordValueShouldHaveLessThan60Characters(t *testing.T) {
	value := random.String(61, random.Alphanumeric)
	itemID := uuid.NewString()

	itemKeyword, err := NewItemKeyword(itemID, value)

	assert.Error(t, err)
	assert.Nil(t, itemKeyword)
	assert.ErrorIs(t, err, ErrItemKeywordValueShouldHaveLessThan60Characters)
}
