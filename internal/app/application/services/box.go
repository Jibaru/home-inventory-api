package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
)

var (
	ErrRoomDoesNotExists = errors.New("room does not exists")
)

type BoxService struct {
	boxRepository  repositories.BoxRepository
	roomRepository repositories.RoomRepository
}

func NewBoxService(
	boxRepository repositories.BoxRepository,
	roomRepository repositories.RoomRepository,
) *BoxService {
	return &BoxService{
		boxRepository,
		roomRepository,
	}
}

func (s *BoxService) Create(name string, description *string, roomID string) (*entities.Box, error) {
	exists, err := s.roomRepository.ExistsByID(roomID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrRoomDoesNotExists
	}

	box, err := entities.NewBox(name, description, roomID)
	if err != nil {
		return nil, err
	}

	err = s.boxRepository.Create(box)
	if err != nil {
		return nil, err
	}

	return box, nil
}
