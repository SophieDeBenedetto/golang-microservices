package services

import (
	"time"

	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/errors"
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/oauth"
)

type oauthService struct {
}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APIError)
}

var (
	// OauthService is the service instance
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (service *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	token := &oauth.AccessToken{
		UserID:  user.ID,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}
	if err := token.Save(); err != nil {
		return nil, err
	}
	return token, nil
}

func (service *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APIError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}
	if token.IsExpired() {
		return nil, errors.NotFound("Expired token")
	}
	return token, nil
}
