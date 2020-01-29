package oauth

import (
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

// Save persists the access token (mocked)
func (at *AccessToken) Save() errors.APIError {
	at.AccessToken = "1234"
	tokens[at.AccessToken] = at
	return nil
}

// GetAccessTokenByToken retrieves the token for the given token string
func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.APIError) {
	token := tokens[accessToken]
	if token != nil {
		return token, nil
	}
	return nil, errors.NotFound("No token found")
}
