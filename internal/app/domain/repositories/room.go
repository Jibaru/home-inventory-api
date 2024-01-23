package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrCanNotCreateRoom            = errors.New("can not create room")
	ErrCanNotCheckIfRoomExistsByID = errors.New("can not check if room exists by id")
)

type RoomRepository interface {
	Create(room *entities.Room) error
	ExistsByID(id string) (bool, error)
}
