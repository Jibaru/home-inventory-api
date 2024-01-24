package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type BoxRepositoryMock struct {
	mock.Mock
}

func (m *BoxRepositoryMock) Create(box *entities.Box) error {
	args := m.Called(box)
	return args.Error(0)
}
