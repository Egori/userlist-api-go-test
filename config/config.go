package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP_PORT         string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("[ERROR] Unable to load .env file: %v", err)
		return nil, err
	}

	config := &Config{
		HTTP_PORT:         os.Getenv("HTTP_PORT"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// Validate the configuration to ensure all required fields are set
func validateConfig(config *Config) error {
	if config.HTTP_PORT == "" || config.POSTGRES_USER == "" || config.POSTGRES_PASSWORD == "" || config.POSTGRES_DB == "" {
		return fmt.Errorf("missing required configuration")
	}
	return nil
}
