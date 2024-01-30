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

func (m *BoxRepositoryMock) GetBoxItem(boxID string, itemID string) (*entities.BoxItem, error) {
	args := m.Called(boxID, itemID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.BoxItem), args.Error(1)
}

func (m *BoxRepositoryMock) CreateBoxItem(boxItem *entities.BoxItem) error {
	args := m.Called(boxItem)
	return args.Error(0)
}

func (m *BoxRepositoryMock) UpdateBoxItem(boxItem *entities.BoxItem) error {
	args := m.Called(boxItem)
	return args.Error(0)
}

func (m *BoxRepositoryMock) CreateBoxTransaction(boxTransaction *entities.BoxTransaction) error {
	args := m.Called(boxTransaction)
	return args.Error(0)
}
