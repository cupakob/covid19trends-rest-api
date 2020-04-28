package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

var debugLogger = log.New(os.Stdout, "DEBUG ", log.Llongfile)

func handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	debugLogger.Printf("BODY: ", request.Body)
	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	lambda.Start(handle)
}
