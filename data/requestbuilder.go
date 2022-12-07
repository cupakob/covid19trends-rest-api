package data

import "net/http"

// Builder is interface with a method to build the request
type Builder interface {
	BuildRequest() (*http.Request, error)
}

// RequestBuilder implements Builder
type RequestBuilder struct {
	Method string
	URL    string
}

// NewRequestBuilder returns an instance of the request builder
func NewRequestBuilder(method string, url string) Builder {
	return &RequestBuilder{
		Method: method,
		URL:    url,
	}
}

// BuildRequest builds the request
func (r RequestBuilder) BuildRequest() (*http.Request, error) {
	var request *http.Request
	var err error
	if request, err = http.NewRequest(r.Method, r.URL, nil); err != nil {
		return nil, err
	}

	return request, nil
}
