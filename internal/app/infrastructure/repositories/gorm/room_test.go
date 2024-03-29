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

func TestRoomRepositoryCreateRoom(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	description := random.String(255, random.Alphanumeric)
	room := &entities.Room{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `rooms` (`id`,`name`,`description`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(room.ID, room.Name, *room.Description, room.UserID, room.CreatedAt, room.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := roomRepository.Create(room)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryCreateRoomErrorCanNotCreateRoom(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	description := random.String(255, random.Alphanumeric)
	room := &entities.Room{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: &description,
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `rooms` (`id`,`name`,`description`,`user_id`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs(room.ID, room.Name, *room.Description, room.UserID, room.CreatedAt, room.UpdatedAt).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := roomRepository.Create(room)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrRoomRepositoryCanNotCreateRoom)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryExistsByID(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	roomID := uuid.NewString()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `rooms` WHERE id = ?")).
		WithArgs(roomID).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	exists, err := roomRepository.ExistsByID(roomID)

	assert.NoError(t, err)
	assert.True(t, exists)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryExistsByIDErrorCanNotCheckIfRoomExistsByID(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	roomID := uuid.NewString()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `rooms` WHERE id = ?")).
		WithArgs(roomID).
		WillReturnError(errors.New("database error"))

	exists, err := roomRepository.ExistsByID(roomID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrRoomRepositoryCanNotCheckIfRoomExistsByID)
	assert.False(t, exists)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryGetByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	roomID := uuid.NewString()
	roomName := random.String(100, random.Alphanumeric)
	roomDescription := random.String(255, random.Alphanumeric)
	userID := uuid.NewString()
	createdAt := time.Now()
	updatedAt := time.Now()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `rooms` WHERE (name LIKE ? OR description LIKE ?) AND user_id = ? LIMIT 1 OFFSET 1")).
		WithArgs("%search%", "%search%", userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "user_id", "created_at", "updated_at"}).
			AddRow(roomID, roomName, roomDescription, userID, createdAt, updatedAt))

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.OrLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    entities.RoomNameField,
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
					{
						Field:    entities.RoomDescriptionField,
						Operator: repositories.LikeComparisonOperator,
						Value:    "%search%",
					},
				},
			},
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    entities.RoomUserIDField,
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}
	pageFilter := &repositories.PageFilter{
		Offset: 1,
		Limit:  1,
	}
	rooms, err := roomRepository.GetByQueryFilters(queryFilter, pageFilter)

	assert.NoError(t, err)
	assert.Len(t, rooms, 1)
	assert.Equal(t, roomID, rooms[0].ID)
	assert.Equal(t, roomName, rooms[0].Name)
	assert.Equal(t, roomDescription, *rooms[0].Description)
	assert.Equal(t, userID, rooms[0].UserID)
	assert.Equal(t, createdAt, rooms[0].CreatedAt)
	assert.Equal(t, updatedAt, rooms[0].UpdatedAt)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryGetByQueryFiltersErrorCanNotGetRooms(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	userID := uuid.NewString()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `rooms` WHERE user_id = ? LIMIT 1 OFFSET 1")).
		WithArgs(userID).
		WillReturnError(errors.New("database error"))

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    entities.RoomUserIDField,
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}
	pageFilter := &repositories.PageFilter{
		Offset: 1,
		Limit:  1,
	}
	rooms, err := roomRepository.GetByQueryFilters(queryFilter, pageFilter)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrRoomRepositoryCanNotGetRooms)
	assert.Nil(t, rooms)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryCountByQueryFilters(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	userID := uuid.NewString()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `rooms` WHERE user_id = ?")).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1))

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    entities.RoomUserIDField,
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}
	count, err := roomRepository.CountByQueryFilters(queryFilter)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryCountByQueryFiltersErrorCanNotCountRooms(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	userID := uuid.NewString()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `rooms` WHERE user_id = ?")).
		WithArgs(userID).
		WillReturnError(errors.New("database error"))

	queryFilter := repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    entities.RoomUserIDField,
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}
	count, err := roomRepository.CountByQueryFilters(queryFilter)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrRoomRepositoryCanNotCountRooms)
	assert.Equal(t, int64(0), count)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryDelete(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	roomID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `rooms` WHERE id = ?")).
		WithArgs(roomID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := roomRepository.Delete(roomID)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryDeleteErrorCanNotDeleteRoom(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	roomID := uuid.NewString()

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `rooms` WHERE id = ?")).
		WithArgs(roomID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := roomRepository.Delete(roomID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrRoomRepositoryCanNotDeleteRoom)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryGetByID(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	roomID := uuid.NewString()
	roomName := random.String(100, random.Alphanumeric)
	roomDescription := random.String(255, random.Alphanumeric)
	userID := uuid.NewString()
	createdAt := time.Now()
	updatedAt := time.Now()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `rooms` WHERE id = ?")).
		WithArgs(roomID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "user_id", "created_at", "updated_at"}).
			AddRow(roomID, roomName, roomDescription, userID, createdAt, updatedAt))

	room, err := roomRepository.GetByID(roomID)

	assert.NoError(t, err)
	assert.Equal(t, roomID, room.ID)
	assert.Equal(t, roomName, room.Name)
	assert.Equal(t, roomDescription, *room.Description)
	assert.Equal(t, userID, room.UserID)
	assert.Equal(t, createdAt, room.CreatedAt)
	assert.Equal(t, updatedAt, room.UpdatedAt)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryGetByIDErrorCanNotGetRoomByID(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	roomID := uuid.NewString()

	dbMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `rooms` WHERE id = ?")).
		WithArgs(roomID).
		WillReturnError(errors.New("database error"))

	room, err := roomRepository.GetByID(roomID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrRoomRepositoryCanNotGetRoomByID)
	assert.Nil(t, room)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryUpdate(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	room := &entities.Room{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: nil,
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("UPDATE `rooms` SET `name`=?,`description`=?,`user_id`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(room.Name, nil, room.UserID, room.CreatedAt, sqlmock.AnyArg(), room.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	err := roomRepository.Update(room)

	assert.NoError(t, err)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRoomRepositoryUpdateErrorCanNotUpdateRoom(t *testing.T) {
	db, dbMock := makeDBMock()
	roomRepository := NewRoomRepository(db)

	room := &entities.Room{
		ID:          uuid.NewString(),
		Name:        random.String(100, random.Alphanumeric),
		Description: nil,
		UserID:      uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec(regexp.QuoteMeta("UPDATE `rooms` SET `name`=?,`description`=?,`user_id`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(room.Name, nil, room.UserID, room.CreatedAt, sqlmock.AnyArg(), room.ID).
		WillReturnError(errors.New("database error"))
	dbMock.ExpectRollback()

	err := roomRepository.Update(room)

	assert.Error(t, err)
	assert.ErrorIs(t, err, repositories.ErrRoomRepositoryCanNotUpdateRoom)
	err = dbMock.ExpectationsWereMet()
	assert.NoError(t, err)
}
