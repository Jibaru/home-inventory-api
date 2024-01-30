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
	assert.ErrorIs(t, err, repositories.ErrCanNotCreateBox)
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
