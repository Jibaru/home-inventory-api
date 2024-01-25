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
