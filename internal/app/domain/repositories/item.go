package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrItemRepositoryCanNotCountByQueryFilters = errors.New("can not count by query filters")
	ErrItemRepositoryCanNotCreateItem          = errors.New("can not create item")
	ErrItemRepositoryCanNotGetByQueryFilters   = errors.New("can not get by query filters")
	ErrItemRepositoryCanNotUpdateItem          = errors.New("can not update item")
	ErrItemRepositoryItemNotFound              = errors.New("item not found")
)

type ItemRepository interface {
	Create(item *entities.Item) error
	GetByID(id string) (*entities.Item, error)
	GetByQueryFilters(queryFilter QueryFilter, pageFilter *PageFilter) ([]*entities.Item, error)
	CountByQueryFilters(queryFilter QueryFilter) (int64, error)
	Update(item *entities.Item) error
}
