package data

import (
	"encoding/json"
	"fmt"
	"github.com/cupakob/covid19trends-rest-api/domain"
	"net/http"
)

type Fetcher interface {
	FetchAndPrepareData() (*domain.Response, error)
}

type Fetch struct {
	HTTPClient CovidHTTPClient
	URL string
}

func NewFetcher(httpClient CovidHTTPClient, url string) Fetcher {
	return &Fetch{
		HTTPClient: httpClient,
		URL: url,
	}
}

func (f *Fetch) FetchAndPrepareData() (*domain.Response, error) {
	var request *http.Request
	var err error
	if request, err = http.NewRequest(http.MethodGet, f.URL, nil); err != nil {
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
