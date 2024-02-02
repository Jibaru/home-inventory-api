package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrItemRepositoryCanNotCreateItem = errors.New("can not create item")
	ErrItemRepositoryItemNotFound     = errors.New("item not found")
)

type ItemRepository interface {
	Create(item *entities.Item) error
	GetByID(id string) (*entities.Item, error)
}
