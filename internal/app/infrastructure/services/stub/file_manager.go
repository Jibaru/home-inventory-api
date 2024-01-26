package stub

import (
	"github.com/stretchr/testify/mock"
	"os"
)

type FileManagerMock struct {
	mock.Mock
}

func (m *FileManagerMock) Upload(file *os.File) (string, error) {
	args := m.Called(file)
	return args.String(0), args.Error(1)
}

func (m *FileManagerMock) GenerateUrl(id string, extension string) string {
	args := m.Called(id, extension)
	return args.String(0)
}

func (m *FileManagerMock) Delete(id string, extension string) error {
	args := m.Called(id, extension)
	return args.Error(0)
}
