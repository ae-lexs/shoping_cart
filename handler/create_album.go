package handler

import (
	"encoding/json"
	"log"

	"github.com/ae-lexs/vinyl_store/adapter"
	"github.com/ae-lexs/vinyl_store/entity"
)

// AlbumData represents the received album information.
type AlbumData struct {
	Title    string  `json:"title"`
	Artist   string  `json:"artist"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Response represents the handler response.
type Response struct {
	ID    uint  `json:"id"`
	Error error `json:"error"`
}

// CreateAlbumHandler represents the CreateAlbum lambda handler.
type CreateAlbumHandler struct {
	repository adapter.AlbumRespositoryInterface
	logger     *log.Logger
}

// NewCreateAlbumHandler returns a new instance of CreateAlbumHandler.
func NewCreateAlbumHandler(repository adapter.AlbumRespositoryInterface, logger *log.Logger) *CreateAlbumHandler {
	return &CreateAlbumHandler{
		repository: repository,
		logger:     logger,
	}
}

// CreateAlbum receives the album information and creates it in the database.
func (h *CreateAlbumHandler) CreateAlbum(data []byte) Response {
	var albumData = AlbumData{}

	if err := json.Unmarshal(data, &albumData); err != nil {
		h.logger.Printf("InvalidJSONError: %s", err)

		return Response{
			ID:    0,
			Error: entity.InvalidJSONError,
		}
	}

	createdAlbum, err := h.repository.CreateAlbum(
		albumData.Title,
		albumData.Artist,
		albumData.Price,
		albumData.Quantity,
	)

	if err != nil {
		return Response{
			ID:    0,
			Error: err,
		}
	}

	return Response{
		ID:    createdAlbum.ID,
		Error: nil,
	}
}
