package mocks

import (
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/errors"
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/oauth"
)

var (
	// CreateAccessTokenFunc gets the mock function
	CreateAccessTokenFunc func(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError)
	// GetAccessTokenFunc gets the mock function
	GetAccessTokenFunc func(accessToken string) (*oauth.AccessToken, errors.APIError)
)

// OauthService the mock oauth service
type OauthService struct {
}

// CreateAccessToken mocks the access token creation
func (s *OauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError) {
	return CreateAccessTokenFunc(request)
}

// GetAccessToken mocks the access token fetch
func (s *OauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APIError) {
	return GetAccessTokenFunc(accessToken)
}
