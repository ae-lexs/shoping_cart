package adapter

import (
	"errors"
	"log"

	"github.com/ae-lexs/vinyl_store/entity"
	"gorm.io/gorm"
)

// AlbumRespositoryCreateError represents a database exception.
var AlbumRespositoryCreateError = errors.New("AlbumRespositoryCreateError")

// Respository represents the adapter for the Album table in PostgreSQL.
type AlbumRespository struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewAlbumRepository returns an instance of AlbumRespository.
func NewAlbumRepository(db *gorm.DB, logger *log.Logger) *AlbumRespository {
	return &AlbumRespository{
		db:     db,
		logger: logger,
	}
}

// Creates an Album in the database.
func (r *AlbumRespository) CreateAlbum(title, artist string, price float64, quantity int) (entity.Album, error) {
	album := entity.Album{
		Title:    title,
		Artist:   artist,
		Price:    price,
		Quantity: quantity,
	}
	result := r.db.Create(&album)

	if result.Error != nil {
		r.logger.Printf("AlbumRespositoryCreateError: %s", result.Error)

		return entity.Album{}, AlbumRespositoryCreateError
	}

	return album, nil
}
