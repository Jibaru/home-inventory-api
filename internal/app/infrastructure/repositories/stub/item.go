package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
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
