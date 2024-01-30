package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	ErrBoxItemQuantityShouldBePositive = errors.New("box item quantity should be positive")
	ErrBoxItemBoxIDShouldNotBeEmpty    = errors.New("box item box id should not be empty")
)

type BoxItem struct {
	ID        string
	Quantity  float64
	BoxID     string
	ItemID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBoxItem(
	quantity float64,
	boxID string,
	item Item,
) (*BoxItem, error) {
	if quantity <= 0 {
		return nil, ErrBoxItemQuantityShouldBePositive
	}

	if strings.TrimSpace(boxID) == "" {
		return nil, ErrBoxItemBoxIDShouldNotBeEmpty
	}

	return &BoxItem{
		ID:        uuid.NewString(),
		Quantity:  quantity,
		BoxID:     boxID,
		ItemID:    item.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
