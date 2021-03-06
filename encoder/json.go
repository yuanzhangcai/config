package encoder

import (
	"encoding/json"
	"io/ioutil"
)

// JSONEncoder json文件编码器
type JSONEncoder struct {
}

// NewJSONEncoder 创建JSON文件编码器
func NewJSONEncoder() *JSONEncoder {
	return &JSONEncoder{}
}

// LoadFile 加载json配置文件
func (c *JSONEncoder) LoadFile(file string) (map[string]interface{}, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return c.LoadMemory(string(buf))
}

// LoadMemory 加载内存配置
func (c *JSONEncoder) LoadMemory(config string) (map[string]interface{}, error) {
	cfg := make(map[string]interface{})
	err := json.Unmarshal([]byte(config), &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
