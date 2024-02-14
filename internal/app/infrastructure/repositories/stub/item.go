package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/stretchr/testify/mock"
)

type ItemRepositoryMock struct {
	mock.Mock
}

func (r *ItemRepositoryMock) Create(item *entities.Item) error {
	args := r.Called(item)
	return args.Error(0)
}

func (r *ItemRepositoryMock) GetByID(id string) (*entities.Item, error) {
	args := r.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Item), args.Error(1)
}

func (r *ItemRepositoryMock) GetByQueryFilters(
	queryFilter repositories.QueryFilter,
	pageFilter *repositories.PageFilter,
) ([]*entities.Item, error) {
	args := r.Called(queryFilter, pageFilter)

	if data := args.Get(0); data != nil {
		return data.([]*entities.Item), args.Error(1)
	}

	return nil, args.Error(1)
}

func (r *ItemRepositoryMock) CountByQueryFilters(
	queryFilter repositories.QueryFilter,
) (int64, error) {
	args := r.Called(queryFilter)
	return args.Get(0).(int64), args.Error(1)
}

func (r *ItemRepositoryMock) Update(item *entities.Item) error {
	args := r.Called(item)
	return args.Error(0)
}
