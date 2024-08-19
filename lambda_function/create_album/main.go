package main

import (
	"log"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	logger = log.Default()
)

type Dependencies struct {
	logger *log.Logger
}

// NewDependencies returns a Dependencies instance.
func NewDependencies() *Dependencies {
	return &Dependencies{
		logger: logger,
	}
}

func (d *Dependencies) handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	d.logger.Printf("Body: %s", request.Body)
	d.logger.Printf("Body: %s", reflect.TypeOf(request.Body))

	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(NewDependencies().handler)
}
