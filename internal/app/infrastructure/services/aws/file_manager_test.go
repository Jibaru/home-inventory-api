package aws

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestFileManagerUpload(t *testing.T) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("S3_BUCKET_NAME")
	filePath := os.Getenv("FILE_MANAGER_TEST_FILE_PATH")

	if accessKey == "" ||
		secretKey == "" ||
		region == "" ||
		bucketName == "" ||
		filePath == "" {
		log.Println("all required environment variables are not set")
		return
	}

	manager := NewFileManager(accessKey, secretKey, region, bucketName)

	file, err := os.Open(filePath)
	assert.NoError(t, err)
	assert.NotEmpty(t, file)

	id, err := manager.Upload(file)

	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestFileManagerGenerateUrl(t *testing.T) {
	bucketName := "bucket-test"
	id := uuid.NewString()
	extension := ".png"

	expectedUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s%s",
		bucketName,
		id,
		extension)

	manager := NewFileManager(uuid.NewString(), uuid.NewString(), random.String(8), bucketName)

	url := manager.GenerateUrl(id, extension)

	assert.Equal(t, expectedUrl, url)
}

func TestFileManagerDelete(t *testing.T) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("S3_BUCKET_NAME")
	filePath := os.Getenv("FILE_MANAGER_TEST_FILE_PATH")
	validObjectID := os.Getenv("FILE_MANAGER_TEST_VALID_OBJECT_ID")
	validExtension := os.Getenv("FILE_MANAGER_TEST_VALID_EXTENSION")

	if accessKey == "" ||
		secretKey == "" ||
		region == "" ||
		bucketName == "" ||
		filePath == "" {
		log.Println("all required environment variables are not set")
		return
	}

	manager := NewFileManager(accessKey, secretKey, region, bucketName)

	err := manager.Delete(validObjectID, validExtension)

	assert.NoError(t, err)
}

func TestFileManagerDeleteErrorCanNotDeleteFile(t *testing.T) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("S3_BUCKET_NAME")
	filePath := os.Getenv("FILE_MANAGER_TEST_FILE_PATH")
	invalidObjectID := os.Getenv("FILE_MANAGER_TEST_INVALID_OBJECT_ID")
	invalidExtension := os.Getenv("FILE_MANAGER_TEST_INVALID_EXTENSION")

	if accessKey == "" ||
		secretKey == "" ||
		region == "" ||
		bucketName == "" ||
		filePath == "" {
		log.Println("all required environment variables are not set")
		return
	}

	manager := NewFileManager(accessKey, secretKey, region, bucketName)

	err := manager.Delete(invalidObjectID, invalidExtension)

	assert.Error(t, err)
	assert.ErrorIs(t, err, services.ErrFileManagerCanNotDeleteFile)
}
