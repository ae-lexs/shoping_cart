package adapter

import (
	"context"
	"log"

	"github.com/ae-lexs/vinyl_store/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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
	Create(VinylItem) error
	Get(string) (VinylItem, error)
	GetAll() ([]VinylItem, error)
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

func (adapter *VinylsDynamoTableAdapter) Create(vinyl VinylItem) error {
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

func (adapter *VinylsDynamoTableAdapter) Get(vinylID string) (VinylItem, error) {
	var vinyls []VinylItem
	keyCondition := expression.Key("vinyl_id").Equal(expression.Value(vinylID))
	expression, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()

	if err != nil {
		log.Printf("DynamoDBNewBuilderExpressionError: %s", err.Error())

		return VinylItem{}, entity.DynamoDBNewBuilderExpressionError
	}

	data, err := adapter.dynamoDBClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String("VinylsTableName"),
		ExpressionAttributeNames:  expression.Names(),
		ExpressionAttributeValues: expression.Values(),
		KeyConditionExpression:    expression.Filter(),
	})

	if err != nil {
		log.Printf("DynamoDBQueryError: %s", err.Error())

		return VinylItem{}, entity.DynamoDBQueryError
	}

	err = attributevalue.UnmarshalListOfMaps(data.Items, &vinyls)

	if err != nil {
		log.Printf("DynamoDBUnmarshalListOfMapsError: %s", err.Error())

		return VinylItem{}, entity.DynamoDBUnmarshalListOfMapsError
	}

	log.Printf("Vinyls: %v", vinyls)

	return vinyls[0], nil
}

func (adapter *VinylsDynamoTableAdapter) GetAll() ([]VinylItem, error) {
	var vinyls []VinylItem

	data, err := adapter.dynamoDBClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(VinylsTableName),
	})

	if err != nil {
		log.Printf("DynamoDBScanError: %s", err.Error())

		return vinyls, entity.DynamoDBScanError
	}

	err = attributevalue.UnmarshalListOfMaps(data.Items, &vinyls)

	if err != nil {
		log.Printf("DynamoDBUnmarshalListOfMapsError: %s", err.Error())

		return vinyls, entity.DynamoDBUnmarshalListOfMapsError
	}

	return vinyls, nil
}
