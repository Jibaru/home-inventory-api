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
	if strings.TrimSpace(name) == "" {
		return nil, ErrBoxNameShouldNotBeEmpty
	}

	if len(name) > 100 {
		return nil, ErrBoxNameShouldHave100OrLessChars
	}

	if description != nil {
		if strings.TrimSpace(*description) == "" {
			return nil, ErrBoxDescriptionShouldNotBeEmpty
		}

		if len(*description) > 255 {
			return nil, ErrBoxDescriptionShouldHave255OrLessChars
		}
	}

	if strings.TrimSpace(roomID) == "" {
		return nil, ErrBoxRoomIDShouldNotBeEmpty
	}

	return &Box{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		RoomID:      roomID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
