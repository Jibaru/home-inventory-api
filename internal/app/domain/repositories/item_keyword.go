package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrItemKeywordRepositoryCanNotCreateItemKeywords = errors.New("can not create item keywords")
)

type ItemKeywordRepository interface {
	CreateMany(itemKeyword []*entities.ItemKeyword) error
}
