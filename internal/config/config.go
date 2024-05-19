package config

import (
	"errors"
	"os"
	"strconv"
)

var (
	ErrInvalidDbUrl = errors.New("db url is invalid")
	ErrInvalidPort  = errors.New("port number is invalid")
	ErrInvalidHost  = errors.New("ws url is invalid")
)

type Config struct {
	Port  int
	DbUrl string
	WsUrl string
} //

func (c *Config) loadEnv() error {
	wsUrl := os.Getenv("WS_URL")
	if wsUrl == "" {
		return ErrInvalidHost
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return ErrInvalidDbUrl
	}
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return ErrInvalidPort
	}

	c.DbUrl = dbUrl
	c.Port = port
	c.WsUrl = wsUrl
	return nil
}
func NewConfig() (*Config, error) {
	config := Config{}
	if err := config.loadEnv(); err != nil {
		return nil, err
	}
	return &config, nil
}
