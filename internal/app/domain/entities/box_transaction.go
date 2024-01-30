package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	BoxTransactionTypeAdd    = "add"
	BoxTransactionTypeRemove = "remove"
)

var (
	ErrBoxTransactionQuantityShouldBePositive = errors.New("quantity should be positive")
	ErrBoxTransactionBoxIDShouldNotBeEmpty    = errors.New("box id should not be empty")
)

type BoxTransaction struct {
	ID         string
	Type       string
	Quantity   float64
	BoxID      string
	ItemID     string
	ItemSku    string
	ItemName   string
	ItemUnit   string
	HappenedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewAddBoxTransaction(
	quantity float64,
	boxID string,
	item Item,
	happenedAt time.Time,
) (*BoxTransaction, error) {
	return newBoxTransaction(
		BoxTransactionTypeAdd,
		quantity,
		boxID,
		item,
		happenedAt,
	)
}

func NewRemoveBoxTransaction(
	quantity float64,
	boxID string,
	item Item,
	happenedAt time.Time,
) (*BoxTransaction, error) {
	return newBoxTransaction(
		BoxTransactionTypeRemove,
		quantity,
		boxID,
		item,
		happenedAt,
	)
}

func newBoxTransaction(
	bType string,
	quantity float64,
	boxID string,
	item Item,
	happenedAt time.Time,
) (*BoxTransaction, error) {
	if quantity <= 0 {
		return nil, ErrBoxTransactionQuantityShouldBePositive
	}

	if strings.TrimSpace(boxID) == "" {
		return nil, ErrBoxTransactionBoxIDShouldNotBeEmpty
	}

	return &BoxTransaction{
		ID:         uuid.NewString(),
		Type:       bType,
		Quantity:   quantity,
		BoxID:      boxID,
		ItemID:     item.ID,
		ItemSku:    item.Sku,
		ItemName:   item.Name,
		ItemUnit:   item.Unit,
		HappenedAt: happenedAt,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
