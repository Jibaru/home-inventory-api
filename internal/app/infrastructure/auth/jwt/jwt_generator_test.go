package jwt

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	secret := "secret-key"
	expirationTime := time.Hour
	jwtGenerator := NewJwtGenerator(secret, expirationTime)

	id := uuid.NewString()
	email := "test@example.com"

	token, err := jwtGenerator.GenerateToken(id, email)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	decodedClaims, err := jwtGenerator.DecodeToken(token)
	assert.NoError(t, err)

	assert.Equal(t, id, decodedClaims.ID)
	assert.Equal(t, email, decodedClaims.Email)
	assert.WithinDuration(t, time.Now().Add(expirationTime), time.Unix(decodedClaims.ExpiresAt.Unix(), 0), 5*time.Second)
}
