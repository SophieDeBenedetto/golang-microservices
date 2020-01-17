package services

import (
	"net/http"
	"sync"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/config"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/github"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/providers/githubprovider"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/repositories"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/errors"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
	CreateRepos(input []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.APIError)
}

var (
	// RepoService of type repoServiceInterface
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	if err := repositories.Validate(input); err != nil {
		return nil, err
	}
	request := &github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := githubprovider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		apiError := errors.NewAPIError(err.StatusCode, err.Message, err.Errors)
		return nil, apiError
	}
	return &repositories.CreateRepoResponse{
		ID:    response.ID,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}, nil
}

// CreateRepos creates repos
func (s *repoService) CreateRepos(request []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.APIError) {
	input := make(chan *repositories.CreateReposResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)
	var wg sync.WaitGroup

	go s.handleRepoResults(&wg, input, output)

	for _, current := range request {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	response := <-output
	calculateResponseStatus(request, &response)
	return &response, nil
}

func (s *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan *repositories.CreateReposResult) {
	if err := repositories.Validate(input); err != nil {
		output <- &repositories.CreateReposResult{
			Error: err,
		}
		return
	}
	result, err := s.CreateRepo(input)
	if err != nil {
		output <- &repositories.CreateReposResult{
			Error: err,
		}
		return
	}
	output <- &repositories.CreateReposResult{
		Response: result,
	}
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan *repositories.CreateReposResult, output chan repositories.CreateReposResponse) {
	var response repositories.CreateReposResponse
	for result := range input {
		response.Results = append(response.Results, result)
		wg.Done()
	}
	output <- response
}

func calculateResponseStatus(request []repositories.CreateRepoRequest, response *repositories.CreateReposResponse) {
	successCreations := 0
	for _, result := range response.Results {
		if result.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		(*response).StatusCode = response.Results[0].Error.GetStatus()
	} else if successCreations == len(request) {
		(*response).StatusCode = http.StatusCreated
	} else {
		(*response).StatusCode = http.StatusPartialContent
	}
}
