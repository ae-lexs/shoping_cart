package entity

import "errors"

var (
	// InvalidJSONError represents a Unmarshal JSON parsing error,
	InvalidJSONError = errors.New("InvalidJSON")

	// AlbumRespositoryCreateError represents a database exception.
	AlbumRespositoryCreateError = errors.New("AlbumRespositoryCreate")

	// JSONParsingError represents a Marshal JSON parsing error.
	InvalidResponseStructError = errors.New("InvalidResponseStructError")
)
