package encoder

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var configFilePath = ""

func init() {
	configFilePath = os.Getenv("GITHUB_WORKSPACE")
	if configFilePath == "" {
		configFilePath = "/Users/zacyuan/MyWork/config"
	}
	configFilePath = configFilePath + "/examples/"
}

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

func TestJSONLoadMemory(t *testing.T) {
	str := `
		{
			"redis" : "127.0.0.1:3679"
		}
	`
	encoder := NewJSONEncoder()
	cfg, err := encoder.LoadMemory(str)
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "127.0.0.1:3679", cfg["redis"].(string))
}
