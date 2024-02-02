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
		return repositories.ErrRoomRepositoryCanNotCreateRoom
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
		return false, repositories.ErrRoomRepositoryCanNotCheckIfRoomExistsByID
	}

	return count > 0, nil
}

func (r *RoomRepository) GetByQueryFilters(queryFilter repositories.QueryFilter, pageFilter *repositories.PageFilter) ([]*entities.Room, error) {
	var rooms []*entities.Room
	err := applyFilters(r.db, queryFilter).
		Offset(pageFilter.Offset).
		Limit(pageFilter.Limit).
		Find(&rooms).
		Error

	if err != nil {
		return nil, repositories.ErrRoomRepositoryCanNotGetRooms
	}

	return rooms, nil
}

func (r *RoomRepository) CountByQueryFilters(queryFilter repositories.QueryFilter) (int64, error) {
	var count int64
	err := applyFilters(r.db.Model(&entities.Room{}), queryFilter).
		Count(&count).
		Error

	if err != nil {
		return 0, repositories.ErrRoomRepositoryCanNotCountRooms
	}

	return count, nil
}
