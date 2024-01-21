package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

type MockEntity struct {
	ID string
}

func (e *MockEntity) EntityID() string {
	return e.ID
}

func (e *MockEntity) EntityName() string {
	return "mock_entity"
}

func TestNewAssetFromFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "*_"+uuid.NewString())
	defer tempFile.Close()
	assert.NoError(t, err)

	name := filepath.Base(tempFile.Name())
	extension := filepath.Ext(tempFile.Name())
	info, err := tempFile.Stat()
	assert.NoError(t, err)

	fileID := uuid.NewString()
	entity := &MockEntity{uuid.NewString()}

	asset, err := NewAssetFromFile(tempFile, fileID, entity)

	assert.NoError(t, err)
	assert.NotEmpty(t, asset)
	assert.NotEmpty(t, asset.ID)
	assert.Equal(t, name, asset.Name)
	assert.Equal(t, extension, asset.Extension)
	assert.Equal(t, info.Size(), asset.Size)
	assert.Equal(t, fileID, asset.FileID)
	assert.Equal(t, entity.EntityID(), asset.EntityID)
	assert.Equal(t, entity.EntityName(), asset.EntityName)
	assert.NotEmpty(t, asset.CreatedAt)
	assert.NotEmpty(t, asset.UpdatedAt)
}

func TestNewAssetFromFileErrorCanNotCreateAssetFromFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "*_"+uuid.NewString())
	tempFile.Close()
	assert.NoError(t, err)

	asset, err := NewAssetFromFile(tempFile, uuid.NewString(), &MockEntity{uuid.NewString()})

	assert.Error(t, err)
	assert.Empty(t, asset)
	assert.ErrorIs(t, err, ErrCanNotCreateAssetFromFile)
}
