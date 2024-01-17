package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/auth"
	"time"
)

type Generator struct {
	Secret         string
	ExpirationTime time.Duration
}

type CustomClaims struct {
	ID    string `json:"id" mapstructure:"id"`
	Email string `json:"email" mapstructure:"email"`
	jwt.RegisteredClaims
}

func NewJwtGenerator(
	secret string,
	expirationTime time.Duration,
) *Generator {
	return &Generator{
		Secret:         secret,
		ExpirationTime: expirationTime,
	}
}

func (s *Generator) GenerateToken(id string, email string) (string, error) {
	claims := &CustomClaims{
		id,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.ExpirationTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", auth.ErrCanNotGenerateToken
	}

	return encodedToken, nil
}

func (s *Generator) DecodeToken(tokenString string) (*CustomClaims, error) {
	verifier, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("can not verify token")
		}
		return []byte(s.Secret), nil
	})
	if err != nil {
		return nil, errors.New("can not verify token")
	}

	if !verifier.Valid {
		return nil, errors.New("token is not valid")
	}

	claims, ok := verifier.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unable to parse claims")
	}

	return &CustomClaims{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(int64(claims["exp"].(float64)), 0)),
		},
	}, nil
}
