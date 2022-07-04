package data

import "net/http"

type Builder interface {
	BuildRequest() (*http.Request, error)
}

type RequestBuilder struct {
	Method string
	URL    string
}

func NewRequestBuilder(method string, url string) Builder {
	return &RequestBuilder{
		Method: method,
		URL:    url,
	}
}

func (r RequestBuilder) BuildRequest() (*http.Request, error) {
	var request *http.Request
	var err error
	if request, err = http.NewRequest(r.Method, r.URL, nil); err != nil {
		return nil, err
	}

	return request, nil
}
