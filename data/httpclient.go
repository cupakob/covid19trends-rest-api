package data

import "net/http"

// HTTPClient is a simple interface for our HTTP Client
type CovidHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is implementation of DDFHTTPClient
type CovidClient struct {
	covidclient *http.Client
}

// Do sends the request
func (c *CovidClient) Do(req *http.Request) (*http.Response, error) {
	return c.covidclient.Do(req)
}
