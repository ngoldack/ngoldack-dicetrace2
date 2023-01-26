package app

import (
	"github.com/go-playground/validator/v10"
	"os"
)
import _ "github.com/joho/godotenv/autoload"

type Config struct {
	Neo4JURI      string `validate:"required,uri"`
	Neo4JUsername string `validate:"required"`
	Neo4JPassword string `validate:"required"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{
		Neo4JURI:      os.Getenv("NEO4J_URI"),
		Neo4JUsername: os.Getenv("NEO4J_USERNAME"),
		Neo4JPassword: os.Getenv("NEO4J_PASSWORD"),
	}

	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
