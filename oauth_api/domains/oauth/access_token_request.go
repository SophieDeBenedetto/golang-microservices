package oauth

import (
	"strings"

	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/errors"
)

//AccessTokenRequest describes the request for a new access token
type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate validates the password and username of ther request
func (r *AccessTokenRequest) Validate() errors.APIError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.BadRequestError("invalid username")
	}

	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.BadRequestError("invalid Password")
	}
	return nil
}
