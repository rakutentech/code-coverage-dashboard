package config

import "os"

// AppConfig is the configuration for the application
type AppConfig struct {
	AppHost         string
	AppEnv          string
	AppDir          string
	AppBaseURL      string
	FileServerURL   string
	NextJSServerURL string
	AppName         string
	AssetsDir       string
	AppURL          string
	AppSecure       bool
	LogDir          string
	PidDir          string
	SessionSecret   string
}

// NewAppConfig creates a new AppConfig
func NewAppConfig() *AppConfig {

	return &AppConfig{
		AppHost:         os.Getenv("APP_HOST"),
		AppEnv:          os.Getenv("APP_ENV"),
		AppDir:          os.Getenv("APP_DIR"),
		AppBaseURL:      os.Getenv("APP_BASE_URL"),
		FileServerURL:   os.Getenv("FILE_SERVER_URL"),
		NextJSServerURL: os.Getenv("NEXTJS_SERVER_URL"),
		AppName:         os.Getenv("APP_NAME"),
		AssetsDir:       os.Getenv("ASSETS_DIR"),
		AppURL:          os.Getenv("APP_URL"),
		AppSecure:       os.Getenv("APP_SECURE") == "true",
		LogDir:          os.Getenv("LOG_DIR"),
		PidDir:          os.Getenv("PID_DIR"),
		SessionSecret:   os.Getenv("SESSION_SECRET"),
	}
}
