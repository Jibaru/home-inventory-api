package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type ItemKeywordRepositoryMock struct {
	mock.Mock
}

func (r *ItemKeywordRepositoryMock) CreateMany(itemKeywords []*entities.ItemKeyword) error {
	args := r.Called(itemKeywords)
	return args.Error(0)
}

func (r *ItemKeywordRepositoryMock) DeleteByItemID(itemID string) error {
	args := r.Called(itemID)
	return args.Error(0)
}
