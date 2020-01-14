package repositories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/testutils"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/clients/restclient"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/repositories"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/errors"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func init() {
	restclient.Client = &mocks.MockClient{}
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	response := httptest.NewRecorder()
	request := &http.Request{}
	c := testutils.MockContext(request, response)

	CreateRepo(c)
	apiError, err := errors.NewAPIErrorFromBytes(response.Body.Bytes())

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, "invalid JSON body", apiError.GetMessage())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	mocks.MockHTTPResponse(`{"message":"Invalid token", "status_code":400}`, 400)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"test repo"}`))
	c := testutils.MockContext(request, response)

	CreateRepo(c)
	apiError, err := errors.NewAPIErrorFromBytes(response.Body.Bytes())

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, "Invalid token", apiError.GetMessage())
}

func TestCreateRepoSucces(t *testing.T) {
	mocks.MockHTTPResponse(`{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`, 200)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"Test Name"}`))
	c := testutils.MockContext(request, response)

	CreateRepo(c)
	responseStruct, _ := newAPIResponseFromBytes(response.Body.Bytes())

	assert.EqualValues(t, http.StatusCreated, response.Code)
	assert.EqualValues(t, "Test Name", responseStruct.Name)
}

func newAPIResponseFromBytes(bytes []byte) (*repositories.CreateRepoResponse, error) {
	var result repositories.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
