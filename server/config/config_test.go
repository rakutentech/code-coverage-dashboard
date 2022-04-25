package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var conf *Config

func init() {
	Setup()
	conf = NewConfig()
}

func TestEnv(t *testing.T) {
	assert.Equal(t, conf.AppConfig.AppEnv, "testing")
	assert.NotNil(t, conf.AppConfig.AppHost)
	assert.NotNil(t, conf.AppConfig.AppURL)
	assert.NotNil(t, conf.AppConfig.LogDir)
	assert.NotNil(t, conf.AppConfig.PidDir)
}
