package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
)

type Config struct {
	DatabaseHost     string `validator:"required,hostname"`
	DatabasePort     string `validator:"required,numeric"`
	DatabaseUsername string `validator:"required"`
	DatabasePassword string `validator:"required"`
	DatabaseName     string `validator:"required"`
}

var localConfig *Config

func GetConfig() (*Config, error) {
	if localConfig == nil {
		localConfig = &Config{
			DatabaseHost:     os.Getenv("DATABASE_HOST"),
			DatabasePort:     os.Getenv("DATABASE_PORT"),
			DatabaseUsername: os.Getenv("DATABASE_USERNAME"),
			DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
			DatabaseName:     os.Getenv("DATABASE_NAME"),
		}
	}

	return localConfig, localConfig.validate()
}

func (c *Config) validate() error {
	return validator.New().Struct(c)
}

func (c *Config) DatabaseURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", c.DatabaseUsername, c.DatabasePassword, c.DatabaseHost, c.DatabasePort)
}

func (c *Config) Clear() {
	c = nil
}
