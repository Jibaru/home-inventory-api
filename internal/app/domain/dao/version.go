package dao

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"gorm.io/gorm"
)

var (
	ErrVersionNotFound = errors.New("version not found")
)

type VersionDAO struct {
	DB *gorm.DB
}

func (d *VersionDAO) Create(version entities.Version) {
	d.DB.Create(&version)
}

func (d *VersionDAO) GetLatest() (*entities.Version, error) {
	version := &entities.Version{}
	err := d.DB.First(version).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrVersionNotFound
	}

	return version, nil
}
