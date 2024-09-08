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

func (a *vinylsTableAdapterMock) Create(vinyl adapter.VinylItem) error {
	return a.dynamoError
}

func (a *vinylsTableAdapterMock) Get(vinylID string) (adapter.VinylItem, error) {
	return adapter.VinylItem{}, a.dynamoError
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
		albumService := NewVinyl(&vinylsTableAdapterMock{
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

func TestAlbumServiceGet(t *testing.T) {
	testCases := []struct {
		name                 string
		vinylID              string
		expectedAdapterError error
		expectedError        error
	}{
		{
			name:                 "GetsAlbum",
			vinylID:              "4a7f6d57-c324-4854-bf0a-f77926fa5e6c",
			expectedAdapterError: nil,
			expectedError:        nil,
		},
		{
			name:                 "VinylsTableAdapterError",
			vinylID:              "4a7f6d57-c324-4854-bf0a-f77926fa5e6c",
			expectedAdapterError: errors.New("ANY_DYNAMO_ERROR"),
			expectedError:        entity.VinylsTableAdapterError,
		},
	}

	for _, testCase := range testCases {
		albumService := NewVinyl(&vinylsTableAdapterMock{
			dynamoError: testCase.expectedAdapterError,
		})

		t.Run(testCase.name, func(t *testing.T) {

			_, actualError := albumService.Get(testCase.vinylID)

			if actualError != testCase.expectedError {
				t.Errorf("Expected %v but got %v", testCase.expectedError, actualError)
			}
		})
	}
}
