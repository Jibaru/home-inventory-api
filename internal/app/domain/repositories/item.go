package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrItemRepositoryCanNotCreateItem          = errors.New("can not create item")
	ErrItemRepositoryItemNotFound              = errors.New("item not found")
	ErrItemRepositoryCanNotGetByQueryFilters   = errors.New("can not get by query filters")
	ErrItemRepositoryCanNotCountByQueryFilters = errors.New("can not count by query filters")
)

type ItemRepository interface {
	Create(item *entities.Item) error
	GetByID(id string) (*entities.Item, error)
	GetByQueryFilters(queryFilter QueryFilter, pageFilter *PageFilter) ([]*entities.Item, error)
	CountByQueryFilters(queryFilter QueryFilter) (int64, error)
}
