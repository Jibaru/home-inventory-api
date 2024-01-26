package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"os"
	"path/filepath"
)

type FileManager struct {
	accessKey  string
	secretKey  string
	region     string
	bucketName string
}

func NewFileManager(
	accessKey string,
	secretKey string,
	region string,
	bucketName string,
) *FileManager {
	return &FileManager{
		accessKey,
		secretKey,
		region,
		bucketName,
	}
}

func (m *FileManager) getUploader() (*s3manager.Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(m.region),
		Credentials: credentials.NewStaticCredentialsFromCreds(
			credentials.Value{
				AccessKeyID:     m.accessKey,
				SecretAccessKey: m.secretKey,
			},
		),
	})
	if err != nil {
		return nil, err
	}

	return s3manager.NewUploader(sess), nil
}

func (m *FileManager) getNewS3Client() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(m.region),
		Credentials: credentials.NewStaticCredentialsFromCreds(
			credentials.Value{
				AccessKeyID:     m.accessKey,
				SecretAccessKey: m.secretKey,
			},
		),
	})
	if err != nil {
		return nil, err
	}

	return s3.New(sess), nil
}

func (m *FileManager) Upload(file *os.File) (string, error) {
	uploader, err := m.getUploader()
	if err != nil {
		return "", services.ErrCanNotUploadFile
	}

	id := uuid.NewString()

	extension := filepath.Ext(file.Name())

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(m.bucketName),
		Key:    aws.String(id + extension),
		Body:   file,
	})
	if err != nil {
		return "", services.ErrUploadingFile
	}

	return id, nil
}

func (m *FileManager) GenerateUrl(id string, extension string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s%s",
		m.bucketName,
		id,
		extension,
	)
}

func (m *FileManager) Delete(id string, extension string) error {
	client, err := m.getNewS3Client()
	if err != nil {
		return services.ErrCanNotDeleteFile
	}

	_, err = client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(m.bucketName),
		Key:    aws.String(id + extension),
	})
	if err != nil {
		return services.ErrCanNotDeleteFile
	}

	err = client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(m.bucketName),
		Key:    aws.String(id + extension),
	})
	if err != nil {
		return services.ErrCanNotDeleteFile
	}

	return nil
}
