package entity

import "errors"

var (
	JSONUnmarshalError                = errors.New("JSONUnmarshalError")
	AWSLoadDefaultConfigError         = errors.New("AWSLoadDefaultConfigError")
	DynamoDBMarshalMapError           = errors.New("DynamoDBMarshalMapError")
	DynamoDBPutItemError              = errors.New("DynamoDBPutItemError")
	VinylsTableAdapterError           = errors.New("VinylsTableAdapterError")
	JSONMarshalError                  = errors.New("JSONMarshalError")
	DynamoDBNewBuilderExpressionError = errors.New("DynamoDBNewBuilderExpressionError")
	DynamoDBQueryError                = errors.New("DynamoDBQueryError")
	DynamoDBUnmarshalListOfMapsError  = errors.New("DynamoDBUnmarshalListOfMapsError")
	VinylIDRequiredError              = errors.New("VinylIDRequiredError")
)
