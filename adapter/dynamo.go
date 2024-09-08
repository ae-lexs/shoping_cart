package adapter

import (
	"context"
	"log"

	"github.com/ae-lexs/vinyl_store/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const AWSRegion string = "us-east-2"
const VinylsTableName string = "vinyls"

func NewDynamoDBClient() (*dynamodb.Client, error) {
	config, err := config.LoadDefaultConfig(context.TODO(), func(options *config.LoadOptions) error {
		options.Region = AWSRegion
		return nil
	})

	if err != nil {
		log.Printf("AWSLoadDefaultConfigError: %s", err.Error())

		return nil, entity.AWSLoadDefaultConfigError
	}

	return dynamodb.NewFromConfig(config), nil
}

type VinylsTableInterface interface {
	CreateVinyl(vinyl VinylItem) error
}

type VinylsDynamoTableAdapter struct {
	dynamoDBClient *dynamodb.Client
}

func NewVinylsDynamoTableAdapter(dynamoDBClient *dynamodb.Client) *VinylsDynamoTableAdapter {
	return &VinylsDynamoTableAdapter{
		dynamoDBClient: dynamoDBClient,
	}
}

type VinylItem struct {
	ID     string  `dynamodbav:"vinyl_id"`
	Title  string  `dynamodbav:"title"`
	Artist string  `dynamodbav:"artist"`
	Price  float32 `dynamodbav:"price"`
}

func (adapter *VinylsDynamoTableAdapter) CreateVinyl(vinyl VinylItem) error {
	data, err := attributevalue.MarshalMap(vinyl)

	if err != nil {
		log.Printf("DynamoDBMarshalMapError: %s", err.Error())

		return entity.DynamoDBMarshalMapError
	}

	_, err = adapter.dynamoDBClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(VinylsTableName),
		Item:      data,
	})

	if err != nil {
		log.Printf("DynamoDBPutItemError: %s", err.Error())

		return entity.DynamoDBPutItemError
	}

	return nil
}
