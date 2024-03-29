package main

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppHost            string `mapstructure:"APP_HOST"`
	AppPort            int    `mapstructure:"APP_PORT"`
	DatabaseName       string `mapstructure:"DB_NAME"`
	DatabaseHost       string `mapstructure:"DB_HOST"`
	DatabasePort       int    `mapstructure:"DB_PORT"`
	DatabaseUsername   string `mapstructure:"DB_USERNAME"`
	DatabasePassword   string `mapstructure:"DB_PASSWORD"`
	JwtSecret          string `mapstructure:"JWT_SECRET"`
	JwtDuration        int    `mapstructure:"JWT_DURATION"`
	AwsAccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsRegion          string `mapstructure:"AWS_REGION"`
	S3BucketName       string `mapstructure:"S3_BUCKET_NAME"`
	SentryDSN          string `mapstructure:"SENTRY_DSN"`
	SmtpHost           string `mapstructure:"SMTP_HOST"`
	SmtpPort           int    `mapstructure:"SMTP_PORT"`
	SmtpEmail          string `mapstructure:"SMTP_EMAIL"`
	SmtpPassword       string `mapstructure:"SMTP_PASSWORD"`
}

func ReadConfig() (*AppConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &AppConfig{}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
