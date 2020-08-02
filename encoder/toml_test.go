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

func TestTOMLLoadMemory(t *testing.T) {
	t.Run("Toml LoadMemory success", func(t *testing.T) {
		str := `
		redis= "127.0.0.1:3679"
	`
		encoder := NewTomlEncoder()
		cfg, err := encoder.LoadMemory(str)
		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		assert.Equal(t, "127.0.0.1:3679", cfg["redis"].(string))
	})

	t.Run("Toml LoadMemory error", func(t *testing.T) {
		str := `
		{}
	`
		encoder := NewTomlEncoder()
		cfg, err := encoder.LoadMemory(str)
		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})

}
