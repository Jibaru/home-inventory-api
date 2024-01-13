package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
)

var (
	ErrVersionNotSet = errors.New("current version is not set")
)

type VersionService struct {
	versionRepository repositories.VersionRepository
}

func NewVersionService(versionRepository repositories.VersionRepository) *VersionService {
	return &VersionService{
		versionRepository,
	}
}

func (s *VersionService) GetLatestVersion() (*entities.Version, error) {
	version, err := s.versionRepository.GetLatest()

	if err != nil && errors.Is(err, repositories.ErrVersionNotFound) {
		return nil, ErrVersionNotSet
	}

	return version, nil
}
