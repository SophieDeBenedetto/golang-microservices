package mocks

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// MockHTTPResponse sets the mock do function
func MockHTTPResponse(body string, status int) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(body)))
	response := &http.Response{
		StatusCode: status,
		Body:       r,
	}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return response, nil
	}
}

// MockHTTPError sets the mock do function
func MockHTTPError(err string) {
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(err)
	}
}
