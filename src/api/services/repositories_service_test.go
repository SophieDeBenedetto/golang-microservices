package services

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/clients/restclient"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/repositories"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func init() {
	restclient.Client = &mocks.MockClient{}
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{Name: ""}
	resp, err := RepoService.CreateRepo(request)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Invalid repo name", err.GetMessage())
}

func TestCreateRepoGithubError(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(
			"Error creating repo",
		)
	}
	request := repositories.CreateRepoRequest{Name: "Test Repo Name"}
	resp, err := RepoService.CreateRepo(request)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Error creating repo", err.GetMessage())
	assert.EqualValues(t, 500, err.GetStatus())
}

func TestCreateRepoSuccess(t *testing.T) {
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	request := repositories.CreateRepoRequest{Name: "Test Name"}
	resp, err := RepoService.CreateRepo(request)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, "Test Name", resp.Name)
	assert.EqualValues(t, "octocat", resp.Owner)
}
