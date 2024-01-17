package auth

import "errors"

var (
	ErrCanNotGenerateToken = errors.New("can not generate token")
)

type TokenGenerator interface {
	GenerateToken(id string, email string) (string, error)
}