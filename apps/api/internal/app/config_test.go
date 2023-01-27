package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {

	t.Run("valid config test", func(t *testing.T) {
		uri := "neo4j://localhost:1111"

		os.Clearenv()
		err := os.Setenv("NEO4J_URI", uri)
		assert.NoError(t, err)
		err = os.Setenv("NEO4J_USERNAME", "username")
		assert.NoError(t, err)
		err = os.Setenv("NEO4J_PASSWORD", "password")
		assert.NoError(t, err)

		cfg, err := GetConfig()
		if assert.NoError(t, err) {
			assert.Equal(t, uri, cfg.Neo4JURI)
			assert.Equal(t, "username", cfg.Neo4JUsername)
			assert.Equal(t, "password", cfg.Neo4JPassword)
		}
	})

	t.Run("empty config test", func(t *testing.T) {
		os.Clearenv()
		cfg, err := GetConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
	})

	t.Run("invalid config test", func(t *testing.T) {
		os.Clearenv()
		err := os.Setenv("NEO4J_URI", "this should fail")
		assert.NoError(t, err)
		err = os.Setenv("NEO4J_USERNAME", "ok")
		assert.NoError(t, err)
		err = os.Setenv("NEO4J_PASSWORD", "ok")
		assert.NoError(t, err)

		cfg, err := GetConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
	})

}
