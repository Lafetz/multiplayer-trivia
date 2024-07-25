package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/Lafetz/showdown-trivia-game/internal/logger"
)

var (
	ErrInvalidDbUrl = errors.New("db url is invalid")
	ErrInvalidPort  = errors.New("port number is invalid")
	ErrLogLevel     = errors.New("log level not set")
	ErrInvliadEnv   = errors.New("env not set or invliad")
	ErrInvalidLevel = errors.New("invliad log level")
)

type Config struct {
	Port     int
	DbUrl    string
	LogLevel string
	Env      string
}

var Environment = map[string]string{
	"dev":  "development",
	"prod": "production",
}

func (c *Config) loadEnv() error {
	env := os.Getenv("ENV")

	if env == "" {
		return ErrInvliadEnv
	}
	if evalue, ok := Environment[env]; !ok {

		return ErrInvliadEnv
	} else {
		c.Env = evalue

	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return ErrInvalidDbUrl
	}
	//
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		return ErrLogLevel
	}
	_, ok := logger.LogLevels[logLevel]
	if !ok {
		return ErrInvalidLevel
	}
	//
	portStr := os.Getenv("PORT")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return ErrInvalidPort
	}

	c.DbUrl = dbUrl
	c.Port = port

	c.LogLevel = logLevel
	return nil
}
func NewConfig() (*Config, error) {
	config := Config{}
	if err := config.loadEnv(); err != nil {
		return nil, err
	}
	return &config, nil
}
