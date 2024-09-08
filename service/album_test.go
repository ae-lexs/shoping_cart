package service

import (
	"errors"
	"testing"

	"github.com/ae-lexs/vinyl_store/adapter"
	"github.com/ae-lexs/vinyl_store/entity"
)

type vinylsTableAdapterMock struct {
	dynamoError error
}

func (adapter *vinylsTableAdapterMock) CreateVinyl(vinyl adapter.VinylItem) error {
	return adapter.dynamoError
}

func TestAlbumServiceCreate(t *testing.T) {
	testCases := []struct {
		name                 string
		bodyRequest          string
		expectedAdapterError error
		expectedError        error
	}{
		{
			name:                 "CreatesAlbum",
			bodyRequest:          `{"title":"Grace","artist":"Jeff Buckley","price":99.99}`,
			expectedAdapterError: nil,
			expectedError:        nil,
		},
		{
			name:                 "JSONUnmarshalError",
			bodyRequest:          `{"title":"Grace","artist":"Jeff Buckley","price":99.99`,
			expectedAdapterError: nil,
			expectedError:        entity.JSONUnmarshalError,
		},
		{
			name:                 "VinylsTableAdapterError",
			bodyRequest:          `{"title":"Grace","artist":"Jeff Buckley","price":99.99}`,
			expectedAdapterError: errors.New("ANY_DYNAMO_ERROR"),
			expectedError:        entity.VinylsTableAdapterError,
		},
	}

	for _, testCase := range testCases {
		albumService := NewAlbum(&vinylsTableAdapterMock{
			dynamoError: testCase.expectedAdapterError,
		})

		t.Run(testCase.name, func(t *testing.T) {

			_, actualError := albumService.Create(testCase.bodyRequest)

			if actualError != testCase.expectedError {
				t.Errorf("Expected %v but got %v", testCase.expectedError, actualError)
			}
		})
	}
}
