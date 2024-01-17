package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type VersionRepositoryMock struct {
	mock.Mock
}

func (m *VersionRepositoryMock) GetLatest() (*entities.Version, error) {
	args := m.Called()

	if args.Get(0) != nil {
		return args.Get(0).(*entities.Version), args.Error(1)
	}

	return nil, args.Error(1)
}
