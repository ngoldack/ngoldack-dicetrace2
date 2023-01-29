package internal

import (
	"github.com/go-playground/validator/v10"
	"os"
)
import _ "github.com/joho/godotenv/autoload"

type Config struct {
	KafkaHost     string `validate:"required"`
	KafkaUsername string `validate:"required"`
	KafkaPassword string `validate:"required"`

	RedisURI string `validate:"required,uri"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{
		KafkaHost:     os.Getenv("KAFKA_HOST"),
		KafkaUsername: os.Getenv("KAFKA_USERNAME"),
		KafkaPassword: os.Getenv("KAFKA_PASSWORD"),
		RedisURI:      os.Getenv("REDIS_URI"),
	}

	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
