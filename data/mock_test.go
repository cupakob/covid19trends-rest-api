package data_test

import (
	"net/http"

	"github.com/cupakob/covid19trends-rest-api/data"
)

type MockCovidHTTPClient struct {
	data.CovidHTTPClient
	countDo int
	callDo  func(req *http.Request) (*http.Response, error)
}

func (m *MockCovidHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.countDo++
	return m.callDo(req)
}

type MockRequestBuilder struct {
	data.Builder
	countBuildRequest int
	callBuildRequest func() (*http.Request, error)
}

func (m *MockRequestBuilder) BuildRequest() (*http.Request, error) {
	m.countBuildRequest++
	return m.callBuildRequest()
}
