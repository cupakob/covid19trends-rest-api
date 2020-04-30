package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cupakob/covid19trends-rest-api/data"
	requestHandler "github.com/cupakob/covid19trends-rest-api/handler"
	"github.com/cupakob/covid19trends-rest-api/router"
	"log"
	"net/http"
)

var fetcher data.Fetcher
var handler requestHandler.Handler

func init() {
	httpClient := &http.Client{}
	fetcher = data.NewFetcher(httpClient, "https://api.covid19api.com/summary")
	handler = requestHandler.NewHandler(fetcher)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return router.NewRouter(handler).InvokeRequest(request)
}

func main() {
	log.Printf("Start lambda")
	lambda.Start(Handler)
}
