package gorm

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func TestVersionRepositoryGetLatestVersion(t *testing.T) {
	db, dbMock := makeDBMock()
	versionRepository := NewVersionRepository(db)

	expectedVersion := entities.Version{
		ID:    uuid.NewString(),
		Value: "1.0.0",
	}

	rows := sqlmock.NewRows([]string{"id", "value"}).
		AddRow(expectedVersion.ID, expectedVersion.Value)

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `versions` ORDER BY `versions`.`id` LIMIT 1")).
		WillReturnRows(rows).
		WithoutArgs()

	version, err := versionRepository.GetLatest()

	assert.NotNil(t, version)
	assert.Equal(t, expectedVersion.ID, version.ID)
	assert.Equal(t, expectedVersion.Value, version.Value)
	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestVersionRepositoryGetLatestVersionErrorVersionNotFound(t *testing.T) {
	db, dbMock := makeDBMock()
	versionRepository := NewVersionRepository(db)

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `versions` ORDER BY `versions`.`id` LIMIT 1")).
		WithoutArgs().
		WillReturnError(gorm.ErrRecordNotFound)

	version, err := versionRepository.GetLatest()

	assert.Nil(t, version)
	assert.Error(t, err)
	assert.ErrorIs(t, repositories.ErrVersionNotFound, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
