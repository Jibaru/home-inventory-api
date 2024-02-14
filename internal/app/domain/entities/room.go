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
	room := &Room{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := room.Update(name, description)
	if err != nil {
		return nil, err
	}

	err = room.ChangeUserID(userID)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *Room) Update(name string, description *string) error {
	err := r.ChangeName(name)
	if err != nil {
		return err
	}

	err = r.ChangeDescription(description)
	if err != nil {
		return err
	}

	return nil
}

func (r *Room) ChangeUserID(userID string) error {
	if strings.TrimSpace(userID) == "" {
		return ErrRoomUserIDShouldNotBeEmpty
	}

	r.UserID = userID
	return nil
}

func (r *Room) ChangeName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrRoomNameShouldNotBeEmpty
	}

	if len(name) > 100 {
		return ErrRoomNameShouldHave100OrLessChars
	}

	r.Name = name
	return nil
}

func (r *Room) ChangeDescription(description *string) error {
	if description != nil {
		if strings.TrimSpace(*description) == "" {
			return ErrRoomDescriptionShouldNotBeEmpty
		}

		if len(*description) > 255 {
			return ErrRoomDescriptionShouldHave255OrLessChars
		}
	}

	r.Description = description
	return nil
}
