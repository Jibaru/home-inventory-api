package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrVersionRepositoryVersionNotFound = errors.New("version not found")
)

type VersionRepository interface {
	GetLatest() (*entities.Version, error)
}
