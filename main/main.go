package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cupakob/covid19trends-rest-api/data"
	requestHandler "github.com/cupakob/covid19trends-rest-api/handler"
	"github.com/cupakob/covid19trends-rest-api/resources"
	"github.com/cupakob/covid19trends-rest-api/router"
	"log"
	"net/http"
	"os"
)

var fetcher data.Fetcher
var handler requestHandler.Handler
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func init() {
	httpClient := &http.Client{}
	resources, err := resources.NewResources()
	if err != nil {
			errorLogger.Fatalf("failed to initialize awsResources. %v", err)
	}
	fetcher = data.NewFetcher(httpClient, resources.URL)
	handler = requestHandler.NewHandler(fetcher)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return router.NewRouter(handler).InvokeRequest(request)
}

func main() {
	log.Printf("Start lambda")
	lambda.Start(Handler)
}
