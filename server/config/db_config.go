package config

import (
	"os"
	"strconv"
)

// DBConfig is the configuration for the DB client.
type DBConfig struct {
	DBConnection string
	DBHost       string
	DBHostSlave  string
	DBPort       string
	DBDatabase   string
	DBUsername   string
	DBPassword   string
	DBLogLevel   int
}

// NewDBConfig returns a new DBConfig.
// DBHosts is a list of DB hosts, separated by commas.
func NewDBConfig() *DBConfig {
	dbLogLevel, _ := strconv.Atoi(os.Getenv("DB_LOG_LEVEL"))
	return &DBConfig{
		DBConnection: os.Getenv("DB_CONNECTION"),
		DBHost:       os.Getenv("DB_HOST"),
		DBHostSlave:  os.Getenv("DB_HOST_SLAVE"),
		DBPort:       os.Getenv("DB_PORT"),
		DBDatabase:   os.Getenv("DB_DATABASE"),
		DBUsername:   os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBLogLevel:   dbLogLevel,
	}
}
