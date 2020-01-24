package config

import "os"

const (
	secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	// LogLevel describes the log level
	LogLevel      = "info"
	goEnvironment = "GO_ENV"
	production    = "production"
)

var (
	githubAccessToken = os.Getenv(secretGithubAccessToken)
)

// GetGithubAccessToken returns the GH access token
func GetGithubAccessToken() string {
	return githubAccessToken
}

// IsProduction returns true for a prod env
func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
