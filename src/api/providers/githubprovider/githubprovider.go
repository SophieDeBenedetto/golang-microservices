package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/clients/restclient"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	createRepoURL             = "https://api.github.com/user/repos"
)

// CreateRepo sends the create repo request to the GitHub API
func CreateRepo(accessToken string, request *github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	// create headers
	header := getAuthorizationHeader(accessToken)
	headers := http.Header{}
	headers.Set(headerAuthorization, header)

	// make request
	response, err := restclient.Post(createRepoURL, request, headers)
	// check request success/failure
	if err != nil {
		x := err.Error()
		fmt.Println(x)
		log.Println(fmt.Sprintf("Error creating GitHub repo: %s", err.Error()))
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	// read response body
	bytes, err := ioutil.ReadAll(response.Body)
	// leverage defer stack to defer closing of response body read operation
	// this will defer until this function is ready to return
	defer response.Body.Close()
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}

	// check response status
	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse
		// decode response body JSON into github.ErrorResponse struct
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			log.Println(err)
			return nil, &github.ErrorResponse{
				StatusCode: response.StatusCode,
				Message:    "Invalid JSON response body",
			}
		}
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: response.StatusCode,
			Message:    "Error trying to decode successfull create repo JSON response body",
		}
	}
	return &result, nil
}

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}
