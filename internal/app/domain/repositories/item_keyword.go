package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrItemKeywordRepositoryCanNotCreateItemKeywords = errors.New("can not create item keywords")
	ErrItemKeywordRepositoryCanNotDeleteByItemID     = errors.New("can not delete by item id")
)

type ItemKeywordRepository interface {
	CreateMany(itemKeyword []*entities.ItemKeyword) error
	DeleteByItemID(itemID string) error
}
