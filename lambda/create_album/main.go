package main

import (
	"log"

	"github.com/ae-lexs/vinyl_store/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Lambda struct {
	logger  *log.Logger
	service service.AlbumInterface
}

// NewLambda returns a Lambda instance.
func NewLambda() *Lambda {
	logger := log.Default()

	return &Lambda{
		logger:  logger,
		service: service.NewAlbum(),
	}
}

func (l *Lambda) handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, err := l.service.Create(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "InternalServerError",
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
