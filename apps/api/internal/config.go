package internal

import (
	"github.com/go-playground/validator/v10"
	"os"
)
import _ "github.com/joho/godotenv/autoload"

type Config struct {
	ApiPort          string `validate:"required,number"`
	AuthClientID     string `validate:"required"`
	AuthClientSecret string `validate:"required"`
	AuthIssuer       string `validate:"required,uri"`
	AuthAudience     string `validate:"required"`

	KafkaHost     string `validate:"required"`
	KafkaUsername string `validate:"required"`
	KafkaPassword string `validate:"required"`

	RedisURI string `validate:"required,uri"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{
		ApiPort:          os.Getenv("API_PORT"),
		AuthClientID:     os.Getenv("AUTH_CLIENT_ID"),
		AuthClientSecret: os.Getenv("AUTH_CLIENT_SECRET"),
		AuthIssuer:       os.Getenv("AUTH_ISSUER"),
		AuthAudience:     os.Getenv("AUTH_AUDIENCE"),
		KafkaHost:        os.Getenv("KAFKA_HOST"),
		KafkaUsername:    os.Getenv("KAFKA_USERNAME"),
		KafkaPassword:    os.Getenv("KAFKA_PASSWORD"),
		RedisURI:         os.Getenv("REDIS_URI"),
	}

	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
