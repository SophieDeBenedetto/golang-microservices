package services

import (
	"bytes"
	errs "errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/clients/restclient"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/repositories"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/errors"
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
		return nil, errs.New(
			"Error creating repo",
		)
	}
	request := repositories.CreateRepoRequest{Name: "Test Repo Name"}
	resp, err := RepoService.CreateRepo(request)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Error creating repo", err.GetMessage())
	assert.EqualValues(t, http.StatusInternalServerError, err.GetStatus())
}

func TestCreateRepoSuccess(t *testing.T) {
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
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

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{Name: ""}
	output := make(chan *repositories.CreateReposResult)
	service := &repoService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, "Invalid repo name", result.Error.GetMessage())
}

func TestCreateRepoConcurrentErrorFromGitHub(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errs.New(
			"Error creating repo",
		)
	}
	request := repositories.CreateRepoRequest{Name: "Test Repo Name"}
	output := make(chan *repositories.CreateReposResult)
	service := &repoService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, "Error creating repo", result.Error.GetMessage())
	assert.EqualValues(t, http.StatusInternalServerError, result.Error.GetStatus())
}

func TestCreateRepoConcurrentSuccess(t *testing.T) {
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}, nil
	}
	request := repositories.CreateRepoRequest{Name: "Test Name"}
	output := make(chan *repositories.CreateReposResult)
	service := &repoService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.NotNil(t, result)
	assert.NotNil(t, result.Response)
	assert.Nil(t, result.Error)
	assert.EqualValues(t, "Test Name", result.Response.Name)
	assert.EqualValues(t, "octocat", result.Response.Owner)
}

func TestHandleRepoResults(t *testing.T) {
	input := make(chan *repositories.CreateReposResult)
	output := make(chan *repositories.CreateReposResponse)
	var wg sync.WaitGroup

	service := &repoService{}
	go service.handleRepoResults(&wg, input, output)
	wg.Add(1)
	go func() {
		input <- &repositories.CreateReposResult{
			Error: errors.BadRequestError("Invalid repo name"),
		}
	}()
	wg.Wait()
	close(input)
	response := <-output
	result := response.Results[0]
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.GetStatus())
	assert.Nil(t, result.Response)
}

func TestCreateReposInvalidRequest(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: ""},
	}
	resp, err := RepoService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	first := resp.Results[0]
	second := resp.Results[1]
	assert.EqualValues(t, http.StatusBadRequest, resp.StatusCode)
	assert.EqualValues(t, "Invalid repo name", first.Error.GetMessage())
	assert.EqualValues(t, "Invalid repo name", second.Error.GetMessage())

}

func TestCreateReposPartialInvalidRequest(t *testing.T) {
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	r := ioutil.NopCloser(strings.NewReader(json))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}, nil
	}
	requests := []repositories.CreateRepoRequest{
		{Name: "Test Name"},
		{Name: ""},
	}
	resp, err := RepoService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	failure := resp.Results[0]
	success := resp.Results[1]
	assert.EqualValues(t, http.StatusPartialContent, resp.StatusCode)
	assert.EqualValues(t, "Test Name", success.Response.Name)
	assert.EqualValues(t, "Invalid repo name", failure.Error.GetMessage())
}

func TestCreateReposSuccess(t *testing.T) {
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(strings.NewReader(json))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}, nil
	}
	requests := []repositories.CreateRepoRequest{
		{Name: "Test Name"},
		{Name: "Test Name 2"},
	}
	resp, err := RepoService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	first := resp.Results[0]
	second := resp.Results[1]
	assert.EqualValues(t, http.StatusCreated, resp.StatusCode)
	assert.EqualValues(t, "Test Name", first.Response.Name)
	assert.EqualValues(t, "Test Name", second.Response.Name)
}
