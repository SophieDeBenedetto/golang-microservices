package services

import (
	"strings"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/github"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/repositories"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/errors"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
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
	request := github.CreateRepoRequest{
		Name: input.Name,
	}

	return &repositories.CreateRepoResponse{}, nil
}
