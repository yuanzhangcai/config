package encoder

// Encoder 编码器
type Encoder interface {
	LoadFile(file string) (map[string]interface{}, error)
}
