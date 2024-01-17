package stub

import (
	"github.com/stretchr/testify/mock"
)

type TokenGeneratorMock struct {
	mock.Mock
}

func (m *TokenGeneratorMock) GenerateToken(userID, userEmail string) (string, error) {
	args := m.Called(userID, userEmail)
	return args.String(0), args.Error(1)
}

func (s *TokenGeneratorMock) ParseToken(token string) (
	*struct {
		ID    string
		Email string
	},
	error,
) {
	args := s.Called(token)
	if args.Get(0) != nil {
		return args.Get(0).(*struct {
			ID    string
			Email string
		}), args.Error(1)
	}

	return nil, args.Error(1)
}
