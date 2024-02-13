package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
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

func (m *RoomRepositoryMock) GetByQueryFilters(
	queryFilter repositories.QueryFilter,
	pageFilter *repositories.PageFilter,
) ([]*entities.Room, error) {
	args := m.Called(queryFilter, pageFilter)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*entities.Room), args.Error(1)
}

func (m *RoomRepositoryMock) CountByQueryFilters(
	queryFilter repositories.QueryFilter,
) (int64, error) {
	args := m.Called(queryFilter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *RoomRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *RoomRepositoryMock) GetByID(id string) (*entities.Room, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Room), args.Error(1)
}

func (m *RoomRepositoryMock) Update(room *entities.Room) error {
	args := m.Called(room)
	return args.Error(0)
}
