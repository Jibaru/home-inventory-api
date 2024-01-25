package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	ErrItemKeywordItemIDShouldNotBeEmpty              = errors.New("item id should not be empty")
	ErrItemKeywordValueShouldNotBeEmpty               = errors.New("value should not be empty")
	ErrItemKeywordValueShouldHaveLessThan60Characters = errors.New("value should have less than 60 characters")
)

type ItemKeyword struct {
	ID        string
	Value     string
	ItemID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewItemKeyword(itemID, value string) (*ItemKeyword, error) {
	if strings.TrimSpace(itemID) == "" {
		return nil, ErrItemKeywordItemIDShouldNotBeEmpty
	}

	if strings.TrimSpace(value) == "" {
		return nil, ErrItemKeywordValueShouldNotBeEmpty
	}

	if len(value) > 60 {
		return nil, ErrItemKeywordValueShouldHaveLessThan60Characters
	}

	return &ItemKeyword{
		ID:        uuid.NewString(),
		Value:     value,
		ItemID:    itemID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
