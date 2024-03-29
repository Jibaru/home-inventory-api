package main

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/database"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/http"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/jibaru/home-inventory-api/m/notifier"
	"strconv"
	"time"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		logger.LogError(err)
		return
	}

	err = notifier.Init(config.SentryDSN)
	if err != nil {
		logger.LogError(err)
		return
	}
	defer notifier.Flush()

	db, err := database.CreateConnection(
		database.DBConfig{
			Name:     config.DatabaseName,
			Host:     config.DatabaseHost,
			Port:     config.DatabasePort,
			Username: config.DatabaseUsername,
			Password: config.DatabasePassword,
		},
	)
	if err != nil {
		logger.LogError(err)
		return
	}

	http.RunServer(
		config.AppHost,
		strconv.Itoa(config.AppPort),
		config.JwtSecret,
		time.Duration(config.JwtDuration)*time.Hour,
		config.AwsAccessKeyID,
		config.AwsSecretAccessKey,
		config.AwsRegion,
		config.S3BucketName,
		config.SmtpHost,
		config.SmtpPort,
		config.SmtpEmail,
		config.SmtpPassword,
		db,
	)
}
