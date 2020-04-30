package router

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/cupakob/covid19trends-rest-api/domain"
	"github.com/gorilla/mux"
	"net/http"
)

// Router is an interface that provides methods to handle incoming requests
type Router interface {
	InvokeRequest(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

// Route is a struct that implements the interface Router
type Route struct {
	MuxRouter                 *mux.Router
	gorillaAdapter            *gorillamux.GorillaMuxAdapter
}

// NewRouter is a constructor to create a Router object
func NewRouter() Router {
	muxRouter := mux.NewRouter()
	gorillaAdapter := gorillamux.New(muxRouter)
	route := &Route{
		MuxRouter:      muxRouter,
		gorillaAdapter: gorillaAdapter,
	}
	route.MuxRouter.HandleFunc("/{countrycode}", route.FetchDataForGivenCountry).Methods("GET")
	return route
}

func (r *Route) FetchDataForGivenCountry(writer http.ResponseWriter, request *http.Request) {
	pathParameters := mux.Vars(request)
	inputCountryCode, err := findPathParameter(pathParameters, "countrycode")
	if err != nil {
		buildResponse(writer, http.StatusBadRequest, fmt.Sprintf("missing path parameter. error: %v", err))
		return
	}

	countryCode := domain.CountryCode(*inputCountryCode)
	valid := countryCode.Validate()
	if !valid {
		buildResponse(writer, http.StatusBadRequest, fmt.Sprintf("given countrycode '%v' is not valid", countryCode))
		return
	}
	buildResponse(writer, http.StatusOK, fmt.Sprintf("Hello Countrycode '%v'!", *inputCountryCode))
	return
}

func (r *Route) InvokeRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return r.gorillaAdapter.Proxy(request)
}

func findPathParameter(givenMap map[string]string, key string) (*string, error) {
	value, ok := givenMap[key]
	if !ok {
		return nil, fmt.Errorf("failed to find key '%v' in path parameters", key)
	}
	return &value, nil
}

func buildResponse(writer http.ResponseWriter, code int, payload string) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	_, _ = writer.Write([]byte(payload))
}
