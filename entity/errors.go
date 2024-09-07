package entity

import "errors"

var (
	InvalidJSONError = errors.New("InvalidJSON")

	AlbumRespositoryCreateError = errors.New("AlbumRespositoryCreate")

	InvalidResponseStructError = errors.New("InvalidResponseStructError")
)
