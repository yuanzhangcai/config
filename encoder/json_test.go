package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var configFilePath = "/Users/zacyuan/MyWork/config/examples/"

func TestNewJSONEncoder(t *testing.T) {
	encoder := NewJSONEncoder()
	assert.NotNil(t, encoder)
}

func TestJSONLoadFile(t *testing.T) {
	encoder := NewJSONEncoder()

	t.Run("LoadFile success", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "config.json")
		assert.Nil(t, err)
		assert.NotNil(t, cfg)
	})

	t.Run("LoadFile error", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "config.toml")
		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})

	t.Run("LoadFile file is not exist", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "db.json")
		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})
}
