package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cupakob/covid19trends-rest-api/domain"
)

func TestCountryCodeValidate(t *testing.T) {
	t.Run("validation result should be true", func(t *testing.T) {
		// given
		countryCode := domain.CountryCode("DE")

		// when
		validationResult := countryCode.Validate()

		// then
		assert.True(t, validationResult)
	})

	t.Run("validation result should be false", func(t *testing.T) {
		// given
		countryCode := domain.CountryCode("de")

		// when
		validationResult := countryCode.Validate()

		// then
		assert.False(t, validationResult)
	})
}
