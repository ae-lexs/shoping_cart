package service

import (
	"encoding/json"
	"log"

	"github.com/ae-lexs/vinyl_store/adapter"
	"github.com/ae-lexs/vinyl_store/entity"
)

type CreateAlbumInterface interface {
	CreateAlbum(data string) (string, error)
}

type AlbumData struct {
	Title    string  `json:"title"`
	Artist   string  `json:"artist"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Response struct {
	ID    uint   `json:"id"`
	Error string `json:"error"`
}

type CreateAlbum struct {
	repository adapter.AlbumRespositoryInterface
	logger     *log.Logger
}

func NewCreateAlbum(repository adapter.AlbumRespositoryInterface, logger *log.Logger) CreateAlbum {
	return CreateAlbum{
		repository: repository,
		logger:     logger,
	}
}

func (service *CreateAlbum) CreateAlbum(data string) (string, error) {
	var albumData = AlbumData{}

	if err := json.Unmarshal([]byte(data), &albumData); err != nil {
		service.logger.Printf("InvalidJSONError: %s", err)

		return service.parseResponse(0, entity.InvalidJSONError)
	}

	createdAlbum, err := service.repository.CreateAlbum(
		albumData.Title,
		albumData.Artist,
		albumData.Price,
		albumData.Quantity,
	)

	if err != nil {
		return service.parseResponse(0, entity.InvalidJSONError)
	}

	return service.parseResponse(createdAlbum.ID, nil)
}

func (service *CreateAlbum) parseResponse(albumID uint, err error) (string, error) {
	response, responseError := json.Marshal(&Response{
		ID:    albumID,
		Error: err.Error(),
	})

	if responseError != nil {
		service.logger.Printf("InvalidResponseStructError: %s", responseError)

		return "", entity.InvalidResponseStructError
	}

	return string(response), nil
}
