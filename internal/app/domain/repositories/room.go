package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrCanNotCreateRoom               = errors.New("can not create room")
	ErrCanNotCheckIfRoomExistsByID    = errors.New("can not check if room exists by id")
	ErrRoomRepositoryCanNotGetRooms   = errors.New("can not get rooms")
	ErrRoomRepositoryCanNotCountRooms = errors.New("can not count rooms")
)

type RoomRepository interface {
	Create(room *entities.Room) error
	ExistsByID(id string) (bool, error)
	GetByQueryFilters(queryFilter QueryFilter, pageFilter *PageFilter) ([]*entities.Room, error)
	CountByQueryFilters(queryFilter QueryFilter) (int64, error)
}
