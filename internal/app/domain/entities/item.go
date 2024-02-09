package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	ErrItemSkuShouldNotBeEmpty                          = errors.New("sku should not be empty")
	ErrItemSkuShouldHaveLessThan100Characters           = errors.New("sku should have less than 100 characters")
	ErrItemNameShouldNotBeEmpty                         = errors.New("name should not be empty")
	ErrItemNameShouldHaveLessThan255Characters          = errors.New("name should have less than 255 characters")
	ErrItemDescriptionShouldNotBeEmpty                  = errors.New("description should not be empty")
	ErrItemDescriptionShouldHaveLessThan65535Characters = errors.New("description should have less than 65535 characters")
	ErrItemUnitShouldBeValid                            = errors.New("unit should be valid")
	ErrItemUserIDShouldNotBeEmpty                       = errors.New("user id should not be empty")

	ValidItemUnits = map[string]bool{
		"kg":    true,
		"l":     true,
		"m":     true,
		"m2":    true,
		"m3":    true,
		"unit":  true,
		"cm":    true,
		"inch":  true,
		"ft":    true,
		"g":     true,
		"oz":    true,
		"piece": true,
		"dozen": true,
		"yd":    true,
		"gal":   true,
		"pt":    true,
		"qt":    true,
	}
)

type Item struct {
	ID          string
	Sku         string
	Name        string
	Description *string
	Unit        string
	UserID      string
	Keywords    []*ItemKeyword
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewItem(sku, name string, description *string, unit string, userID string) (*Item, error) {
	if strings.TrimSpace(sku) == "" {
		return nil, ErrItemSkuShouldNotBeEmpty
	}

	if len(sku) > 100 {
		return nil, ErrItemSkuShouldHaveLessThan100Characters
	}

	if strings.TrimSpace(name) == "" {
		return nil, ErrItemNameShouldNotBeEmpty
	}

	if len(name) > 255 {
		return nil, ErrItemNameShouldHaveLessThan255Characters
	}

	if description != nil {
		if strings.TrimSpace(*description) == "" {
			return nil, ErrItemDescriptionShouldNotBeEmpty
		}

		if len(*description) > 65535 {
			return nil, ErrItemDescriptionShouldHaveLessThan65535Characters
		}
	}

	if strings.TrimSpace(userID) == "" {
		return nil, ErrItemUserIDShouldNotBeEmpty
	}

	if _, ok := ValidItemUnits[unit]; !ok {
		return nil, ErrItemUnitShouldBeValid
	}

	return &Item{
		ID:          uuid.NewString(),
		Sku:         sku,
		Name:        name,
		Description: description,
		Unit:        unit,
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (i *Item) EntityID() string {
	return i.ID
}

func (i *Item) EntityName() string {
	return "item"
}
