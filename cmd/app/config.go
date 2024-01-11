package main

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppHost          string `mapstructure:"APP_HOST"`
	AppPort          int    `mapstructure:"APP_PORT"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     int    `mapstructure:"DB_PORT"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
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
