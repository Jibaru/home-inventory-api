package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockVersionRepository struct {
	mock.Mock
}

func (m *mockVersionRepository) GetLatest() (*entities.Version, error) {
	args := m.Called()

	if args.Get(0) != nil {
		return args.Get(0).(*entities.Version), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestGetLatestVersion(t *testing.T) {
	mockRepo := new(mockVersionRepository)
	versionService := NewVersionService(mockRepo)

	expectedVersion := &entities.Version{Value: "1.0.0", ID: uuid.NewString()}
	mockRepo.On("GetLatest").Return(expectedVersion, nil)

	version, err := versionService.GetLatestVersion()

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetLatest")
	assert.Equal(t, expectedVersion, version)
}

func TestGetLatestVersionErrorNotFound(t *testing.T) {
	mockRepo := new(mockVersionRepository)
	versionService := NewVersionService(mockRepo)

	mockRepo.On("GetLatest").Return(nil, repositories.ErrVersionNotFound).Once()

	version, err := versionService.GetLatestVersion()

	assert.Error(t, err)
	assert.EqualError(t, err, ErrVersionNotSet.Error())
	mockRepo.AssertCalled(t, "GetLatest")
	assert.Nil(t, version)
}

func TestGetLatestVersionErrorCanNotGetLatestVersion(t *testing.T) {
	mockRepo := new(mockVersionRepository)
	versionService := NewVersionService(mockRepo)

	mockRepo.On("GetLatest").Return(nil, errors.New("repository error")).Once()

	version, err := versionService.GetLatestVersion()

	assert.Error(t, err)
	assert.EqualError(t, err, ErrCanNotGetLatestVersion.Error())
	mockRepo.AssertCalled(t, "GetLatest")
	assert.Nil(t, version)
}
