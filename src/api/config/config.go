package config

import "os"

const (
	secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(secretGithubAccessToken)
)

// GetGithubAccessToken returns the GH access token
func GetGithubAccessToken() string {
	return githubAccessToken
}
