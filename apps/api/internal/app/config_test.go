package app

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {

	t.Run("valid config test", func(t *testing.T) {
		mongoDBURI := "mongodb+srv://dicetrace:abc@dev.dbc.net/?retryWrites=true&w=majority"
		os.Clearenv()
		err := os.Setenv("MONGODB_URI", mongoDBURI)
		assert.NoError(t, err)

		cfg, err := GetConfig()
		if assert.NoError(t, err) {
			assert.Equal(t, mongoDBURI, cfg.MongoDBURI)
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
