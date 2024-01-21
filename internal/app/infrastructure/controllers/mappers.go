package controllers

import (
	"io"
	"mime/multipart"
	"os"
)

func mapFileHeaderToTempFolderAndFile(fileHeader *multipart.FileHeader) (string, *os.File, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	tempDir, err := os.MkdirTemp("", "temp_folder")
	if err != nil {
		return "", nil, err
	}

	tempFile, err := os.CreateTemp(tempDir, "*_"+fileHeader.Filename)
	if err != nil {
		return "", nil, err
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return "", nil, err
	}

	return tempDir, tempFile, nil
}
