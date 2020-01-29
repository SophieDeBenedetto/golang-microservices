package oauth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/errors"
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/oauth"
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/services"
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/utils/mocks"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/testutils"
	"github.com/stretchr/testify/assert"
)

func init() {
	services.OauthService = &mocks.OauthService{}
}

func TestCreateAccessToken(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/oauth/access_token", strings.NewReader(`{"username":"sophie", "password":"password"}`))
	c := testutils.MockContext(request, response)

	mocks.CreateAccessTokenFunc = func(req oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError) {
		return &oauth.AccessToken{UserID: 123}, nil
	}

	CreateAccessToken(c)
	responseStruct, _ := newAPIResponseFromBytes(response.Body.Bytes())
	assert.EqualValues(t, responseStruct.UserID, 123)
}

func newAPIResponseFromBytes(bytes []byte) (*oauth.AccessToken, error) {
	var result oauth.AccessToken
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
