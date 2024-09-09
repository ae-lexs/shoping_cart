package main

import (
	"log"

	"github.com/ae-lexs/vinyl_store/adapter"
	"github.com/ae-lexs/vinyl_store/entity"
	"github.com/ae-lexs/vinyl_store/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Lambda struct {
	logger  *log.Logger
	service service.VinylInterface
}

func NewLambda() *Lambda {
	logger := log.Default()
	dynamoClient, err := adapter.NewDynamoDBClient()

	if err != nil {
		panic(err)
	}

	return &Lambda{
		logger: logger,
		service: service.NewVinyl(
			adapter.NewVinylsDynamoTableAdapter(dynamoClient),
		),
	}
}

func (l *Lambda) handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	vinylID := request.PathParameters["id"]

	if vinylID == "" {
		return events.APIGatewayProxyResponse{
			Body:       entity.VinylIDRequiredError.Error(),
			StatusCode: 400,
		}, nil
	}

	response, err := l.service.Get(vinylID)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       response,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(NewLambda().handler)
}
