package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewYamlEncoder(t *testing.T) {
	encoder := NewYamlEncoder()
	assert.NotNil(t, encoder)
}

func TestYamlLoadFile(t *testing.T) {
	encoder := NewYamlEncoder()

	t.Run("LoadFile success", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "config.yaml")
		assert.Nil(t, err)
		assert.NotNil(t, cfg)
	})

	t.Run("LoadFile error", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "config.toml")
		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})

	t.Run("LoadFile file is not exist", func(t *testing.T) {
		cfg, err := encoder.LoadFile(configFilePath + "db.yaml")
		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})
}

func TestYamlLoadMemory(t *testing.T) {
	str := `
---
redis: 127.0.0.1:3679
`
	encoder := NewYamlEncoder()
	cfg, err := encoder.LoadMemory(str)
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "127.0.0.1:3679", cfg["redis"].(string))
}
