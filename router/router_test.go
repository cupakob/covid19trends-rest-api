package router_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cupakob/covid19trends-rest-api/domain"
	"github.com/cupakob/covid19trends-rest-api/handler"
	"github.com/cupakob/covid19trends-rest-api/router"
)

func TestNewRouter(t *testing.T) {
	t.Run("should create new router", func(t *testing.T) {
		// given
		var h handler.Handler

		// when
		newRouter := router.NewRouter(h)

		// then
		assert.NotNil(t, newRouter)
	})
}

type MockHandler struct {
	handler.Handler
	countProcess int
	callProcess  func(requestParams map[string]string) (*domain.Output, int, error)
}

func (m *MockHandler) Process(requestParams map[string]string) (*domain.Output, int, error) {
	m.countProcess++
	return m.callProcess(requestParams)
}

func TestRouteFetchDataForGivenCountry(t *testing.T) {
	t.Run("should fetch data for the given country successfully", func(t *testing.T) {
		// given
		h := &MockHandler{
			callProcess: func(requestParams map[string]string) (*domain.Output, int, error) {
				return nil, 200, nil
			},
		}
		r := router.Route{
			Handler: h,
		}
		reqest := &http.Request{
			Method: "GET",
			Header: map[string][]string{
				"test": {},
			},
		}
		writer := httptest.NewRecorder()

		// when
		r.FetchDataForGivenCountry(writer, reqest)

		// then
		assert.Equal(t, 1, h.countProcess)
		assert.Equal(t, 200, writer.Code)
	})

	t.Run("should fails on fetch data for the given country", func(t *testing.T) {
		// given
		h := &MockHandler{
			callProcess: func(requestParams map[string]string) (*domain.Output, int, error) {
				return nil, 500, fmt.Errorf("failed")
			},
		}
		r := router.Route{
			Handler: h,
		}
		reqest := &http.Request{
			Method: "GET",
			Header: map[string][]string{
				"test": {},
			},
		}
		writer := httptest.NewRecorder()

		// when
		r.FetchDataForGivenCountry(writer, reqest)

		// then
		assert.Equal(t, 1, h.countProcess)
		assert.Equal(t, 500, writer.Code)
	})
}
