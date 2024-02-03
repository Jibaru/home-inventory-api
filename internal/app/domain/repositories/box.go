package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrBoxRepositoryBoxItemNotFound             = errors.New("box item not found")
	ErrBoxRepositoryCanBotCreateBoxItem         = errors.New("can not create box item")
	ErrBoxRepositoryCanNotCreateBox             = errors.New("can not create box")
	ErrBoxRepositoryCanNotCreateBoxTransaction  = errors.New("can not create box transaction")
	ErrBoxRepositoryCanNotDeleteBoxItem         = errors.New("can not delete box item")
	ErrBoxRepositoryCanNotUpdateBoxItem         = errors.New("can not update box item")
	ErrorBoxRepositoryCanNotGetByQueryFilters   = errors.New("can not get by query filters")
	ErrorBoxRepositoryCanNotCountByQueryFilters = errors.New("can not count by query filters")
)

type BoxRepository interface {
	Create(box *entities.Box) error
	GetBoxItem(boxID string, itemID string) (*entities.BoxItem, error)
	CreateBoxItem(boxItem *entities.BoxItem) error
	UpdateBoxItem(boxItem *entities.BoxItem) error
	CreateBoxTransaction(boxTransaction *entities.BoxTransaction) error
	DeleteBoxItem(boxID string, itemID string) error
	GetByQueryFilters(queryFilter QueryFilter, pageFilter *PageFilter) ([]*entities.Box, error)
	CountByQueryFilters(queryFilter QueryFilter) (int64, error)
}
