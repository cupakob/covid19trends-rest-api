package data_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cupakob/covid19trends-rest-api/data"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewFetcher(t *testing.T) {
	t.Run("should create fetcher successfully", func(t *testing.T) {
		// given
		mockHTTPClient := &MockCovidHTTPClient{}
		requestBuilder := &MockRequestBuilder{}

		// when
		fetcher := data.NewFetcher(mockHTTPClient, requestBuilder)

		// then
		assert.NotNil(t, fetcher)
	})
}

func TestFetchImport(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
	    // given
		fileContent, _ := ioutil.ReadFile("testdata/validResponse.json")
		mockHTTPClient := &MockCovidHTTPClient{
			callDo: func(req *http.Request) (*http.Response, error) {
				reader := bytes.NewReader(fileContent)
				nopCloser := ioutil.NopCloser(reader)
				return &http.Response{
					StatusCode: 200,
					Body:       nopCloser,
				}, nil
			},
		}
		mockRequestBuilder := &MockRequestBuilder{
			callBuildRequest: func() (*http.Request, error) {
				return nil, nil
			},
		}
	    fetcher := data.Fetch{
	    	HTTPClient:     mockHTTPClient,
	    	RequestBuilder: mockRequestBuilder,
		}

	    // when
		output, err := fetcher.FetchAndPrepareData()

		// then
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, 1, mockHTTPClient.countDo)
		assert.Equal(t, 1, mockRequestBuilder.countBuildRequest)
	})

	t.Run("sending request returns wrong status", func(t *testing.T) {
		// given
		mockHTTPClient := &MockCovidHTTPClient{
			callDo: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 500,
				}, nil
			},
		}
		mockRequestBuilder := &MockRequestBuilder{
			callBuildRequest: func() (*http.Request, error) {
				return nil, nil
			},
		}
		fetcher := data.Fetch{
			HTTPClient:     mockHTTPClient,
			RequestBuilder: mockRequestBuilder,
		}


		// when
		output, err := fetcher.FetchAndPrepareData()

		// then
		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, 1, mockHTTPClient.countDo)
		assert.Equal(t, 1, mockRequestBuilder.countBuildRequest)
	})

	t.Run("sending request failed", func(t *testing.T) {
		// given
		mockHTTPClient := &MockCovidHTTPClient{
			callDo: func(req *http.Request) (*http.Response, error) {
				return &http.Response{}, errors.New("an error")
			},
		}
		mockRequestBuilder := &MockRequestBuilder{
			callBuildRequest: func() (*http.Request, error) {
				return nil, nil
			},
		}
		fetcher := data.Fetch{
			HTTPClient:     mockHTTPClient,
			RequestBuilder: mockRequestBuilder,
		}


		// when
		output, err := fetcher.FetchAndPrepareData()

		// then
		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, 1, mockHTTPClient.countDo)
		assert.Equal(t, 1, mockRequestBuilder.countBuildRequest)
	})

	t.Run("should return error when building request fails", func(t *testing.T) {
		// given
		mockHTTPClient := &MockCovidHTTPClient{}
		mockRequestBuilder := &MockRequestBuilder{
			callBuildRequest: func() (*http.Request, error) {
				return nil, fmt.Errorf("failed to build request")
			},
		}
		fetcher := data.Fetch{
			HTTPClient:     mockHTTPClient,
			RequestBuilder: mockRequestBuilder,
		}


		// when
		output, err := fetcher.FetchAndPrepareData()

		// then
		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, 1, mockRequestBuilder.countBuildRequest)
		assert.Equal(t, 0, mockHTTPClient.countDo)
	})
}
