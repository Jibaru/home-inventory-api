package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/jibaru/home-inventory-api/m/notifier"
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
		logger.LogError(err)
		notifier.NotifyError(err)
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
		logger.LogError(err)
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
		logger.LogError(err)
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
		logger.LogError(err)
		return 0, repositories.ErrRoomRepositoryCanNotCountRooms
	}

	return count, nil
}

func (r *RoomRepository) Delete(id string) error {
	err := r.db.Where("id = ?", id).Delete(&entities.Room{}).Error
	if err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrRoomRepositoryCanNotDeleteRoom
	}

	return nil
}

func (r *RoomRepository) GetByID(id string) (*entities.Room, error) {
	var room entities.Room
	err := r.db.Where("id = ?", id).First(&room).Error
	if err != nil {
		logger.LogError(err)
		return nil, repositories.ErrRoomRepositoryCanNotGetRoomByID
	}

	return &room, nil
}

func (r *RoomRepository) Update(room *entities.Room) error {
	err := r.db.Save(room).Error
	if err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return repositories.ErrRoomRepositoryCanNotUpdateRoom
	}

	return nil
}
