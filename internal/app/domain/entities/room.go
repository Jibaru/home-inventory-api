package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	ErrRoomDescriptionShouldHave255OrLessChars = errors.New("description should have 255 or less characters")
	ErrRoomDescriptionShouldNotBeEmpty         = errors.New("description should not be empty")
	ErrRoomNameShouldHave100OrLessChars        = errors.New("name should have 100 or less characters")
	ErrRoomNameShouldNotBeEmpty                = errors.New("name should not be empty")
	ErrRoomUserIDShouldNotBeEmpty              = errors.New("user id should not be empty")
)

const (
	RoomNameField        = "name"
	RoomDescriptionField = "description"
	RoomUserIDField      = "user_id"
)

type Room struct {
	ID          string
	Name        string
	Description *string
	UserID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewRoom(name string, description *string, userID string) (*Room, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrRoomNameShouldNotBeEmpty
	}

	if len(name) > 100 {
		return nil, ErrRoomNameShouldHave100OrLessChars
	}

	if description != nil {
		if strings.TrimSpace(*description) == "" {
			return nil, ErrRoomDescriptionShouldNotBeEmpty
		}

		if len(*description) > 255 {
			return nil, ErrRoomDescriptionShouldHave255OrLessChars
		}
	}

	if strings.TrimSpace(userID) == "" {
		return nil, ErrRoomUserIDShouldNotBeEmpty
	}

	return &Room{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
