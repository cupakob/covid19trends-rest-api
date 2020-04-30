package data_test

import (
	"github.com/cupakob/covid19trends-rest-api/data"
	"net/http"
)

type MockHTTPClient struct {
	data.CovidHTTPClient
	countDo int
	callDo  func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.countDo++
	return m.callDo(req)
}

