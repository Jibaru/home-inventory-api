package services

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
)

type RoomService struct {
	roomRepository repositories.RoomRepository
}

func NewRoomService(
	roomRepository repositories.RoomRepository,
) *RoomService {
	return &RoomService{
		roomRepository,
	}
}

func (s *RoomService) Create(name string, description *string, userID string) (*entities.Room, error) {
	room, err := entities.NewRoom(name, description, userID)
	if err != nil {
		return nil, err
	}

	err = s.roomRepository.Create(room)
	if err != nil {
		return nil, err
	}

	return room, nil
}
