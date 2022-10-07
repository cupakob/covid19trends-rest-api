package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"

	"github.com/cupakob/covid19trends-rest-api/data"
	requestHandler "github.com/cupakob/covid19trends-rest-api/handler"
	r "github.com/cupakob/covid19trends-rest-api/resources"
	"github.com/cupakob/covid19trends-rest-api/router"
)

var handler requestHandler.Handler
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func init() {
	resources, err := r.NewResources()
	if err != nil {
		errorLogger.Fatalf("failed to initialize awsResources. %v", err)
	}
	handler = createHandler(resources)
}

func createHandler(resources *r.Resources) requestHandler.Handler {
	fetcher := createFetcher(resources)
	return requestHandler.NewHandler(fetcher)
}

func createFetcher(resources *r.Resources) data.Fetcher {
	covidClient := createCovidClient()
	requestBuilder := data.NewRequestBuilder("GET", resources.URL)
	return data.NewFetcher(covidClient, requestBuilder)
}

func createCovidClient() data.CovidHTTPClient {
	httpClient := &http.Client{}
	covidClient := data.NewCovidClient(httpClient)
	return covidClient
}

func Handler(request core.SwitchableAPIGatewayRequest) (*core.SwitchableAPIGatewayResponse, error) {
	return router.NewRouter(handler).InvokeRequest(request)
}

func main() {
	log.Printf("Start lambda")
	lambda.Start(Handler)
}
