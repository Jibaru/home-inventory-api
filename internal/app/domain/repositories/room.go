package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrCanNotCreateRoom = errors.New("can not create room")
)

type RoomRepository interface {
	Create(room *entities.Room) error
}
