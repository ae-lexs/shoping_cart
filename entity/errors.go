package entity

import "errors"

var (
	// InvalidJSONError represents a JSON parsing error
	InvalidJSONError = errors.New("InvalidJSON")

	// AlbumRespositoryCreateError represents a database exception.
	AlbumRespositoryCreateError = errors.New("AlbumRespositoryCreate")
)
