package service

import (
	"log"
	"testing"

	"github.com/ae-lexs/vinyl_store/entity"
)

func TestAlbumServiceCreate(t *testing.T) {
	testCases := []struct {
		name             string
		bodyRequest      string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "CreatesAlbum",
			bodyRequest:      `{"title":"Grace","artist":"Jeff Buckley","price":99.99}`,
			expectedResponse: "",
			expectedError:    nil,
		},
		{
			name:             "JSONUnmarshalError",
			bodyRequest:      `{"title":"Grace","artist":"Jeff Buckley","price":99.99`,
			expectedResponse: "",
			expectedError:    entity.JSONUnmarshalError,
		},
	}

	albumService := NewAlbum(log.Default())

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResponse, actualError := albumService.Create(testCase.bodyRequest)

			if actualResponse != testCase.expectedResponse {
				t.Errorf("Expected %v but got %v", testCase.expectedResponse, actualResponse)
			}

			if actualError != testCase.expectedError {
				t.Errorf("Expected %v but got %v", testCase.expectedError, actualError)
			}
		})
	}
}
