package services

import (
	"errors"
	"os"
)

var (
	ErrFileManagerCanNotDeleteFile = errors.New("can not delete file")
	ErrFileManagerCanNotUploadFile = errors.New("can not upload file")
	ErrFileManagerUploadingFile    = errors.New("error uploading file")
)

type FileManager interface {
	Upload(file *os.File) (string, error)
	GenerateUrl(id string, extension string) string
	Delete(id string, extension string) error
}
