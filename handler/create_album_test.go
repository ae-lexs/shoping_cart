package handler

import (
	"log"
	"testing"

	"github.com/ae-lexs/vinyl_store/entity"
)

type albumRespositoryMock struct {
	expectedAlbum           entity.Album
	expectedRepositoryError error
}

func (r *albumRespositoryMock) CreateAlbum(title, artist string, price float32, quantity int) (entity.Album, error) {
	return r.expectedAlbum, r.expectedRepositoryError
}

func TestCreateAlbumHandler(t *testing.T) {
	testCases := []struct {
		name                    string
		expectedAlbum           entity.Album
		expectedInput           []byte
		expectedResponse        Response
		expectedRepositoryError error
	}{
		{
			name: "CreateAlbumSuccessfully",
			expectedAlbum: entity.Album{
				Title:    "ANY_TITLE",
				Artist:   "ANY_ARTIST",
				Price:    10.2,
				Quantity: 10,
			},
			expectedInput: []byte(`{
				"title":    "ANY_TITLE",
				"artist":   "ANY_ARTIST",
				"price":    10.2,
				"quantity": 10
			}`),
			expectedResponse: Response{
				ID:    0,
				Error: nil,
			},
			expectedRepositoryError: nil,
		},
		{
			name: "JSONParseError",
			expectedAlbum: entity.Album{
				Title:    "ANY_TITLE",
				Artist:   "ANY_ARTIST",
				Price:    10.2,
				Quantity: 10,
			},
			expectedInput: []byte(`{
				"title":    "ANY_TITLE",
				"artist":   "ANY_ARTIST",
				"price":    10.2,
				"quantity": 10,
			}`),
			expectedResponse: Response{
				ID:    0,
				Error: entity.InvalidJSONError,
			},
			expectedRepositoryError: nil,
		},
		{
			name: "RepositoryError",
			expectedAlbum: entity.Album{
				Title:    "ANY_TITLE",
				Artist:   "ANY_ARTIST",
				Price:    10.2,
				Quantity: 10,
			},
			expectedInput: []byte(`{
				"title":    "ANY_TITLE",
				"artist":   "ANY_ARTIST",
				"price":    10.2,
				"quantity": 10
			}`),
			expectedResponse: Response{
				ID:    0,
				Error: nil,
			},
			expectedRepositoryError: entity.AlbumRespositoryCreateError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			handler := NewCreateAlbumHandler(
				&albumRespositoryMock{
					expectedAlbum:           entity.Album{},
					expectedRepositoryError: nil,
				},
				log.Default(),
			)

			actual_response := handler.CreateAlbum(testCase.expectedInput)

			if actual_response != testCase.expectedResponse {
				t.Errorf("Expected output %v, but got %v", testCase.expectedResponse, actual_response)
			}
		})
	}
}
