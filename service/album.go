package service

import (
	"encoding/json"
	"log"

	"github.com/ae-lexs/vinyl_store/adapter"
	"github.com/ae-lexs/vinyl_store/entity"
	"github.com/google/uuid"
)

type AlbumInterface interface {
	Create(string) (string, error)
}

type Album struct {
	vinylsTableAdapter adapter.VinylsTableInterface
}

func NewAlbum(vinylsTableAdapter adapter.VinylsTableInterface) *Album {
	return &Album{
		vinylsTableAdapter: vinylsTableAdapter,
	}
}

type albumDataRquest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

type response struct {
	ID string `json:"id"`
}

func (service *Album) Create(bodyRequest string) (string, error) {
	albumData := albumDataRquest{}

	if err := json.Unmarshal([]byte(bodyRequest), &albumData); err != nil {
		log.Printf("JSONUnmarshalError: %s", err.Error())

		return "", entity.JSONUnmarshalError
	}

	vinylID := uuid.NewString()
	vinylItem := adapter.VinylItem{
		ID:     vinylID,
		Title:  albumData.Title,
		Artist: albumData.Artist,
		Price:  albumData.Price,
	}

	if err := service.vinylsTableAdapter.CreateVinyl(vinylItem); err != nil {
		log.Printf("VinylsTableAdapterError: %v", err.Error())

		return "", entity.VinylsTableAdapterError
	}

	response := response{
		ID: vinylID,
	}
	jsonResponse, err := json.Marshal(&response)

	if err != nil {
		log.Printf("JSONMarshalError: %s", err.Error())

		return "", entity.JSONMarshalError
	}

	return string(jsonResponse), nil
}
