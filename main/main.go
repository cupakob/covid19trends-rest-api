package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cupakob/covid19trends-rest-api/router"
	"log"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return router.NewRouter().InvokeRequest(request)
}

func main() {
	log.Printf("Start lambda")
	lambda.Start(Handler)
}
