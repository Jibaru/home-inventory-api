package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/dao"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrVersionNotSet = errors.New("current version is not set")
)

type VersionService struct {
	versionDAO *dao.VersionDAO
}

func NewVersionService(versionDAO *dao.VersionDAO) *VersionService {
	return &VersionService{
		versionDAO,
	}
}

func (s *VersionService) GetLatestVersion() (*entities.Version, error) {
	version, err := s.versionDAO.GetLatest()

	if err != nil && errors.Is(err, dao.ErrVersionNotFound) {
		return nil, ErrVersionNotSet
	}

	return version, nil
}
