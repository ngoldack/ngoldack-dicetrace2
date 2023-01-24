package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	DatabaseURI      string `validator:"required_without_all=DatabaseHost DatabasePort DatabaseUsername DatabasePassword|uri"`
	DatabaseHost     string `validator:"required_without=DatabaseURI,hostname"`
	DatabasePort     string `validator:"required_without=DatabaseURI"`
	DatabaseUsername string `validator:"required_without=DatabaseURI"`
	DatabasePassword string `validator:"required_without=DatabaseURI"`
	DatabaseName     string `validator:"required_without=DatabaseURI"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{
		DatabaseURI:      os.Getenv("DATABASE_URI"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUsername: os.Getenv("DATABASE_USERNAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
	}

	if cfg.DatabaseURI == "" {
		cfg.DatabaseURI = cfg.generateDatabaseURI()
	}

	err := cfg.validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c Config) validate() error {
	log.Debug().Interface("config", c).Msg("validating config")
	return validator.New().Struct(c)
}

func (c Config) generateDatabaseURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", c.DatabaseUsername, c.DatabasePassword, c.DatabaseHost, c.DatabasePort)
}
