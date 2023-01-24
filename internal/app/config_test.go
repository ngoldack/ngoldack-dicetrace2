package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {

	t.Run("valid config test", func(t *testing.T) {
		err := os.Setenv("DATABASE_HOST", "test")
		assert.NoError(t, err)
		err = os.Setenv("DATABASE_PORT", "1111")
		assert.NoError(t, err)
		err = os.Setenv("DATABASE_USERNAME", "test")
		assert.NoError(t, err)
		err = os.Setenv("DATABASE_PASSWORD", "test")
		assert.NoError(t, err)
		err = os.Setenv("DATABASE_NAME", "test")
		assert.NoError(t, err)

		cfg, err := GetConfig()
		if assert.NoError(t, err) {
			assert.Equal(t, "test", cfg.DatabaseHost)
			assert.Equal(t, "1111", cfg.DatabasePort)
			assert.Equal(t, "test", cfg.DatabaseUsername)
			assert.Equal(t, "test", cfg.DatabasePassword)
			assert.Equal(t, "test", cfg.DatabaseName)
		}
	})

	t.Run("invalid config test", func(t *testing.T) {

		err := os.Setenv("DATABASE_HOST", "1")
		assert.NoError(t, err)
		err = os.Setenv("DATABASE_PORT", "test")
		assert.NoError(t, err)
		err = os.Setenv("DATABASE_USERNAME", "test")
		assert.NoError(t, err)
		err = os.Setenv("DATABASE_PASSWORD", "test")
		assert.NoError(t, err)
		err = os.Setenv("DATABASE_NAME", "test")
		assert.NoError(t, err)

		cfg, err := GetConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
	})

}

func TestGetDatabaseURI(t *testing.T) {

	err := os.Setenv("DATABASE_HOST", "1")
	assert.NoError(t, err)
	err = os.Setenv("DATABASE_HOST", "abc")
	assert.NoError(t, err)
	err = os.Setenv("DATABASE_USERNAME", "test")
	assert.NoError(t, err)
	err = os.Setenv("DATABASE_PASSWORD", "test")
	assert.NoError(t, err)
	err = os.Setenv("DATABASE_NAME", "test")
	assert.NoError(t, err)

	cfg, err := GetConfig()
	assert.NoError(t, err)

	t.Run("valid uri test", func(t *testing.T) {
		err := validator.New().Var(cfg.DatabaseURI(), "uri")
		assert.NoError(t, err)
	})

	t.Run("invalid uri test", func(t *testing.T) {
		cfg.DatabaseHost = "123"
		cfg.DatabasePort = "test"

		err := validator.New().Var(cfg.DatabaseURI(), "uri")
		assert.Error(t, err)
	})

}
