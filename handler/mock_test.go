package handler_test

import (
	"github.com/cupakob/covid19trends-rest-api/data"
	"github.com/cupakob/covid19trends-rest-api/domain"
)

type MockFetcher struct {
	data.Fetcher
	countFetchAndPrepareData int
	callFetchAndPrepareData  func() (*domain.Response, error)
}

func (m *MockFetcher) FetchAndPrepareData() (*domain.Response, error) {
	m.countFetchAndPrepareData++
	return m.callFetchAndPrepareData()
}
