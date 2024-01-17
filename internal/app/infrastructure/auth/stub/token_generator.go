package stub

import "github.com/stretchr/testify/mock"

type TokenGeneratorMock struct {
	mock.Mock
}

func (m *TokenGeneratorMock) GenerateToken(userID, userEmail string) (string, error) {
	args := m.Called(userID, userEmail)
	return args.String(0), args.Error(1)
}
