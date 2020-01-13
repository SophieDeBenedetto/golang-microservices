package config

import "os"

const (
	apiGithubAcessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv("SECRET_GITHUB_ACCESS_TOKEN")
)

// GetGithubAccessToken returns the GH access token
func GetGithubAccessToken() string {
	return githubAccessToken
}
