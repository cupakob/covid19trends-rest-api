package data_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cupakob/covid19trends-rest-api/data"
)

func TestCovidClientDo(t *testing.T) {
	t.Run("should send request successful", func(t *testing.T) {
		// given
		httpClient := &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		}
		covidClient := data.NewCovidClient(httpClient)
		request := &http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "http",
				Host:   "google.com",
			},
		}

		// when
		_, err := covidClient.Do(request)

		// then
		assert.NoError(t, err)
	})
}
