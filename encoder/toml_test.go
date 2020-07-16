package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTomlEncoder(t *testing.T) {
	encoder := NewTomlEncoder()
	assert.NotNil(t, encoder)
}

func TestTomlLoadFile(t *testing.T) {
	encoder := NewTomlEncoder()

	t.Run("LoadFile success", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "config.toml")
		assert.Nil(t, err)
		assert.NotNil(t, cfg)
	})

	t.Run("LoadFile error", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "config.json")
		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})

	t.Run("LoadFile file is not exist", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "db.toml")
		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})
}
