package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DataEndpoint          string
	DataProcessedEndpoint string
	AuthToken             string // токен доступа к источнику данных для репоста
	VkToken               string
}

func NewConfig() (*Config, error) {
	pd, _ := os.Getwd()
	parts := strings.Split(pd, "/internal/")
	projectDir := parts[0]

	if err := godotenv.Load(projectDir + "/.env"); err != nil {
		return nil, errors.New("No .env file found")
	}

	cfg := &Config{
		DataEndpoint:          os.Getenv("DATA_ENDPOINT"),
		DataProcessedEndpoint: os.Getenv("DATA_PROCESSED_ENDPOINT"),
		AuthToken:             os.Getenv("AUTH_TOKEN"),
		VkToken:               os.Getenv("VK_TOKEN"),
	}

	return cfg, nil
}
