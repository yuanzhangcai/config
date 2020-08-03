# config
[![build](https://github.com/yuanzhangcai/config/workflows/Go/badge.svg)](https://github.com/yuanzhangcai/config/actions)
[![codecov](https://codecov.io/gh/yuanzhangcai/config/branch/master/graph/badge.svg)](https://codecov.io/gh/yuanzhangcai/config)
[![GitHub](https://img.shields.io/github/license/yuanzhangcai/config)](https://github.com/yuanzhangcai/config/blob/master/LICENSE)

配置文件解析，目前支持toml、yaml、json文件格式。支持多配置文件重载。默认初次加载所有环境变量。支持配置文件热更新。

example:
```
    import github.com/yuanzhangcai/config

    // 加载配置
    config.LoadFile(configFilePath + "config.toml")

    // 读取配置参数
    value := config.GetString("key")

    // 修改配置参数
    config.Set("newKey", "value")

    value := strut {
      List   []string `json:"list"`
      Server string   `json:"server"`
      Count  int      `json:"count"`
    }{}
    err = Scan([]string{"db"}, &value)

```

性能压测：
```
goos: darwin
goarch: amd64
pkg: github.com/yuanzhangcai/config
BenchmarkGet/Get_______________-12         	 6036589	       196 ns/op
BenchmarkGet/viper_Get_________-12       	  462037	      2604 ns/op
BenchmarkGet/micro_config_Get__-12      	 4340478	       275 ns/op
BenchmarkGet/beego_config_Get__-12      	 2691524	       446 ns/op
BenchmarkGet/Parallel_Get______-12      	18792312	        68.7 ns/op
BenchmarkGet/Parallel_viper_Get-12       	 1757396	       668 ns/op
BenchmarkGet/Parallel_micro_Get-12       	13200244	        87.0 ns/op
BenchmarkGet/Parallel_beego_Get-12       	 9161379	       149 ns/op
```