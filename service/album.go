package service

import (
	"encoding/json"
	"log"

	"github.com/ae-lexs/vinyl_store/entity"
)

type AlbumInterface interface {
	Create(string) (string, error)
}

type Album struct {
	logger *log.Logger
}

func NewAlbum(logger *log.Logger) *Album {
	return &Album{
		logger: logger,
	}
}

type albumDataRquest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func (a *Album) Create(bodyRequest string) (string, error) {
	albumData := albumDataRquest{}

	if err := json.Unmarshal([]byte(bodyRequest), &albumData); err != nil {
		a.logger.Printf("JSONUnmarshalError: %v", err.Error())

		return "", entity.JSONUnmarshalError
	}

	return "", nil
}
