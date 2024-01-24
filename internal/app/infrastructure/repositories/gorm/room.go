package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		db,
	}
}

func (r *RoomRepository) Create(room *entities.Room) error {
	if err := r.db.Create(room).Error; err != nil {
		return repositories.ErrCanNotCreateRoom
	}

	return nil
}

func (r *RoomRepository) ExistsByID(id string) (bool, error) {
	var count int64
	err := r.db.Model(&entities.Room{}).
		Where("id = ?", id).
		Count(&count).
		Error
	if err != nil {
		return false, repositories.ErrCanNotCheckIfRoomExistsByID
	}

	return count > 0, nil
}
