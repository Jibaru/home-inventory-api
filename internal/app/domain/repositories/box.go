package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrBoxRepositoryBoxItemNotFound                    = errors.New("box item not found")
	ErrBoxRepositoryCanBotCreateBoxItem                = errors.New("can not create box item")
	ErrBoxRepositoryCanNotCreateBox                    = errors.New("can not create box")
	ErrBoxRepositoryCanNotCreateBoxTransaction         = errors.New("can not create box transaction")
	ErrBoxRepositoryCanNotDeleteBox                    = errors.New("can not delete box")
	ErrBoxRepositoryCanNotDeleteBoxItem                = errors.New("can not delete box item")
	ErrBoxRepositoryCanNotDeleteBoxItemsByBoxID        = errors.New("can not delete box items by box id")
	ErrBoxRepositoryCanNotDeleteBoxTransactionsByBoxID = errors.New("can not delete box transactions by box id")
	ErrBoxRepositoryCanNotUpdateBoxItem                = errors.New("can not update box item")
	ErrorBoxRepositoryCanNotCountByQueryFilters        = errors.New("can not count by query filters")
	ErrorBoxRepositoryCanNotGetByQueryFilters          = errors.New("can not get by query filters")
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
	DeleteBoxItemsByBoxID(boxID string) error
	DeleteBoxTransactionsByBoxID(boxID string) error
	Delete(id string) error
}
