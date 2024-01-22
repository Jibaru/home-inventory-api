package gorm

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func makeDBMock() (*gorm.DB, sqlmock.Sqlmock) {
	dialector, dbMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dialector,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	return db, dbMock
}
