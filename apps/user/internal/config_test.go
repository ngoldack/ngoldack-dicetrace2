package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {

	t.Run("valid config test", func(t *testing.T) {
		os.Clearenv()

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

		cfg, err := GetConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
	})

}
