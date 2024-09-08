package service

import (
	"encoding/json"
	"log"

	"github.com/ae-lexs/vinyl_store/adapter"
	"github.com/ae-lexs/vinyl_store/entity"
	"github.com/google/uuid"
)

type VinylInterface interface {
	Create(string) (string, error)
	Get(string) (string, error)
}

type Vinyl struct {
	vinylsTableAdapter adapter.VinylsTableInterface
}

func NewVinyl(vinylsTableAdapter adapter.VinylsTableInterface) *Vinyl {
	return &Vinyl{
		vinylsTableAdapter: vinylsTableAdapter,
	}
}

type albumDataRequest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

type response struct {
	ID string `json:"id"`
}

type vinylResponse struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func (service *Vinyl) Create(bodyRequest string) (string, error) {
	albumData := albumDataRequest{}

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

	if err := service.vinylsTableAdapter.Create(vinylItem); err != nil {
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

func (service *Vinyl) Get(vinylID string) (string, error) {
	vinyl, err := service.vinylsTableAdapter.Get(vinylID)

	if err != nil {
		log.Printf("VinylsTableAdapterError: %s", err.Error())

		return "", entity.VinylsTableAdapterError
	}

	response := vinylResponse{
		ID:     vinyl.ID,
		Title:  vinyl.Title,
		Artist: vinyl.Artist,
		Price:  vinyl.Price,
	}
	jsonResponse, err := json.Marshal(&response)

	if err != nil {
		log.Printf("JSONMarshalError: %s", err.Error())

		return "", entity.JSONMarshalError
	}

	return string(jsonResponse), nil
}