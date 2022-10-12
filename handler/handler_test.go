package handler_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/cupakob/covid19trends-rest-api/domain"
	"github.com/cupakob/covid19trends-rest-api/handler"
)

func TestNewHandler(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// given
		mockFetcher := &MockFetcher{}

		// when
		h := handler.NewHandler(mockFetcher)

		// then
		assert.NotNil(t, h)
	})
}

func TestHandleProcess(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// given
		mockFetcher := &MockFetcher{
			callFetchAndPrepareData: func() (*domain.Response, error) {
				return &domain.Response{
					Countries: []domain.Country{
						{
							CountryCode: "DE",
						},
					},
				}, nil
			},
		}
		h := handler.Handle{
			Fetcher: mockFetcher,
		}
		requestParams := map[string]string{
			"countrycode": "DE",
		}

		// when
		_, statusCode, err := h.Process(requestParams)

		// then
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Equal(t, 1, mockFetcher.countFetchAndPrepareData)
	})

	t.Run("country not found", func(t *testing.T) {
		// given
		mockFetcher := &MockFetcher{
			callFetchAndPrepareData: func() (*domain.Response, error) {
				return &domain.Response{
					Countries: []domain.Country{
						{
							CountryCode: "AA",
						},
					},
				}, nil
			},
		}
		h := handler.Handle{
			Fetcher: mockFetcher,
		}
		requestParams := map[string]string{
			"countrycode": "DE",
		}

		// when
		_, statusCode, err := h.Process(requestParams)

		// then
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, statusCode)
		assert.Equal(t, 1, mockFetcher.countFetchAndPrepareData)
	})

	t.Run("path parameter can't be found", func(t *testing.T) {
		// given
		mockFetcher := &MockFetcher{}
		h := handler.Handle{
			Fetcher: mockFetcher,
		}
		requestParams := map[string]string{
			"notcountrycodeparameter": "DE",
		}

		// when
		_, statusCode, err := h.Process(requestParams)

		// then
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, statusCode)
	})

	t.Run("path parameter is not valid", func(t *testing.T) {
		// given
		mockFetcher := &MockFetcher{}
		h := handler.Handle{
			Fetcher: mockFetcher,
		}
		requestParams := map[string]string{
			"countrycode": "de",
		}

		// when
		_, statusCode, err := h.Process(requestParams)

		// then
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, statusCode)
	})

	t.Run("fetching data failed", func(t *testing.T) {
		// given
		mockFetcher := &MockFetcher{
			callFetchAndPrepareData: func() (*domain.Response, error) {
				return &domain.Response{}, errors.New("an error")
			},
		}
		h := handler.Handle{
			Fetcher: mockFetcher,
		}
		requestParams := map[string]string{
			"countrycode": "DE",
		}

		// when
		_, statusCode, err := h.Process(requestParams)

		// then
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.Equal(t, 1, mockFetcher.countFetchAndPrepareData)
	})
}
