package domain

import (
	"regexp"
)

type CountryCode string

const (
	countryCodePattern = "^[A-Z]{2}$"
)

func (cc CountryCode) Validate() bool {
	match, err := regexp.MatchString(countryCodePattern, string(cc))
	if err != nil {
		return false
	}
	return match
}
