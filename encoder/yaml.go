package encoder

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// YamlEncoder yaml编码器
type YamlEncoder struct {
}

// NewYamlEncoder 创建yaml编码器
func NewYamlEncoder() *YamlEncoder {
	return &YamlEncoder{}
}

// LoadFile 加载yaml配置文件
func (c *YamlEncoder) LoadFile(file string) (map[string]interface{}, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return c.LoadMemory(string(buf))
}

// LoadMemory 加内存配置文件
func (c *YamlEncoder) LoadMemory(config string) (map[string]interface{}, error) {
	cfg := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(config), &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
