package aws

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
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

func TestGenerateUrl(t *testing.T) {
	bucketName := "bucket-test"
	id := uuid.NewString()
	extension := ".png"

	expectedUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s%s",
		bucketName,
		id,
		extension,
	)

	manager := NewFileManager(uuid.NewString(), uuid.NewString(), random.String(8), bucketName)

	url := manager.GenerateUrl(id, extension)

	assert.Equal(t, expectedUrl, url)
}
