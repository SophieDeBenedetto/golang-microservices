package config

import "os"

// GetGithubAccessToken returns the GH access token
func GetGithubAccessToken() string {
	return os.Getenv("apiGithubAccessToken")
}
