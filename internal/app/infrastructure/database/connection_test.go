package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBConfigDSN(t *testing.T) {
	dbConfig := DBConfig{
		Name:     "test_db",
		Host:     "localhost",
		Port:     3306,
		Username: "test",
		Password: "password",
	}

	expectedDSN := "test:password@tcp(localhost:3306)/test_db?parseTime=true"
	actualDSN := dbConfig.DSN()

	assert.Equal(t, expectedDSN, actualDSN)
}
