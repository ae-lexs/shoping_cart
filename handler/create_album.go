package handler

import (
	"encoding/json"
	"log"

	"github.com/ae-lexs/vinyl_store/adapter"
	"github.com/ae-lexs/vinyl_store/entity"
)

type CreateAlbumHandlerInterface interface {
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

type CreateAlbumHandler struct {
	repository adapter.AlbumRespositoryInterface
	logger     *log.Logger
}

func NewCreateAlbumHandler(repository adapter.AlbumRespositoryInterface, logger *log.Logger) *CreateAlbumHandler {
	return &CreateAlbumHandler{
		repository: repository,
		logger:     logger,
	}
}

func (h *CreateAlbumHandler) CreateAlbum(data string) (string, error) {
	var albumData = AlbumData{}

	if err := json.Unmarshal([]byte(data), &albumData); err != nil {
		h.logger.Printf("InvalidJSONError: %s", err)

		return h.parseResponse(0, entity.InvalidJSONError)
	}

	createdAlbum, err := h.repository.CreateAlbum(
		albumData.Title,
		albumData.Artist,
		albumData.Price,
		albumData.Quantity,
	)

	if err != nil {
		return h.parseResponse(0, entity.InvalidJSONError)
	}

	return h.parseResponse(createdAlbum.ID, nil)
}

func (h *CreateAlbumHandler) parseResponse(albumID uint, err error) (string, error) {
	response, responseError := json.Marshal(&Response{
		ID:    albumID,
		Error: err.Error(),
	})

	if responseError != nil {
		h.logger.Printf("InvalidResponseStructError: %s", responseError)

		return "", entity.InvalidResponseStructError
	}

	return string(response), nil
}
