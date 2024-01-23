package repositories

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
)

var (
	ErrCanNotCreateBox = errors.New("can not create box")
)

type BoxRepository interface {
	Create(box *entities.Box) error
}
