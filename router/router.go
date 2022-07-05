package router

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"

	"github.com/cupakob/covid19trends-rest-api/handler"
)

// Router is an interface that provides methods to handle incoming requests
type Router interface {
	InvokeRequest(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

// Route is a struct that implements the interface Router
type Route struct {
	MuxRouter      *mux.Router
	gorillaAdapter *gorillamux.GorillaMuxAdapter
	Handler        handler.Handler
}

// NewRouter is a constructor to create a Router object
func NewRouter(handler handler.Handler) Router {
	muxRouter := mux.NewRouter()
	gorillaAdapter := gorillamux.New(muxRouter)
	route := &Route{
		MuxRouter:      muxRouter,
		gorillaAdapter: gorillaAdapter,
		Handler:        handler,
	}
	route.MuxRouter.HandleFunc("/{countrycode}", route.FetchDataForGivenCountry).Methods("GET")
	return route
}

func (r *Route) FetchDataForGivenCountry(writer http.ResponseWriter, request *http.Request) {
	pathParameters := mux.Vars(request)
	output, statusCode, err := r.Handler.Process(pathParameters)
	if err != nil {
		buildResponse(writer, statusCode, err.Error())
		return
	}

	outputJSON, _ := json.Marshal(output)
	buildResponse(writer, http.StatusOK, string(outputJSON))
	return
}

func (r *Route) InvokeRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return r.gorillaAdapter.Proxy(request)
}

func buildResponse(writer http.ResponseWriter, code int, payload string) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	_, _ = writer.Write([]byte(payload))
}
