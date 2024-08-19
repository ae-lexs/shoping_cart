package adapter

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ae-lexs/vinyl_store/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func TestCreateAlbumDBError(t *testing.T) {
	expectedAlbum := entity.Album{
		Title:    "ANY_TITLE",
		Artist:   "ANY_ARTIST",
		Price:    20.0,
		Quantity: 10,
	}
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	album_repository := NewAlbumRepository(db)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"albums\" (.+) VALUES (.+)").WillReturnResult(
		sqlmock.NewErrorResult(errors.New("ANY_DATABASE_ERROR")),
	)
	mock.ExpectCommit()

	_, err := album_repository.CreateAlbum(
		expectedAlbum.Title,
		expectedAlbum.Artist,
		expectedAlbum.Price,
		expectedAlbum.Quantity,
	)

	if err != AlbumRespositoryCreateError {
		t.Errorf("Expected error %v, but got %v", AlbumRespositoryCreateError, err)
	}
}

func TestCreateAlbumSuccessfully(t *testing.T) {
	expectedAlbum := entity.Album{
		Title:    "ANY_TITLE",
		Artist:   "ANY_ARTIST",
		Price:    20.0,
		Quantity: 10,
	}
	sqlDB, db, mock := DbMock(t)

	defer sqlDB.Close()

	album_repository := NewAlbumRepository(db)

	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	expectedSQL := "INSERT INTO \"albums\" (.+) VALUES (.+)"

	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
	mock.ExpectCommit()

	_, err := album_repository.CreateAlbum(
		expectedAlbum.Title,
		expectedAlbum.Artist,
		expectedAlbum.Price,
		expectedAlbum.Quantity,
	)

	if err != nil {
		t.Errorf("Expected error %v, but got %v", nil, err)
	}
}
