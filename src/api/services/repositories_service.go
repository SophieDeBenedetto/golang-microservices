package services

import (
	"strings"

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
	CreateRepos(input repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.APIError)
}

var (
	// RepoService of type repoServiceInterface
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	repoName := strings.TrimSpace(input.Name)
	if repoName == "" {
		err := errors.BadRequestError("Invalid repo name")
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
func CreateRepos(request []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.APIError) {

}
