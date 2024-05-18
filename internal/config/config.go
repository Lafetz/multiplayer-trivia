package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port  int
	DbUrl string
}

func (c *Config) loadEnv() {

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("couldn't find Db Url")
	}
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Error converting PORT to int: %v\n", err)
		return
	}

	c.DbUrl = dbUrl
	c.Port = port

}
func NewConfig() *Config {
	config := Config{}
	config.loadEnv()
	return &config
}
