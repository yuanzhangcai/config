package config

// Loader 加载器接口
type Loader interface {
	LoadMemory(config, t string) error // 加载内存配置
	LoadFile(file string) error        // 加载配置文件
	LoadOsEnv() error                  // 加载环境变量
}

// Getter 获取接口
type Getter interface {
	Get(keys ...string) interface{}
	GetString(keys ...string) string
	GetStringArray(keys ...string) []string
	GetBool(keys ...string) bool
	GetInt(keys ...string) int
	GetInt64(keys ...string) int64
	GetUint64(keys ...string) uint64
	GetFloat64(keys ...string) float64
	GetArray(keys ...string) []interface{}
	GetMap(keys ...string) map[string]interface{}
	GetStringMap(keys ...string) map[string]string
	Scan(keys []string, value interface{}) error
}

// Setter 设置接口
type Setter interface {
	Set(key string, value interface{})
	SetPath(keys []string, value interface{})
}

// Config 配置结构体
type Config interface {
	Loader
	Getter
	Setter
}

var defaultConfig = newJSONConfig()

// New 创建配置文件
func New() Config {
	return newJSONConfig()
}

// LoadFile 加载配置文件
func LoadFile(file string) error {
	return defaultConfig.LoadFile(file)
}

// Get 获取interface配置
func Get(keys ...string) interface{} {
	return defaultConfig.Get(keys...)
}

// GetString 获取配置
func GetString(keys ...string) string {
	return defaultConfig.GetString(keys...)
}

// GetStringArray 获取map数据
func GetStringArray(keys ...string) []string {
	return defaultConfig.GetStringArray(keys...)
}

// GetBool 获取bool配置
func GetBool(keys ...string) bool {
	return defaultConfig.GetBool(keys...)
}

// GetInt 获取int配置
func GetInt(keys ...string) int {
	return defaultConfig.GetInt(keys...)
}

// GetInt64 获取int64配置
func GetInt64(keys ...string) int64 {
	return defaultConfig.GetInt64(keys...)
}

// GetUint64 获取uint64配置
func GetUint64(keys ...string) uint64 {
	return defaultConfig.GetUint64(keys...)
}

// GetFloat64 获取float64配置
func GetFloat64(keys ...string) float64 {
	return defaultConfig.GetFloat64(keys...)
}

// GetArray 获取interface数组
func GetArray(keys ...string) []interface{} {
	return defaultConfig.GetArray(keys...)
}

// GetMap 获取map数据
func GetMap(keys ...string) map[string]interface{} {
	return defaultConfig.GetMap(keys...)
}

// GetStringMap 获取string map数据
func GetStringMap(keys ...string) map[string]string {
	return defaultConfig.GetStringMap(keys...)
}

// Scan 读取配置到指定对象
func Scan(keys []string, value interface{}) error {
	return defaultConfig.Scan(keys, value)
}

// Set 写入配置
func Set(key string, value interface{}) {
	defaultConfig.Set(key, value)
}

// SetPath 写入多级配置
func SetPath(keys []string, value interface{}) {
	defaultConfig.SetPath(keys, value)
}

// LoadOsEnv 加载环境变量
func LoadOsEnv() error {
	return defaultConfig.LoadOsEnv()
}

// LoadMemory 从内存中加载配置文件
func LoadMemory(config, t string) error {
	return defaultConfig.LoadMemory(config, t)
}
