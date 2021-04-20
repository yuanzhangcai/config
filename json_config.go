package config

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"
	"sync"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/fsnotify/fsnotify"
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

type loaderFunc func(config string) (map[string]interface{}, error)

const (
	fileConfig   = 1
	memoryConfig = 2
	envConfig    = 3
)

type source struct {
	Type   int
	Data   string
	loader loaderFunc
}

// JSONConfig json配置文件
type JSONConfig struct {
	m      sync.RWMutex
	cfg    map[string]interface{}
	json   *simplejson.Json
	source []*source
	watch  *fsnotify.Watcher
}

func newJSONConfig() *JSONConfig {
	cfg := &JSONConfig{
		json: simplejson.New(),
	}
	_ = cfg.LoadOsEnv() // 默认载入环境变量
	return cfg
}

func (c *JSONConfig) loadOsEnv(config string) (map[string]interface{}, error) {
	cfg := make(map[string]interface{})
	arr := os.Environ()
	for _, one := range arr {
		tmp := strings.Split(one, "=")
		if len(tmp) < 2 {
			continue
		}
		cfg[tmp[0]] = tmp[1]
	}
	return cfg, nil
}

// LoadOsEnv 载入环境变量
func (c *JSONConfig) LoadOsEnv() error {
	s := &source{Type: envConfig, loader: c.loadOsEnv}
	return c.loadAndAddSource(s)
}

// LoadMemory 加载内存配置
func (c *JSONConfig) LoadMemory(config, t string) error {
	loader, ok := encoding[t]
	if !ok {
		return ErrUnsupportedFileFormat
	}

	s := &source{Type: memoryConfig, Data: config, loader: loader.LoadMemory}
	return c.loadAndAddSource(s)
}

// LoadFile 加载配置文件
func (c *JSONConfig) LoadFile(file string) error {
	ext := strings.ToLower(path.Ext(file))
	if ext == "" {
		return ErrUnsupportedFileFormat
	}

	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	ext = ext[1:]
	loader, ok := encoding[ext]
	if !ok {
		return ErrUnsupportedFileFormat
	}

	s := &source{Type: fileConfig, Data: file, loader: loader.LoadFile}
	return c.loadAndAddSource(s)
}

func (c *JSONConfig) load(s *source) error {
	c.m.Lock()
	defer c.m.Unlock()

	cfg, err := s.loader(s.Data)
	if err != nil {
		return err
	}

	err = mergo.MergeWithOverwrite(&c.cfg, cfg)
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

func (c *JSONConfig) loadAndAddSource(s *source) error {
	if err := c.load(s); err != nil {
		return err
	}

	return c.addSource(s)
}

func (c *JSONConfig) addSource(s *source) error {
	c.m.Lock()
	defer c.m.Unlock()

	needWatch := true
	for index, one := range c.source {
		if one.Type == s.Type && one.Data == s.Data {
			c.source = append(c.source[0:index], c.source[index+1:]...)
			needWatch = false
		}
	}
	c.source = append(c.source, s)

	if s.Type == fileConfig && needWatch {
		if c.watch == nil {
			watch, err := fsnotify.NewWatcher()
			if err != nil {
				return err
			}
			c.watch = watch
			c.listenWatch()
		}
		err := c.watch.Add(s.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *JSONConfig) listenWatch() {
	go func() {
		for event := range c.watch.Events {
			c.reload(event.Name)
		}
	}()
}

func (c *JSONConfig) reload(file string) {
	index := 0
	for i, one := range c.source {
		if one.Type == fileConfig && one.Data == file {
			index = i
			break
		}
	}

	for i := index; i < len(c.source); i++ {
		_ = c.load(c.source[i])
	}
}

// Get 获取interface配置
func (c *JSONConfig) Get(keys ...string) interface{} {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).Interface()
}

// GetString 获取string配置
func (c *JSONConfig) GetString(keys ...string) string {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustString()
}

// GetStringArray 获取map数据
func (c *JSONConfig) GetStringArray(keys ...string) []string {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustStringArray()
}

// GetBool 获取bool配置
func (c *JSONConfig) GetBool(keys ...string) bool {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustBool()
}

// GetInt 获取int配置
func (c *JSONConfig) GetInt(keys ...string) int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustInt()
}

// GetInt64 获取int64配置
func (c *JSONConfig) GetInt64(keys ...string) int64 {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustInt64()
}

// GetUint64 获取uint64配置
func (c *JSONConfig) GetUint64(keys ...string) uint64 {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustUint64()
}

// GetFloat64 获取float64配置
func (c *JSONConfig) GetFloat64(keys ...string) float64 {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustFloat64()
}

// GetArray 获取interface数组
func (c *JSONConfig) GetArray(keys ...string) []interface{} {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustArray()
}

// GetMap 获取map数据
func (c *JSONConfig) GetMap(keys ...string) map[string]interface{} {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.json.GetPath(keys...).MustMap()
}

// GetStringMap 获取map数据
func (c *JSONConfig) GetStringMap(keys ...string) map[string]string {
	cfg := make(map[string]string)
	err := c.Scan(keys, &cfg)
	if err != nil {
		return nil
	}
	return cfg
}

// Scan 读取配置到指定对象
func (c *JSONConfig) Scan(keys []string, value interface{}) error {
	c.m.RLock()
	defer c.m.RUnlock()

	buf, err := c.json.GetPath(keys...).MarshalJSON()
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
	c.m.Lock()
	c.json.Set(key, value)
	c.m.Unlock()
}

// SetPath 写入多级配置
func (c *JSONConfig) SetPath(keys []string, value interface{}) {
	c.m.Lock()
	c.json.SetPath(keys, value)
	c.m.Unlock()
}
