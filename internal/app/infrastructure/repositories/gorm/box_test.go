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

func TestBoxRepositoryCreateBox(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	description := random.String(255, random.Alphanumeric)
	box := &entities.Box{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		RoomID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `boxes` (`id`,`name`,`description`,`room_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(box.ID, box.Name, *box.Description, box.RoomID, box.CreatedAt, box.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.Create(box)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCreateBoxErrorCanNotCreateBox(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	description := random.String(255, random.Alphanumeric)
	box := &entities.Box{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		RoomID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `boxes` (`id`,`name`,`description`,`room_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(box.ID, box.Name, *box.Description, box.RoomID, box.CreatedAt, box.UpdatedAt).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.Create(box)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotCreateBox)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryGetBoxItem(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()
	itemID := uuid.NewString()
	boxItem := entities.BoxItem{
		ID:        uuid.NewString(),
		Quantity:  100.0,
		BoxID:     boxID,
		ItemID:    itemID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `box_items` WHERE box_id = ? AND item_id = ? ORDER BY `box_items`.`id` LIMIT 1")).
		WithArgs(boxID, itemID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "quantity", "box_id", "item_id", "created_at", "updated_at"}).
			AddRow(boxItem.ID, boxItem.Quantity, boxItem.BoxID, boxItem.ItemID, boxItem.CreatedAt, boxItem.UpdatedAt))

	result, err := boxRepository.GetBoxItem(boxID, itemID)

	assert.NoError(t, err)
	assert.Equal(t, &boxItem, result)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryGetBoxItemErrorBoxRepositoryBoxItemNotFound(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()
	itemID := uuid.NewString()
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `box_items` WHERE box_id = ? AND item_id = ? ORDER BY `box_items`.`id` LIMIT 1")).
		WithArgs(boxID, itemID).
		WillReturnError(errors.New("record not found"))

	result, err := boxRepository.GetBoxItem(boxID, itemID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryBoxItemNotFound)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCreateBoxItem(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxItem := entities.BoxItem{
		ID:        uuid.NewString(),
		Quantity:  100.0,
		BoxID:     uuid.NewString(),
		ItemID:    uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `box_items` (`id`,`quantity`,`box_id`,`item_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(boxItem.ID, boxItem.Quantity, boxItem.BoxID, boxItem.ItemID, boxItem.CreatedAt, boxItem.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.CreateBoxItem(&boxItem)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCreateBoxItemErrorBoxRepositoryCanBotCreateBoxItem(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxItem := entities.BoxItem{
		ID:        uuid.NewString(),
		Quantity:  100.0,
		BoxID:     uuid.NewString(),
		ItemID:    uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `box_items` (`id`,`quantity`,`box_id`,`item_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(boxItem.ID, boxItem.Quantity, boxItem.BoxID, boxItem.ItemID, boxItem.CreatedAt, boxItem.UpdatedAt).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.CreateBoxItem(&boxItem)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanBotCreateBoxItem)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryUpdateBoxItem(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxItem := entities.BoxItem{
		ID:        uuid.NewString(),
		Quantity:  100.0,
		BoxID:     uuid.NewString(),
		ItemID:    uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("UPDATE `box_items` SET `quantity`=?,`box_id`=?,`item_id`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(boxItem.Quantity, boxItem.BoxID, boxItem.ItemID, boxItem.CreatedAt, sqlmock.AnyArg(), boxItem.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.UpdateBoxItem(&boxItem)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryUpdateBoxItemErrorBoxRepositoryCanNotUpdateBoxItem(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxItem := entities.BoxItem{
		ID:        uuid.NewString(),
		Quantity:  100.0,
		BoxID:     uuid.NewString(),
		ItemID:    uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("UPDATE `box_items` SET `quantity`=?,`box_id`=?,`item_id`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(boxItem.Quantity, boxItem.BoxID, boxItem.ItemID, boxItem.CreatedAt, sqlmock.AnyArg(), boxItem.ID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.UpdateBoxItem(&boxItem)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotUpdateBoxItem)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCreateBoxTransaction(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxTransaction := entities.BoxTransaction{
		ID:         uuid.NewString(),
		Type:       entities.BoxTransactionTypeAdd,
		Quantity:   100.0,
		BoxID:      uuid.NewString(),
		ItemID:     uuid.NewString(),
		ItemSku:    random.String(4, random.Alphanumeric),
		ItemName:   random.String(10, random.Alphanumeric),
		ItemUnit:   "unit",
		HappenedAt: time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `box_transactions` (`id`,`type`,`quantity`,`box_id`,`item_id`,`item_sku`,`item_name`,`item_unit`,`happened_at`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(
			boxTransaction.ID,
			boxTransaction.Type,
			boxTransaction.Quantity,
			boxTransaction.BoxID,
			boxTransaction.ItemID,
			boxTransaction.ItemSku,
			boxTransaction.ItemName,
			boxTransaction.ItemUnit,
			boxTransaction.HappenedAt,
			boxTransaction.CreatedAt,
			boxTransaction.UpdatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.CreateBoxTransaction(&boxTransaction)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCreateBoxTransactionErrorBoxRepositoryCanNotCreateBoxTransaction(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxTransaction := entities.BoxTransaction{
		ID:         uuid.NewString(),
		Type:       entities.BoxTransactionTypeAdd,
		Quantity:   100.0,
		BoxID:      uuid.NewString(),
		ItemID:     uuid.NewString(),
		ItemSku:    random.String(4, random.Alphanumeric),
		ItemName:   random.String(10, random.Alphanumeric),
		ItemUnit:   "unit",
		HappenedAt: time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `box_transactions` (`id`,`type`,`quantity`,`box_id`,`item_id`,`item_sku`,`item_name`,`item_unit`,`happened_at`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(
			boxTransaction.ID,
			boxTransaction.Type,
			boxTransaction.Quantity,
			boxTransaction.BoxID,
			boxTransaction.ItemID,
			boxTransaction.ItemSku,
			boxTransaction.ItemName,
			boxTransaction.ItemUnit,
			boxTransaction.HappenedAt,
			boxTransaction.CreatedAt,
			boxTransaction.UpdatedAt,
		).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.CreateBoxTransaction(&boxTransaction)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotCreateBoxTransaction)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryDeleteBoxItem(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()
	itemID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `box_items` WHERE box_id = ? AND item_id = ?")).
		WithArgs(boxID, itemID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.DeleteBoxItem(boxID, itemID)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryDeleteBoxItemErrorBoxRepositoryCanNotDeleteBoxItem(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()
	itemID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `box_items` WHERE box_id = ? AND item_id = ?")).
		WithArgs(boxID, itemID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.DeleteBoxItem(boxID, itemID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotDeleteBoxItem)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryGetByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	userID := uuid.NewString()
	roomID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "rooms.user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
					{
						Field:    "room_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    roomID,
					},
				},
			},
		},
	}
	pageFilter := &repositories.PageFilter{
		Offset: 0,
		Limit:  10,
	}

	boxes := make([]*entities.Box, 0)
	for i := 0; i < 10; i++ {
		description := random.String(255, random.Alphanumeric)
		boxes = append(boxes, &entities.Box{
			ID:          uuid.NewString(),
			Name:        random.String(100, random.Alphanumeric),
			Description: &description,
			RoomID:      uuid.NewString(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	rows := sqlmock.NewRows([]string{"id", "name", "description", "room_id", "created_at", "updated_at"})
	for _, box := range boxes {
		rows.AddRow(box.ID, box.Name, *box.Description, box.RoomID, box.CreatedAt, box.UpdatedAt)
	}
	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT `boxes`.`id`,`boxes`.`name`,`boxes`.`description`,`boxes`.`room_id`,`boxes`.`created_at`,`boxes`.`updated_at` FROM `boxes` inner join rooms on boxes.room_id = rooms.id WHERE rooms.user_id = ? AND room_id = ? LIMIT 10")).
		WithArgs(userID, roomID).
		WillReturnRows(rows)

	result, err := boxRepository.GetByQueryFilters(queryFilter, pageFilter)

	assert.NoError(t, err)
	assert.Equal(t, boxes, result)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryGetByQueryFiltersErrorBoxRepositoryCanNotGetByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	userID := uuid.NewString()
	roomID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "rooms.user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
					{
						Field:    "room_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    roomID,
					},
				},
			},
		},
	}
	pageFilter := &repositories.PageFilter{
		Offset: 0,
		Limit:  10,
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT `boxes`.`id`,`boxes`.`name`,`boxes`.`description`,`boxes`.`room_id`,`boxes`.`created_at`,`boxes`.`updated_at` FROM `boxes` inner join rooms on boxes.room_id = rooms.id WHERE rooms.user_id = ? AND room_id = ? LIMIT 10")).
		WithArgs(userID, roomID).
		WillReturnError(errors.New("database error"))

	result, err := boxRepository.GetByQueryFilters(queryFilter, pageFilter)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotGetByQueryFilters)
	assert.Nil(t, result)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCountByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	userID := uuid.NewString()
	roomID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "rooms.user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
					{
						Field:    "room_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    roomID,
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `boxes` inner join rooms on boxes.room_id = rooms.id WHERE rooms.user_id = ? AND room_id = ?")).
		WithArgs(userID, roomID).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(10))

	result, err := boxRepository.CountByQueryFilters(queryFilter)

	assert.NoError(t, err)
	assert.Equal(t, int64(10), result)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCountByQueryFiltersErrorBoxRepositoryCanNotCountByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	userID := uuid.NewString()
	roomID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "rooms.user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
					{
						Field:    "room_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    roomID,
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `boxes` inner join rooms on boxes.room_id = rooms.id WHERE rooms.user_id = ? AND room_id = ?")).
		WithArgs(userID, roomID).
		WillReturnError(errors.New("database error"))

	result, err := boxRepository.CountByQueryFilters(queryFilter)

	assert.Error(t, err)
	assert.Empty(t, result)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryDeleteBoxItemsByBoxID(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `box_items` WHERE box_id = ?")).
		WithArgs(boxID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.DeleteBoxItemsByBoxID(boxID)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryDeleteBoxItemsByBoxIDErrorBoxRepositoryCanNotDeleteBoxItemsByBoxID(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `box_items` WHERE box_id = ?")).
		WithArgs(boxID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.DeleteBoxItemsByBoxID(boxID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotDeleteBoxItemsByBoxID)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryDeleteBoxTransactionsByBoxID(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `box_transactions` WHERE box_id = ?")).
		WithArgs(boxID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.DeleteBoxTransactionsByBoxID(boxID)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryDeleteBoxTransactionsByBoxIDErrorBoxRepositoryCanNotDeleteBoxTransactionsByBoxID(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `box_transactions` WHERE box_id = ?")).
		WithArgs(boxID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.DeleteBoxTransactionsByBoxID(boxID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotDeleteBoxTransactionsByBoxID)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryDelete(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `boxes` WHERE id = ?")).
		WithArgs(boxID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.Delete(boxID)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryDeleteErrorBoxRepositoryCanNotDeleteBox(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `boxes` WHERE id = ?")).
		WithArgs(boxID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.Delete(boxID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotDeleteBox)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryGetByID(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()
	description := random.String(255, random.Alphanumeric)
	box := &entities.Box{
		ID:          boxID,
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		RoomID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `boxes` WHERE id = ? ORDER BY `boxes`.`id` LIMIT 1")).
		WithArgs(boxID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "room_id", "created_at", "updated_at"}).
			AddRow(box.ID, box.Name, *box.Description, box.RoomID, box.CreatedAt, box.UpdatedAt))

	result, err := boxRepository.GetByID(boxID)

	assert.NoError(t, err)
	assert.Equal(t, box, result)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryGetByIDErrorBoxRepositoryCanNotGetByID(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `boxes` WHERE id = ? ORDER BY `boxes`.`id` LIMIT 1")).
		WithArgs(boxID).
		WillReturnError(errors.New("record not found"))

	result, err := boxRepository.GetByID(boxID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotGetByID)
	assert.Nil(t, result)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryUpdate(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	description := random.String(255, random.Alphanumeric)
	box := &entities.Box{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		RoomID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("UPDATE `boxes` SET `name`=?,`description`=?,`room_id`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(box.Name, *box.Description, box.RoomID, box.CreatedAt, sqlmock.AnyArg(), box.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := boxRepository.Update(box)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryUpdateErrorBoxRepositoryCanNotUpdate(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	description := random.String(255, random.Alphanumeric)
	box := &entities.Box{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		RoomID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("UPDATE `boxes` SET `name`=?,`description`=?,`room_id`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(box.Name, *box.Description, box.RoomID, box.CreatedAt, sqlmock.AnyArg(), box.ID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := boxRepository.Update(box)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotUpdate)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryGetBoxTransactionsByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "box_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    boxID,
					},
				},
			},
		},
	}
	pageFilter := &repositories.PageFilter{
		Offset: 0,
		Limit:  10,
	}

	boxTransactions := []*entities.BoxTransaction{
		{
			ID:       uuid.NewString(),
			Type:     "add",
			Quantity: 12,
			BoxID:    boxID,
			ItemID:   uuid.NewString(),
			ItemSku:  random.String(4, random.Alphanumeric),
			ItemName: random.String(10, random.Alphanumeric),
			ItemUnit: "unit",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "type", "quantity", "box_id", "item_id", "item_sku", "item_name", "item_unit", "happened_at", "created_at", "updated_at"})
	for _, boxTransaction := range boxTransactions {
		rows.AddRow(boxTransaction.ID, boxTransaction.Type, boxTransaction.Quantity, boxTransaction.BoxID, boxTransaction.ItemID, boxTransaction.ItemSku, boxTransaction.ItemName, boxTransaction.ItemUnit, boxTransaction.HappenedAt, boxTransaction.CreatedAt, boxTransaction.UpdatedAt)
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `box_transactions` WHERE box_id = ? ORDER BY created_at desc LIMIT 10")).
		WithArgs(boxID).
		WillReturnRows(rows)

	transactions, err := boxRepository.GetBoxTransactionsByQueryFilters(queryFilter, pageFilter)

	assert.NoError(t, err)
	assert.Equal(t, boxTransactions, transactions)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryGetBoxTransactionsByQueryFiltersErrorCanNotGetBoxTransactionsByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "box_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    boxID,
					},
				},
			},
		},
	}
	pageFilter := &repositories.PageFilter{
		Offset: 0,
		Limit:  10,
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `box_transactions` WHERE box_id = ? ORDER BY created_at desc LIMIT 10")).
		WithArgs(boxID).
		WillReturnError(errors.New("database error"))

	transactions, err := boxRepository.GetBoxTransactionsByQueryFilters(queryFilter, pageFilter)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrBoxRepositoryCanNotGetBoxTransactionsByQueryFilters)
	assert.Nil(t, transactions)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCountBoxTransactionsByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "box_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    boxID,
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `box_transactions` WHERE box_id = ?")).
		WithArgs(boxID).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(10))

	count, err := boxRepository.CountBoxTransactionsByQueryFilters(queryFilter)

	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestBoxRepositoryCountBoxTransactionsByQueryFiltersErrorCanNotCountBoxTransactionsByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	boxRepository := NewBoxRepository(db)

	boxID := uuid.NewString()

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "box_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    boxID,
					},
				},
			},
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `box_transactions` WHERE box_id = ?")).
		WithArgs(boxID).
		WillReturnError(errors.New("database error"))

	count, err := boxRepository.CountBoxTransactionsByQueryFilters(queryFilter)

	assert.Error(t, err)
	assert.Empty(t, count)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
