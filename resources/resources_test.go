package resources_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cupakob/covid19trends-rest-api/resources"
)

func TestNewResources(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// given
		givenURL := "sourceURL"
		_ = os.Setenv("source_url", givenURL)

		// when
		r, err := resources.NewResources()

		// then
		assert.NoError(t, err)
		assert.Equal(t, r.URL, givenURL)
	})

	t.Run("url is missing", func(t *testing.T) {
		// given
		_ = os.Unsetenv("source_url")

		// when
		r, err := resources.NewResources()

		// then
		assert.Error(t, err)
		assert.Nil(t, r)
	})
}
