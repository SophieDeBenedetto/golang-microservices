package githubprovider

import (
	"errors"
	"net/http"
	"testing"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/clients/restclient"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/github"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	accessToken := "ABC123"
	assert.EqualValues(t, "token ABC123", getAuthorizationHeader(accessToken))
}

func TestCreateRepoFailure(t *testing.T) {
	restclient.StartMocks()
	// & generates a pointer
	restclient.AddMock(&restclient.Mock{
		URL:        createRepoURL,
		HTTPMethod: http.MethodPost,
		Error:      errors.New("Invalid rest client response"),
	})
	response, err := CreateRepo("invalidToken", &github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
}
