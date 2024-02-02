package gorm

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"gorm.io/gorm"
)

type VersionRepository struct {
	db *gorm.DB
}

func NewVersionRepository(db *gorm.DB) *VersionRepository {
	return &VersionRepository{
		db,
	}
}

func (r *VersionRepository) GetLatest() (*entities.Version, error) {
	version := &entities.Version{}
	err := r.db.First(version).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repositories.ErrVersionRepositoryVersionNotFound
	}

	return version, nil
}
