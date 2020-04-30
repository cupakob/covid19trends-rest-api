package handler

import (
	"fmt"
	"github.com/cupakob/covid19trends-rest-api/data"
	"github.com/cupakob/covid19trends-rest-api/domain"
	"net/http"
	"time"
)

type Handler interface {
	Process(map[string]string) (*domain.Output, int, error)
}

type Handle struct {
	Fetcher data.Fetcher
}

func NewHandler(fetcher data.Fetcher) Handler {
	return &Handle{Fetcher: fetcher}
}

func (h *Handle) Process(requestParams map[string]string) (*domain.Output, int, error) {
	inputCountryCode, err := findPathParameter(requestParams, "countrycode")
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("missing path parameter countrycode. error: %v", err)
	}

	countryCode := domain.CountryCode(*inputCountryCode)
	valid := countryCode.Validate()
	if !valid {
		return nil, http.StatusBadRequest, fmt.Errorf("given countrycode '%v' is not valid. error: %v", countryCode, err)
	}

	responseJSON, err := h.Fetcher.FetchAndPrepareData()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("fetching data failed, error: %v", err)
	}

	for _, country := range responseJSON.Countries {
		if country.CountryCode == string(countryCode) {
			time, _ := time.Parse(time.RFC3339, country.Date)
			return &domain.Output{
				Cases:     country.TotalConfirmed,
				CasesNew:  country.NewConfirmed,
				Deaths:    country.TotalDeaths,
				DeathsNew: country.NewDeaths,
				Timestamp: time.Unix(),
				Date:      country.Date,
				DaysPast:  0,
			}, http.StatusOK, nil
		}
	}

	return nil, http.StatusNotFound, fmt.Errorf("no country found for the given countrycode '%v'", countryCode)
}

func findPathParameter(givenMap map[string]string, key string) (*string, error) {
	value, ok := givenMap[key]
	if !ok {
		return nil, fmt.Errorf("failed to find key '%v' in path parameters", key)
	}
	return &value, nil
}
