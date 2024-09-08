package service

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

func TestCreateAlbumService(t *testing.T) {
	testCases := []struct {
		name                    string
		expectedAlbum           entity.Album
		expectedInput           string
		expectedResponse        string
		expectedHandlerError    error
		expectedRepositoryError error
	}{
		// {
		// 	name: "CreateAlbumSuccessfully",
		// 	expectedAlbum: entity.Album{
		// 		Title:    "ANY_TITLE",
		// 		Artist:   "ANY_ARTIST",
		// 		Price:    10.2,
		// 		Quantity: 10,
		// 	},
		// 	expectedInput:           string(`{'title':'ANY_TITLE','artist':'ANY_ARTIST','price':10.2,'quantity':10}`),
		// 	expectedResponse:        string(`{"id":0,"error":{}}`),
		// 	expectedHandlerError:    nil,
		// 	expectedRepositoryError: nil,
		// },
		{
			name: "JSONParseError",
			expectedAlbum: entity.Album{
				Title:    "ANY_TITLE",
				Artist:   "ANY_ARTIST",
				Price:    10.2,
				Quantity: 10,
			},
			expectedInput:           string("{'title':'ANY_TITLE','artist':'ANY_ARTIST','price':10.2,'quantity':10}"),
			expectedResponse:        string(`{"id":0,"error":"InvalidJSON"}`),
			expectedHandlerError:    nil,
			expectedRepositoryError: nil,
		},
		// {
		// 	name: "RepositoryError",
		// 	expectedAlbum: entity.Album{
		// 		Title:    "ANY_TITLE",
		// 		Artist:   "ANY_ARTIST",
		// 		Price:    10.2,
		// 		Quantity: 10,
		// 	},
		// 	expectedInput:           string("{'title':'ANY_TITLE','artist':'ANY_ARTIST','price':10.2,'quantity':10}"),
		// 	expectedResponse:        string(`{"id": 0,"error": null}`),
		// 	expectedHandlerError:    nil,
		// 	expectedRepositoryError: entity.AlbumRespositoryCreateError,
		// },
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			handler := NewCreateAlbum(
				&albumRespositoryMock{
					expectedAlbum:           entity.Album{},
					expectedRepositoryError: nil,
				},
				log.Default(),
			)

			actualResponse, actualError := handler.CreateAlbum(testCase.expectedInput)

			if actualResponse != testCase.expectedResponse {
				t.Errorf("Expected response %v, but got %v", testCase.expectedResponse, actualResponse)
			}

			if actualError != testCase.expectedHandlerError {
				t.Errorf("Expected handler error %v, but got %v", testCase.expectedHandlerError, actualError)
			}
		})
	}
}