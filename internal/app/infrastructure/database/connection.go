package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

func CreateConnection(
	config DBConfig,
) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DSN()), &gorm.Config{
		Logger: NewGormLogger(),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
