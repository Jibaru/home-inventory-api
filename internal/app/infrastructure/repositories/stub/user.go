package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByEmail(email string) (*entities.User, error) {
	args := m.Called(email)

	if args.Get(0) != nil {
		return args.Get(0).(*entities.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetUserByBoxID(boxID string) (*entities.User, error) {
	args := m.Called(boxID)

	if args.Get(0) != nil {
		return args.Get(0).(*entities.User), args.Error(1)
	}

	return nil, args.Error(1)
}
