package githubprovider

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/clients/restclient"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/github"

	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return getDoFunc(req)
}

var (
	// mockClient MockClient
	getDoFunc func(req *http.Request) (*http.Response, error)
)

func init() {
	restclient.Client = &MockClient{}
}

func TestGetAuthorizationHeader(t *testing.T) {
	accessToken := "ABC123"
	assert.EqualValues(t, "token ABC123", getAuthorizationHeader(accessToken))
}

func TestCreateRepoInvalidRestclientResponse(t *testing.T) {
	getDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("Invalid token")
	}
	response, err := CreateRepo("invalidToken", &github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
}

func TestCreateRepoResponseWithError(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"message":"Not Found"}`)))
	getDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 400,
			Body:       r,
		}, nil
	}
	response, err := CreateRepo("validToken", &github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Not Found", err.Message)
}

func TestCreateRepoSuccess(t *testing.T) {
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	getDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	response, err := CreateRepo("validToken", &github.CreateRepoRequest{})
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, "Test Name", response.Name)
	assert.EqualValues(t, "test full name", response.FullName)
	assert.EqualValues(t, "octocat", response.Owner.Login)
}
