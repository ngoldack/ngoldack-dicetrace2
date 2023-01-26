package app

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {

	t.Run("valid config test", func(t *testing.T) {
		os.Clearenv()
		// TODO test config
		_, err := GetConfig()
		if assert.NoError(t, err) {
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
		err := os.Setenv("MONGODB_URI", "this should fail")
		assert.NoError(t, err)

		cfg, err := GetConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
	})

}
