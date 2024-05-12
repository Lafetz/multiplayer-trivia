package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     int
	DbUrl    string
	HashKey  string
	BlockKey string
}

func (c *Config) loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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

	hashKey := os.Getenv("HASH_KEY")
	if hashKey == "" {
		log.Fatal("couldn't find hash key")
	}

	blockKey := os.Getenv("BLOCK_KEY")
	if blockKey == "" {
		log.Fatal("couldn't find block key")
	}
	c.DbUrl = dbUrl
	c.Port = port
	c.HashKey = hashKey
	c.BlockKey = blockKey
}
func NewConfig() *Config {
	config := Config{}
	config.loadEnv()
	return &config
}