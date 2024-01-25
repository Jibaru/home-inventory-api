package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrCanNotCreateItem = errors.New("can not create item")
)

type ItemRepository interface {
	Create(item *entities.Item) error
}
