package oauth

import "github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/errors"

// User describes a user
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

const (
	getUserByUsernameAndPasswordQuery = "SELECT id, username FROM users WHERE username = ? and password = ?"
)

var (
	users = map[string]*User{
		"sophie": &User{
			ID:       123,
			Username: "sophie",
		},
	}
)

// GetUserByUsernameAndPassword returns the user with the given username and password
func GetUserByUsernameAndPassword(username string, password string) (*User, errors.APIError) {
	user := users[username]
	if user == nil {
		return nil, errors.NotFound("User not found")
	}
	return user, nil
}
