package config

import "os"

// GithubConfig is the configuration for the application
type GithubConfig struct {
	GithubURL    string
	GithubApiURL string
	ClientID     string
	ClientSecret string
}

// NewAppConfig creates a new GithubConfig
func NewGithubConfig() *GithubConfig {

	return &GithubConfig{
		GithubURL:    os.Getenv("GITHUB_URL"),
		GithubApiURL: os.Getenv("GITHUB_API_URL"),
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	}
}
