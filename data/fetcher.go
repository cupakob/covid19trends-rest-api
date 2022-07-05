package data

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/cupakob/covid19trends-rest-api/domain"
)

type Fetcher interface {
	FetchAndPrepareData() (*domain.Response, error)
}

type Fetch struct {
	HTTPClient CovidHTTPClient
	RequestBuilder Builder
}

func NewFetcher(httpClient CovidHTTPClient, requestBuilder Builder) Fetcher {
	return &Fetch{
		HTTPClient: httpClient,
		RequestBuilder: requestBuilder,
	}
}

func (f *Fetch) FetchAndPrepareData() (*domain.Response, error) {
	request, err := f.RequestBuilder.BuildRequest()
	if err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = f.HTTPClient.Do(request); err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response is not as expected, statusCode '%v'", response.StatusCode)
	}
	bodyReader := response.Body
	defer func() {
		_ = bodyReader.Close()
	}()

	responseJSON := domain.Response{}
	if err := json.NewDecoder(bodyReader).Decode(&responseJSON); err != nil {
		return nil, err
	}

	return &responseJSON, nil
}
