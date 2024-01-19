package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	ErrNameShouldNotBeEmpty                = errors.New("name should not be empty")
	ErrNameShouldHave100OrLessChars        = errors.New("name should have 100 or less characters")
	ErrDescriptionShouldNotBeEmpty         = errors.New("description should not be empty")
	ErrDescriptionShouldHave255OrLessChars = errors.New("description should have 255 or less characters")
	ErrUserIDShouldNotBeEmpty              = errors.New("user id should not be empty")
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

	if strings.TrimSpace(userID) == "" {
		return nil, ErrUserIDShouldNotBeEmpty
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
