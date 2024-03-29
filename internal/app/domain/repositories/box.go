package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrBoxRepositoryBoxItemNotFound                          = errors.New("box item not found")
	ErrBoxRepositoryCanBotCreateBoxItem                      = errors.New("can not create box item")
	ErrBoxRepositoryCanNotCreateBox                          = errors.New("can not create box")
	ErrBoxRepositoryCanNotCreateBoxTransaction               = errors.New("can not create box transaction")
	ErrBoxRepositoryCanNotDeleteBox                          = errors.New("can not delete box")
	ErrBoxRepositoryCanNotDeleteBoxItem                      = errors.New("can not delete box item")
	ErrBoxRepositoryCanNotDeleteBoxItemsByBoxID              = errors.New("can not delete box items by box id")
	ErrBoxRepositoryCanNotDeleteBoxTransactionsByBoxID       = errors.New("can not delete box transactions by box id")
	ErrBoxRepositoryCanNotUpdateBoxItem                      = errors.New("can not update box item")
	ErrBoxRepositoryCanNotCountByQueryFilters                = errors.New("can not count by query filters")
	ErrBoxRepositoryCanNotGetByQueryFilters                  = errors.New("can not get by query filters")
	ErrBoxRepositoryCanNotGetByID                            = errors.New("can not get by id")
	ErrBoxRepositoryCanNotUpdate                             = errors.New("can not update")
	ErrBoxRepositoryCanNotGetBoxTransactionsByQueryFilters   = errors.New("can not get box transactions by query filters")
	ErrBoxRepositoryCanNotCountBoxTransactionsByQueryFilters = errors.New("can not count box transactions by query filters")
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
	GetByID(id string) (*entities.Box, error)
	Update(box *entities.Box) error
	GetBoxTransactionsByQueryFilters(queryFilter QueryFilter, pageFilter *PageFilter) ([]*entities.BoxTransaction, error)
	CountBoxTransactionsByQueryFilters(queryFilter QueryFilter) (int64, error)
}
