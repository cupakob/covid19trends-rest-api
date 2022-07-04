package data

import "net/http"

// CovidHTTPClient is a simple interface for our HTTP Client
type CovidHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// CovidClient is implementation of CovidHTTPClient
type CovidClient struct {
	covidclient *http.Client
}

func NewCovidClient(covidclient *http.Client) CovidHTTPClient {
	return &CovidClient{covidclient: covidclient}
}

// Do sends the request
func (c *CovidClient) Do(req *http.Request) (*http.Response, error) {
	return c.covidclient.Do(req)
}
