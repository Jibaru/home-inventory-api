package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	ErrBoxDescriptionShouldHave255OrLessChars = errors.New("description should have 255 or less chars")
	ErrBoxDescriptionShouldNotBeEmpty         = errors.New("description should not be empty")
	ErrBoxNameShouldHave100OrLessChars        = errors.New("name should have 100 or less chars")
	ErrBoxNameShouldNotBeEmpty                = errors.New("name should not be empty")
	ErrBoxRoomIDShouldNotBeEmpty              = errors.New("room id should not be empty")
)

type Box struct {
	ID          string
	Name        string
	Description *string
	RoomID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewBox(
	name string,
	description *string,
	roomID string,
) (*Box, error) {
	box := &Box{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := box.Update(name, description)
	if err != nil {
		return nil, err
	}

	err = box.ChangeRoomID(roomID)
	if err != nil {
		return nil, err
	}

	return box, nil
}

func (b *Box) ChangeName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrBoxNameShouldNotBeEmpty
	}

	if len(name) > 100 {
		return ErrBoxNameShouldHave100OrLessChars
	}

	b.Name = name
	return nil
}

func (b *Box) ChangeDescription(description *string) error {
	if description != nil {
		if strings.TrimSpace(*description) == "" {
			return ErrBoxDescriptionShouldNotBeEmpty
		}

		if len(*description) > 255 {
			return ErrBoxDescriptionShouldHave255OrLessChars
		}
	}

	b.Description = description
	return nil
}

func (b *Box) ChangeRoomID(roomID string) error {
	if strings.TrimSpace(roomID) == "" {
		return ErrBoxRoomIDShouldNotBeEmpty
	}

	b.RoomID = roomID
	return nil
}

func (b *Box) Update(
	name string,
	description *string,
) error {
	err := b.ChangeName(name)
	if err != nil {
		return err
	}

	err = b.ChangeDescription(description)
	if err != nil {
		return err
	}

	return nil
}
