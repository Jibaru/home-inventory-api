package repositories

import "github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
}
