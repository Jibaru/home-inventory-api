package services

import (
	"errors"
	"os"
)

var (
	ErrCanNotUploadFile = errors.New("can not upload file")
	ErrUploadingFile    = errors.New("error uploading file")
	ErrCanNotDeleteFile = errors.New("can not delete file")
)

type FileManager interface {
	Upload(file *os.File) (string, error)
	GenerateUrl(id string, extension string) string
	Delete(id string, extension string) error
}
