package gorm

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestItemRepositoryCreate(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	description := random.String(255, random.Alphanumeric)
	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `items` (`id`,`sku`,`name`,`description`,`unit`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(item.ID, item.Sku, item.Name, *item.Description, item.Unit, item.UserID, item.CreatedAt, item.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := itemRepository.Create(item)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryCreateErrorCanNotCreateItem(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	description := random.String(255, random.Alphanumeric)
	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `items` (`id`,`sku`,`name`,`description`,`unit`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(item.ID, item.Sku, item.Name, *item.Description, item.Unit, item.UserID, item.CreatedAt, item.UpdatedAt).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := itemRepository.Create(item)

	assert.Error(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryGetByID(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	description := random.String(255, random.Alphanumeric)
	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id", "sku", "name", "description", "unit", "user_id", "created_at", "updated_at"}).
		AddRow(item.ID, item.Sku, item.Name, item.Description, item.Unit, item.UserID, item.CreatedAt, item.UpdatedAt)
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `items` WHERE id = ? ORDER BY `items`.`id` LIMIT 1")).
		WithArgs(item.ID).
		WillReturnRows(rows)

	itemResult, err := itemRepository.GetByID(item.ID)

	assert.NoError(t, err)
	assert.Equal(t, item.ID, itemResult.ID)
	assert.Equal(t, item.Sku, itemResult.Sku)
	assert.Equal(t, item.Name, itemResult.Name)
	assert.Equal(t, item.Description, itemResult.Description)
	assert.Equal(t, item.Unit, itemResult.Unit)
	assert.Equal(t, item.UserID, itemResult.UserID)
	assert.Equal(t, item.CreatedAt, itemResult.CreatedAt)
	assert.Equal(t, item.UpdatedAt, itemResult.UpdatedAt)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryGetByIDErrorCanNotGetItem(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	itemID := uuid.NewString()
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `items` WHERE id = ? ORDER BY `items`.`id` LIMIT 1")).
		WithArgs(itemID).
		WillReturnError(errors.New("database error"))

	itemResult, err := itemRepository.GetByID(itemID)

	assert.Error(t, err)
	assert.Nil(t, itemResult)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryGetByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	userID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.OrLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "sku",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    "name",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    "description",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
				},
			},
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}

	pageFilter := repositories.PageFilter{
		Offset: 1,
		Limit:  10,
	}

	description := random.String(255, random.Alphanumeric)
	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		Unit:        "unit",
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	itemKeyword := &entities.ItemKeyword{
		ID:        uuid.NewString(),
		Value:     "keyword",
		ItemID:    item.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id", "sku", "name", "description", "unit", "user_id", "created_at", "updated_at"}).
		AddRow(item.ID, item.Sku, item.Name, item.Description, item.Unit, item.UserID, item.CreatedAt, item.UpdatedAt)
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT `items`.`id`,`items`.`sku`,`items`.`name`,`items`.`description`,`items`.`unit`,`items`.`user_id`,`items`.`created_at`,`items`.`updated_at` FROM `items` inner join item_keywords on item_keywords.item_id = items.id WHERE (sku LIKE ? OR name LIKE ? OR description LIKE ?) AND user_id = ? LIMIT 10 OFFSET 1")).
		WithArgs("%search%", "%search%", "%search%", item.UserID).
		WillReturnRows(rows)
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `item_keywords` WHERE `item_keywords`.`item_id` = ?")).
		WithArgs(item.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "value", "item_id", "created_at", "updated_at"}).
			AddRow(itemKeyword.ID, itemKeyword.Value, itemKeyword.ItemID, itemKeyword.CreatedAt, itemKeyword.UpdatedAt))

	itemResults, err := itemRepository.GetByQueryFilters(queryFilter, &pageFilter)

	assert.NoError(t, err)
	assert.Len(t, itemResults, 1)
	assert.Equal(t, item.ID, itemResults[0].ID)
	assert.Equal(t, item.Sku, itemResults[0].Sku)
	assert.Equal(t, item.Name, itemResults[0].Name)
	assert.Equal(t, item.Description, itemResults[0].Description)
	assert.Equal(t, item.Unit, itemResults[0].Unit)
	assert.Equal(t, item.UserID, itemResults[0].UserID)
	assert.Equal(t, item.CreatedAt, itemResults[0].CreatedAt)
	assert.Equal(t, item.UpdatedAt, itemResults[0].UpdatedAt)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryGetByQueryFiltersErrorCanNotGetByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	userID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.OrLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "sku",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    "name",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    "description",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
				},
			},
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}

	pageFilter := repositories.PageFilter{
		Offset: 1,
		Limit:  10,
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT `items`.`id`,`items`.`sku`,`items`.`name`,`items`.`description`,`items`.`unit`,`items`.`user_id`,`items`.`created_at`,`items`.`updated_at` FROM `items` inner join item_keywords on item_keywords.item_id = items.id WHERE (sku LIKE ? OR name LIKE ? OR description LIKE ?) AND user_id = ? LIMIT 10 OFFSET 1")).
		WithArgs("%search%", "%search%", "%search%", userID).
		WillReturnError(errors.New("database error"))

	itemResults, err := itemRepository.GetByQueryFilters(queryFilter, &pageFilter)

	assert.Error(t, err)
	assert.Nil(t, itemResults)
	assert.ErrorIs(t, err, repositories.ErrItemRepositoryCanNotGetByQueryFilters)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryCountByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	userID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.OrLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "sku",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    "name",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    "description",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
				},
			},
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}

	count := int64(10)
	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(count)
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `items` inner join item_keywords on item_keywords.item_id = items.id WHERE (sku LIKE ? OR name LIKE ? OR description LIKE ?) AND user_id = ?")).
		WithArgs("%search%", "%search%", "%search%", userID).
		WillReturnRows(rows)

	countResult, err := itemRepository.CountByQueryFilters(queryFilter)

	assert.NoError(t, err)
	assert.Equal(t, count, countResult)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryCountByQueryFiltersErrorCanNotCountByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	userID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.OrLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "sku",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    "name",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    "description",
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
				},
			},
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `items` inner join item_keywords on item_keywords.item_id = items.id WHERE (sku LIKE ? OR name LIKE ? OR description LIKE ?) AND user_id = ?")).
		WithArgs("%search%", "%search%", "%search%", userID).
		WillReturnError(errors.New("database error"))

	countResult, err := itemRepository.CountByQueryFilters(queryFilter)

	assert.Error(t, err)
	assert.Zero(t, countResult)
	assert.ErrorIs(t, err, repositories.ErrItemRepositoryCanNotCountByQueryFilters)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryUpdate(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: nil,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("UPDATE `items` SET `sku`=?,`name`=?,`description`=?,`unit`=?,`user_id`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(item.Sku, item.Name, item.Description, item.Unit, item.UserID, item.CreatedAt, sqlmock.AnyArg(), item.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := itemRepository.Update(item)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestItemRepositoryUpdateErrorCanNotUpdateItem(t *testing.T) {
	db, dbMock := makeDBMock()
	itemRepository := NewItemRepository(db)

	item := &entities.Item{
		ID:          uuid.NewString(),
		Sku:         random.String(20, random.Alphanumeric),
		Name:        random.String(100, random.Alphanumeric),
		Description: nil,
		Unit:        "unit",
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("UPDATE `items` SET `sku`=?,`name`=?,`description`=?,`unit`=?,`user_id`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(item.Sku, item.Name, item.Description, item.Unit, item.UserID, item.CreatedAt, sqlmock.AnyArg(), item.ID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := itemRepository.Update(item)

	assert.Error(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
