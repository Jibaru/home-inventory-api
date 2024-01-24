package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type RoomRepositoryMock struct {
	mock.Mock
}

func (m *RoomRepositoryMock) Create(room *entities.Room) error {
	args := m.Called(room)
	return args.Error(0)
}

func (m *RoomRepositoryMock) ExistsByID(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
