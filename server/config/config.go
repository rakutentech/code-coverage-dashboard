package config

import (
	"sync"
)

var (
	configInstance *Config
	configOnce     sync.Once
)

// Config is the configuration struct that returns all the configs
type Config struct {
	AppConfig    *AppConfig
	DBConfig     *DBConfig
	GithubConfig *GithubConfig
}

// NewConfig returns a new Config struct with the configs
// Is a singleton with one memory address
func NewConfig() *Config {
	configOnce.Do(func() {
		configInstance = &Config{
			AppConfig:    NewAppConfig(),
			DBConfig:     NewDBConfig(),
			GithubConfig: NewGithubConfig(),
		}
	})
	return configInstance
}
