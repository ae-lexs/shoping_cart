package service

import "encoding/json"

type AlbumInterface interface {
	Create(string) (string, error)
}

type Album struct{}

func NewAlbum() *Album {
	return &Album{}
}

type albumDataRquest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func (a *Album) Create(bodyRequest string) (string, error) {
	albumData := albumDataRquest{}

	if err := json.Unmarshal([]byte(bodyRequest), &albumData); err != nil {
		return "", err
	}

	return "", nil
}
