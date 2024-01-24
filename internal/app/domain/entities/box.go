package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	ErrRoomIDShouldNotBeEmpty = errors.New("room id should not be empty")
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
		return nil, ErrNameShouldNotBeEmpty
	}

	if len(name) > 100 {
		return nil, ErrNameShouldHave100OrLessChars
	}

	if description != nil {
		if strings.TrimSpace(*description) == "" {
			return nil, ErrDescriptionShouldNotBeEmpty
		}

		if len(*description) > 255 {
			return nil, ErrDescriptionShouldHave255OrLessChars
		}
	}

	if strings.TrimSpace(roomID) == "" {
		return nil, ErrRoomIDShouldNotBeEmpty
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
