package data_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cupakob/covid19trends-rest-api/data"
)

func TestRequestBuilder_BuildRequest(t *testing.T) {
	t.Run("should create request successful", func(t *testing.T) {
		// given
		builder := data.NewRequestBuilder("GET", "http://localhost")

		// when
		request, err := builder.BuildRequest()

		// then
		assert.NoError(t, err)
		assert.Equal(t, "http", request.URL.Scheme)
		assert.Equal(t, "localhost", request.URL.Host)
	})

	t.Run("should fails on BuildRequest when the http method is not valid", func(t *testing.T) {
		// given
		builder := data.NewRequestBuilder("B", "\n")

		// when
		_, err := builder.BuildRequest()

		// then
		assert.Error(t, err)
	})
}
