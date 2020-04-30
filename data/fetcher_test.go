package data_test

import (
	"bytes"
	"errors"
	"github.com/cupakob/covid19trends-rest-api/data"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewFetcher(t *testing.T) {

}

func TestNewFetcher1(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// given
		mockHTTPClient := &MockHTTPClient{}

		// when
		fetcher := data.NewFetcher(mockHTTPClient, "url")

		// then
		assert.NotNil(t, fetcher)
	})
}

func TestFetch_Import(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
	    // given
		fileContent, _ := ioutil.ReadFile("testdata/validResponse.json")
		mockHTTPClient := &MockHTTPClient{
			callDo: func(req *http.Request) (*http.Response, error) {
				reader := bytes.NewReader(fileContent)
				nopCloser := ioutil.NopCloser(reader)
				return &http.Response{
					StatusCode: 200,
					Body:       nopCloser,
				}, nil
			},
		}
	    fetcher := data.Fetch{
	    	HTTPClient: mockHTTPClient,
		}

	    // when
		output, err := fetcher.FetchAndPrepareData()

		// then
		assert.NoError(t, err)
		assert.NotNil(t, output)
	})

	t.Run("sending request returns wrong status", func(t *testing.T) {
		// given
		mockHTTPClient := &MockHTTPClient{
			callDo: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 500,
				}, nil
			},
		}
		fetcher := data.Fetch{
			HTTPClient: mockHTTPClient,
		}

		// when
		output, err := fetcher.FetchAndPrepareData()

		// then
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("sending request failed", func(t *testing.T) {
		// given
		mockHTTPClient := &MockHTTPClient{
			callDo: func(req *http.Request) (*http.Response, error) {
				return &http.Response{}, errors.New("an error")
			},
		}
		fetcher := data.Fetch{
			HTTPClient: mockHTTPClient,
		}

		// when
		output, err := fetcher.FetchAndPrepareData()

		// then
		assert.Error(t, err)
		assert.Nil(t, output)
	})

}
