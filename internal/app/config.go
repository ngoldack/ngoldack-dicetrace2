package app

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
}

var localConfig *Config

func GetConfig() (*Config, error) {
	if localConfig == nil {
		localConfig = &Config{
			DatabaseHost:     os.Getenv("DATABASE_HOST"),
			DatabasePort:     os.Getenv("DATABASE_PORT"),
			DatabaseUsername: os.Getenv("DATABASE_USERNAME"),
			DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		}
	}

	return localConfig, localConfig.validate()
}

func (c *Config) validate() error {
	// TODO implement config validation
	return nil
}

func (c *Config) DatabaseURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", c.DatabaseUsername, c.DatabasePassword, c.DatabaseHost, c.DatabasePort)
}
