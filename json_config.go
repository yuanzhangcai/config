package config

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"
	"sync"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/imdario/mergo"
	"github.com/yuanzhangcai/config/encoder"
)

var (
	// ErrUnsupportedFileFormat 暂不支持该配置文件格式。
	ErrUnsupportedFileFormat = errors.New("暂不支持该配置文件格式。")

	encoding = map[string]encoder.Encoder{
		"toml": encoder.NewTomlEncoder(),
		"json": encoder.NewJSONEncoder(),
		"yaml": encoder.NewYamlEncoder(),
	}
)

// JSONConfig json配置文件
type JSONConfig struct {
	sync.RWMutex
	cfg  map[string]interface{}
	json *simplejson.Json
}

func newJSONConfig() *JSONConfig {
	return &JSONConfig{}
}

// LoadFile 加载配置文件
func (c *JSONConfig) LoadFile(file string) error {
	ext := strings.ToLower(path.Ext(file))
	if ext == "" {
		return ErrUnsupportedFileFormat
	}

	ext = ext[1:]
	loader, ok := encoding[ext]
	if !ok {
		return ErrUnsupportedFileFormat
	}
	return c.LoadFileWithEncoder(file, loader)
}

// LoadFileWithEncoder 指定编码器加载配置文件
func (c *JSONConfig) LoadFileWithEncoder(file string, loader encoder.Encoder) error {
	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	temp, err := loader.LoadFile(file)
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()

	err = mergo.Merge(&c.cfg, temp)
	if err != nil {
		return err
	}

	buf, err := json.Marshal(c.cfg)
	if err != nil {
		return err
	}

	c.json, err = simplejson.NewJson(buf)
	if err != nil {
		return err
	}

	return nil
}

// Get 获取interface配置
func (c *JSONConfig) Get(keys ...string) interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).Interface()
}

// GetString 获取string配置
func (c *JSONConfig) GetString(keys ...string) string {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustString()
}

// GetStringArray 获取map数据
func (c *JSONConfig) GetStringArray(keys ...string) []string {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustStringArray()
}

// GetBool 获取bool配置
func (c *JSONConfig) GetBool(keys ...string) bool {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustBool()
}

// GetInt 获取int配置
func (c *JSONConfig) GetInt(keys ...string) int {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustInt()
}

// GetInt64 获取int64配置
func (c *JSONConfig) GetInt64(keys ...string) int64 {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustInt64()
}

// GetUint64 获取uint64配置
func (c *JSONConfig) GetUint64(keys ...string) uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustUint64()
}

// GetFloat64 获取float64配置
func (c *JSONConfig) GetFloat64(keys ...string) float64 {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustFloat64()
}

// GetArray 获取interface数组
func (c *JSONConfig) GetArray(keys ...string) []interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustArray()
}

// GetMap 获取map数据
func (c *JSONConfig) GetMap(keys ...string) map[string]interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.json.GetPath(keys...).MustMap()
}

// Scan 读取配置到指定对象
func (c *JSONConfig) Scan(keys []string, value interface{}) error {
	buf, err := c.json.GetPath(keys...).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, value)
	if err != nil {
		return err
	}
	return nil
}

// Set 写入配置
func (c *JSONConfig) Set(key string, value interface{}) {
	c.Lock()
	c.json.Set(key, value)
	c.Unlock()
}

// SetPath 写入多级配置
func (c *JSONConfig) SetPath(keys []string, value interface{}) {
	c.Lock()
	c.json.SetPath(keys, value)
	c.Unlock()
}
