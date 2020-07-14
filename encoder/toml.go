package encoder

import "github.com/BurntSushi/toml"

// TomlEncoder toml编码器
type TomlEncoder struct {
}

// NewTomlEncoder 创建TomlEncoder编码器
func NewTomlEncoder() *TomlEncoder {
	return &TomlEncoder{}
}

// LoadFile 加载toml配置文件
func (c *TomlEncoder) LoadFile(file string) (map[string]interface{}, error) {
	cfg := make(map[string]interface{})
	_, err := toml.DecodeFile(file, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
