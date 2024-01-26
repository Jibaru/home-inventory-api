package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/stretchr/testify/mock"
)

type AssetRepositoryMock struct {
	mock.Mock
}

func (r *AssetRepositoryMock) Create(asset *entities.Asset) error {
	args := r.Called(asset)
	return args.Error(0)
}

func (r *AssetRepositoryMock) FindByEntity(
	entity entities.Entity,
	page *repositories.PageFilter,
) ([]*entities.Asset, error) {
	args := r.Called(entity, page)

	if data := args.Get(0); data != nil {
		return data.([]*entities.Asset), args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *AssetRepositoryMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
