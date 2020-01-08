package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

// Mock mocks a rest client response
type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Error      error
}

// StartMocks sets mocking to true
func StartMocks() {
	enabledMocks = true
}

// StopMocks sets mocking to false
func StopMocks() {
	enabledMocks = false
}

// AddMock adds a mock to the mocks map
// expects a pointer
func AddMock(mock *Mock) {
	mocks[mock.URL] = mock
}

// Post sends a post request to the URL with the body
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[url]
		if mock == nil {
			return nil, errors.New("No mock for URL")
		}
		return mock.Response, mock.Error
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	client := http.Client{}
	return client.Do(request)
}
