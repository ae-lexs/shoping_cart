package adapter

import (
	"log"

	"github.com/ae-lexs/vinyl_store/entity"
	"gorm.io/gorm"
)

type AlbumRespositoryInterface interface {
	CreateAlbum(title, artist string, price float32, quantity int) (entity.Album, error)
}

type AlbumRespository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewAlbumRepository(db *gorm.DB, logger *log.Logger) *AlbumRespository {
	return &AlbumRespository{
		db:     db,
		logger: logger,
	}
}

func (r *AlbumRespository) CreateAlbum(title, artist string, price float32, quantity int) (entity.Album, error) {
	album := entity.Album{
		Title:    title,
		Artist:   artist,
		Price:    price,
		Quantity: quantity,
	}
	result := r.db.Create(&album)

	if result.Error != nil {
		r.logger.Printf("AlbumRespositoryCreateError: %s", result.Error)

		return entity.Album{}, entity.AlbumRespositoryCreateError
	}

	return album, nil
}
