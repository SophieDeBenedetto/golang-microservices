package repositories

import (
	"strings"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/errors"
)

// CreateRepoRequest describes repo to be created
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateRepoResponse describes created repo
type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"string"`
}

//CreateReposResponse describes created repos
type CreateReposResponse struct {
	StatusCode int                  `json:"status"`
	Results    []*CreateReposResult `json:"results"`
}

// CreateReposResult the result of creating multiple repos
type CreateReposResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.APIError     `json:"error"`
}

// Validate validates the create repo request
func Validate(r CreateRepoRequest) errors.APIError {
	repoName := strings.TrimSpace(r.Name)
	if repoName == "" {
		err := errors.BadRequestError("Invalid repo name")
		return err
	}
	return nil
}
