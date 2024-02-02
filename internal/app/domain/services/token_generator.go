package services

import "errors"

var (
	ErrTokenGeneratorCanNotGenerateToken = errors.New("can not generate token")
	ErrTokenGeneratorCanNotVerifyToken   = errors.New("can not verify token")
	ErrTokenGeneratorTokenIsNotValid     = errors.New("token is not valid")
	ErrTokenGeneratorUnableToParseClaims = errors.New("unable to parse claims")
)

type TokenGenerator interface {
	GenerateToken(id string, email string) (string, error)
	ParseToken(token string) (
		*struct {
			ID    string
			Email string
		},
		error,
	)
}
