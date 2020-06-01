package mocks

import "net/http"

//RoundTripFunc handler
type RoundTripFunc func(req *http.Request) *http.Response

//RoundTrip function
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient test http client
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
