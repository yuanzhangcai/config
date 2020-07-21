package config

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yuanzhangcai/config/encoder"
)

func TestNewJSONConfig(t *testing.T) {
	cfg := newJSONConfig()
	assert.NotNil(t, cfg)
}

func TestLoadFile(t *testing.T) {
	cfg := newJSONConfig()

	t.Run("LoadFile load toml file", func(t *testing.T) {
		err := cfg.LoadFile(configFilePath + "config.toml")
		assert.Nil(t, err)
		assert.Equal(t, "toml_cfg", cfg.GetString("app_desc"))
		assert.Equal(t, "toml", cfg.GetString("common", "config_toml"))
	})

	t.Run("LoadFile load json file", func(t *testing.T) {
		err := cfg.LoadFile(configFilePath + "config.json")
		assert.Nil(t, err)
		assert.Equal(t, "json_cfg", cfg.GetString("app_desc"))
		assert.Equal(t, "json", cfg.GetString("common", "config_json"))
	})

	t.Run("LoadFile load yaml file", func(t *testing.T) {
		err := cfg.LoadFile(configFilePath + "config.yaml")
		assert.Nil(t, err)
		assert.Equal(t, "yaml_cfg", cfg.GetString("app_desc"))
		assert.Equal(t, "yaml", cfg.GetString("common", "config_yaml"))
	})

	t.Run("LoadFile load unsupported file", func(t *testing.T) {
		err := cfg.LoadFile(configFilePath + "config.ini")
		assert.Equal(t, ErrUnsupportedFileFormat, err)
	})

	t.Run("LoadFile load unsupported file 2", func(t *testing.T) {
		err := cfg.LoadFile(configFilePath + "config")
		assert.Equal(t, ErrUnsupportedFileFormat, err)
	})
}

func TestLoadFileWithEncoder(t *testing.T) {
	cfg := newJSONConfig()

	t.Run("LoadFileWithEncoder file is not exist", func(t *testing.T) {
		err := cfg.LoadFileWithEncoder(configFilePath+"db.toml", encoder.NewTomlEncoder())
		assert.NotNil(t, err)
	})

	t.Run("LoadFileWithEncoder file format is not right", func(t *testing.T) {
		err := cfg.LoadFileWithEncoder(configFilePath+"config.toml", encoder.NewYamlEncoder())
		assert.NotNil(t, err)
	})

	t.Run("LoadFileWithEncoder load toml file", func(t *testing.T) {
		err := cfg.LoadFileWithEncoder(configFilePath+"config.toml", encoder.NewTomlEncoder())
		assert.Nil(t, err)
	})

	t.Run("LoadFileWithEncoder load json file", func(t *testing.T) {
		err := cfg.LoadFileWithEncoder(configFilePath+"config.json", encoder.NewJSONEncoder())
		assert.Nil(t, err)
	})

	t.Run("LoadFileWithEncoder load yaml file", func(t *testing.T) {
		err := cfg.LoadFileWithEncoder(configFilePath+"config.yaml", encoder.NewYamlEncoder())
		assert.Nil(t, err)
	})
}

func TestListenWatch(t *testing.T) {
	file := configFilePath + "config.json"
	buf, _ := ioutil.ReadFile(file)

	cfg := newJSONConfig()
	_ = cfg.LoadFile(configFilePath + "config.toml")
	_ = cfg.LoadFile(configFilePath + "config.json")
	_ = cfg.LoadFile(configFilePath + "config.yaml")

	strCfg := `{
		"common" : {
			"redis" : "127.0.0.1"
		}
	}`
	_ = ioutil.WriteFile(file, []byte(strCfg), 0777)
	time.Sleep(time.Second * 1)

	assert.Equal(t, "127.0.0.1", cfg.GetString("common", "redis"))

	_ = ioutil.WriteFile(file, buf, 0777)
}

func TestLoadOsEnv(t *testing.T) {
	os.Setenv("test_env", "dev")
	cfg := newJSONConfig()
	assert.Equal(t, "dev", cfg.GetString("test_env"))
}
