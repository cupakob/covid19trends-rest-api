package domain_test

import (
	"github.com/cupakob/covid19trends-rest-api/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountryCodeValidate(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
	    // given
	    countryCode := domain.CountryCode("DE")

	    // when
		validationResult := countryCode.Validate()

		// then
	    assert.True(t, validationResult)
	})

	t.Run("validation result is false", func(t *testing.T) {
		// given
		countryCode := domain.CountryCode("de")

		// when
		validationResult := countryCode.Validate()

		// then
		assert.False(t, validationResult)
	})
}
