package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/stub"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionServiceGetLatestVersion(t *testing.T) {
	mockRepo := new(stub.VersionRepositoryMock)
	versionService := NewVersionService(mockRepo)

	expectedVersion := &entities.Version{Value: "1.0.0", ID: uuid.NewString()}
	mockRepo.On("GetLatest").Return(expectedVersion, nil)

	version, err := versionService.GetLatestVersion()

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetLatest")
	assert.Equal(t, expectedVersion, version)
}

func TestVersionServiceGetLatestVersionErrorNotFound(t *testing.T) {
	mockRepo := new(stub.VersionRepositoryMock)
	versionService := NewVersionService(mockRepo)

	mockRepo.On("GetLatest").Return(nil, repositories.ErrVersionNotFound).Once()

	version, err := versionService.GetLatestVersion()

	assert.Error(t, err)
	assert.EqualError(t, err, ErrVersionNotSet.Error())
	mockRepo.AssertCalled(t, "GetLatest")
	assert.Nil(t, version)
}

func TestVersionServiceGetLatestVersionErrorCanNotGetLatestVersion(t *testing.T) {
	mockRepo := new(stub.VersionRepositoryMock)
	versionService := NewVersionService(mockRepo)

	mockRepo.On("GetLatest").Return(nil, errors.New("repository error")).Once()

	version, err := versionService.GetLatestVersion()

	assert.Error(t, err)
	assert.EqualError(t, err, ErrCanNotGetLatestVersion.Error())
	mockRepo.AssertCalled(t, "GetLatest")
	assert.Nil(t, version)
}
