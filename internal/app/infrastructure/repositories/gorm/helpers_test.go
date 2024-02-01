package gorm

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

type FakeTable struct{}

func TestApplyFiltersWithEmptyQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()

	queryFilter := repositories.QueryFilter{}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `fake_tables`"))

	db = applyFilters(db.Model(&FakeTable{}), queryFilter)
	db.Find(&FakeTable{})

	err := dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestApplyFiltersWithSingleQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Conditions: []repositories.Condition{
					{
						Field:    "name",
						Operator: repositories.LikeComparisonOperator,
						Value:    "test",
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `fake_tables` WHERE name LIKE ?")).
		WithArgs("test")

	db = applyFilters(db.Model(&FakeTable{}), queryFilter)
	db.Find(&FakeTable{})

	err := dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestApplyFiltersWithMultipleQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Conditions: []repositories.Condition{
					{
						Field:    "name",
						Operator: repositories.LikeComparisonOperator,
						Value:    "test",
					},
				},
			},
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "age",
						Operator: repositories.EqualComparisonOperator,
						Value:    20,
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `fake_tables` WHERE name LIKE ? AND age = ?")).
		WithArgs("test", 20)

	db = applyFilters(db.Model(&FakeTable{}), queryFilter)
	db.Find(&FakeTable{})

	err := dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestApplyFiltersWithMultipleQueryFiltersAndOrLogicalOperator(t *testing.T) {
	db, dbMock := makeDBMock()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "name",
						Operator: repositories.LikeComparisonOperator,
						Value:    "test",
					},
					{
						Field:    "is_ok",
						Operator: repositories.EqualComparisonOperator,
						Value:    true,
					},
				},
			},
			{
				Operator: repositories.OrLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "age",
						Operator: repositories.EqualComparisonOperator,
						Value:    20,
					},
					{
						Field:    "stars",
						Operator: repositories.EqualComparisonOperator,
						Value:    10,
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `fake_tables` WHERE (name LIKE ? AND is_ok = ?) AND (age = ? OR stars = ?)")).
		WithArgs("test", true, 20, 10)

	db = applyFilters(db.Model(&FakeTable{}), queryFilter)
	db.Find(&FakeTable{})

	err := dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
