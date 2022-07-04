package data_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCovidClient_Do(t *testing.T) {
	t.Run("should execute request only once", func(t *testing.T) {
	    // given
		mockHTTPClient := &MockHTTPClient{
			callDo: func(req *http.Request) (*http.Response, error) {
				return nil, nil
			},
		}
		var request *http.Request

		// when
		_, err := mockHTTPClient.Do(request)

		// then
		assert.NoError(t, err)
		assert.Equal(t, 1, mockHTTPClient.countDo)
	})
}
