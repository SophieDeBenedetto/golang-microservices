package config

import "os"

const (
	apiGithubAccessToken = "SECRET_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

// GetGithubAccessToken returns the GH access token
func GetGithubAccessToken() string {
	return githubAccessToken
}
