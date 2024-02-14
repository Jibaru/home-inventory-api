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
	item := &Item{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := item.Update(sku, name, description, unit)
	if err != nil {
		return nil, err
	}

	err = item.ChangeUserID(userID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *Item) Update(sku, name string, description *string, unit string) error {
	err := i.ChangeSku(sku)
	if err != nil {
		return err
	}

	err = i.ChangeName(name)
	if err != nil {
		return err
	}

	err = i.ChangeDescription(description)
	if err != nil {
		return err
	}

	err = i.ChangeUnit(unit)
	if err != nil {
		return err
	}

	i.UpdatedAt = time.Now()
	return nil
}

func (i *Item) EntityID() string {
	return i.ID
}

func (i *Item) EntityName() string {
	return "item"
}

func (i *Item) ChangeSku(sku string) error {
	if strings.TrimSpace(sku) == "" {
		return ErrItemSkuShouldNotBeEmpty
	}

	if len(sku) > 100 {
		return ErrItemSkuShouldHaveLessThan100Characters
	}

	i.Sku = sku
	return nil
}

func (i *Item) ChangeName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrItemNameShouldNotBeEmpty
	}

	if len(name) > 255 {
		return ErrItemNameShouldHaveLessThan255Characters
	}

	i.Name = name
	return nil
}

func (i *Item) ChangeDescription(description *string) error {
	if description != nil {
		if strings.TrimSpace(*description) == "" {
			return ErrItemDescriptionShouldNotBeEmpty
		}

		if len(*description) > 65535 {
			return ErrItemDescriptionShouldHaveLessThan65535Characters
		}
	}

	i.Description = description
	return nil
}

func (i *Item) ChangeUnit(unit string) error {
	if _, ok := ValidItemUnits[unit]; !ok {
		return ErrItemUnitShouldBeValid
	}

	i.Unit = unit
	return nil
}

func (i *Item) ChangeUserID(userID string) error {
	if strings.TrimSpace(userID) == "" {
		return ErrItemUserIDShouldNotBeEmpty
	}

	i.UserID = userID
	return nil
}
