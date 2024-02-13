package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrRoomRepositoryCanNotCheckIfRoomExistsByID = errors.New("can not check if room exists by id")
	ErrRoomRepositoryCanNotCountRooms            = errors.New("can not count rooms")
	ErrRoomRepositoryCanNotCreateRoom            = errors.New("can not create room")
	ErrRoomRepositoryCanNotDeleteRoom            = errors.New("can not delete room")
	ErrRoomRepositoryCanNotGetRoomByID           = errors.New("can not get room by id")
	ErrRoomRepositoryCanNotGetRooms              = errors.New("can not get rooms")
	ErrRoomRepositoryCanNotUpdateRoom            = errors.New("can not update room")
)

type RoomRepository interface {
	Create(room *entities.Room) error
	ExistsByID(id string) (bool, error)
	GetByQueryFilters(queryFilter QueryFilter, pageFilter *PageFilter) ([]*entities.Room, error)
	CountByQueryFilters(queryFilter QueryFilter) (int64, error)
	Delete(id string) error
	GetByID(id string) (*entities.Room, error)
	Update(room *entities.Room) error
}
