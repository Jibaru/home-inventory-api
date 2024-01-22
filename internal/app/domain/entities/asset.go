package entities

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"time"
)

var (
	ErrCanNotCreateAssetFromFile = errors.New("can not create asset from file")
)

type Asset struct {
	ID         string
	Name       string
	Extension  string
	Size       int64
	FileID     string
	EntityID   string
	EntityName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewAssetFromFile(
	file *os.File,
	fileID string,
	entity Entity,
) (*Asset, error) {
	name := filepath.Base(file.Name())
	extension := filepath.Ext(file.Name())
	info, err := file.Stat()
	if err != nil {
		return nil, ErrCanNotCreateAssetFromFile
	}

	return &Asset{
		ID:         uuid.NewString(),
		Name:       name,
		Extension:  extension,
		Size:       info.Size(),
		FileID:     fileID,
		EntityID:   entity.EntityID(),
		EntityName: entity.EntityName(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
