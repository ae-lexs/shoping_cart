package entity

import "errors"

var (
	JSONUnmarshalError        = errors.New("JSONUnmarshalError")
	AWSLoadDefaultConfigError = errors.New("AWSLoadDefaultConfigError")
	DynamoDBMarshalMapError   = errors.New("DynamoDBMarshalMapError")
	DynamoDBPutItemError      = errors.New("DynamoDBPutItemError")
	VinylsTableAdapterError   = errors.New("VinylsTableAdapterError")
	JSONMarshalError          = errors.New("JSONMarshalError")
)
