package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var configFilePath = "/Users/zacyuan/MyWork/config/examples/"

func init() {
	_ = LoadFile(configFilePath + "config.toml")
}

func TestConfig(t *testing.T) {
	cfg := New()
	err := cfg.LoadFile(configFilePath + "config.toml")
	assert.Nil(t, err)

	err = cfg.LoadFile(configFilePath + "config.json")
	assert.Nil(t, err)

	err = cfg.LoadFile(configFilePath + "config.yaml")
	assert.Nil(t, err)
}

func TestNew(t *testing.T) {
	cfg := New()
	assert.NotNil(t, cfg)
}

func TestGet(t *testing.T) {
	t.Run("Get have key", func(t *testing.T) {
		value := Get("app_desc")
		assert.NotNil(t, value)
	})

	t.Run("Get key is not exist", func(t *testing.T) {
		value := Get("no_key")
		assert.Nil(t, value)
	})
}

func TestGetString(t *testing.T) {
	t.Run("GetString have key", func(t *testing.T) {
		value := GetString("app_desc")
		assert.Equal(t, "toml_cfg", value)
	})

	t.Run("GetString key is not exist", func(t *testing.T) {
		value := GetString("no_key")
		assert.Empty(t, value)
	})
}

func TestGetStringArray(t *testing.T) {
	t.Run("GetStringArray have key", func(t *testing.T) {
		value := GetStringArray("str_array")
		assert.Equal(t, "aa", value[0])
		assert.Equal(t, "bb", value[1])
	})

	t.Run("GetStringArray key is not exist", func(t *testing.T) {
		value := GetStringArray("no_key")
		assert.Empty(t, value)
	})
}

func TestGetBool(t *testing.T) {
	t.Run("GetBool true", func(t *testing.T) {
		assert.True(t, GetBool("bool_true"))
	})

	t.Run("GetBool false", func(t *testing.T) {
		assert.False(t, GetBool("bool_false"))
	})

	t.Run("GetBool key is not exist", func(t *testing.T) {
		assert.False(t, GetBool("no_key"))
	})
}

func TestGetInt(t *testing.T) {
	t.Run("GetInt 13", func(t *testing.T) {
		assert.Equal(t, 13, GetInt("int_13"))
	})

	t.Run("GetInt -1", func(t *testing.T) {
		assert.Equal(t, -1, GetInt("int__1"))
	})

	t.Run("GetInt key is not exist", func(t *testing.T) {
		assert.Equal(t, 0, GetInt("no_key"))
	})
}

func TestGetInt64(t *testing.T) {
	t.Run("GetInt64 13", func(t *testing.T) {
		assert.Equal(t, int64(13), GetInt64("int_13"))
	})

	t.Run("GetInt64 -1", func(t *testing.T) {
		assert.Equal(t, int64(-1), GetInt64("int__1"))
	})

	t.Run("GetInt64 key is not exist", func(t *testing.T) {
		assert.Equal(t, int64(0), GetInt64("no_key"))
	})
}

func TestGetUint64(t *testing.T) {
	t.Run("GetUint64 13", func(t *testing.T) {
		assert.Equal(t, uint64(13), GetUint64("int_13"))
	})

	t.Run("GetUint64 -1", func(t *testing.T) {
		assert.Equal(t, uint64(0), GetUint64("int__1"))
	})

	t.Run("GetUint64 key is not exist", func(t *testing.T) {
		assert.Equal(t, uint64(0), GetUint64("no_key"))
	})
}

func TestGetFloat64(t *testing.T) {
	t.Run("GetFloat64 45.3", func(t *testing.T) {
		assert.Equal(t, float64(45.3), GetFloat64("float64"))
	})

	t.Run("GetFloat64 -1", func(t *testing.T) {
		assert.Equal(t, float64(-1), GetFloat64("int__1"))
	})

	t.Run("GetFloat64 key is not exist", func(t *testing.T) {
		assert.Equal(t, float64(0), GetFloat64("no_key"))
	})
}

func TestGetArray(t *testing.T) {
	t.Run("GetArray have key", func(t *testing.T) {
		value := GetArray("str_array")
		assert.Equal(t, 2, len(value))
	})

	t.Run("GetArray key is not exist", func(t *testing.T) {
		value := GetArray("no_key")
		assert.Equal(t, 0, len(value))
	})
}

func TestGetMap(t *testing.T) {
	t.Run("GetMap have key", func(t *testing.T) {
		value := GetMap("common")
		assert.NotNil(t, value)
		assert.Equal(t, "toml", value["config_toml"].(string))
	})

	t.Run("GetMap key is not exist", func(t *testing.T) {
		value := GetMap("no_key")
		assert.Nil(t, value)
	})
}

func TestScan(t *testing.T) {
	value := struct {
		List   []string `json:"list"`
		Server string   `json:"server"`
		Count  int      `json:"count"`
	}{}
	err := Scan([]string{"db"}, &value)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(value.List))
	assert.Equal(t, "127.0.0.1:3306", value.Server)
	assert.Equal(t, 100, value.Count)
}

func TestSet(t *testing.T) {
	Set("set_key", "value")
	assert.Equal(t, "value", GetString("set_key"))
}

func TestSetPath(t *testing.T) {
	SetPath([]string{"common", "set_key"}, "value")
	assert.Equal(t, "value", GetString("common", "set_key"))
}

func TestLoadOsEnv(t *testing.T) {
	os.Setenv("test_env", "dev")
	err := LoadOsEnv()
	assert.Nil(t, err)
	assert.Equal(t, "dev", GetString("test_env"))
}

func TestLoadMemory(t *testing.T) {
	t.Run("LoadMemory success json", func(t *testing.T) {
		str := `{
			"name" : "aaa",
			"db" : {
				"list" : ["db1", "db2"],
				"db1" : "127.0.0.1"
			}
		}`
		err := LoadMemory(str, "json")
		assert.Nil(t, err)
		assert.Equal(t, "aaa", GetString("name"))
		assert.Equal(t, 2, len(GetStringArray("db", "list")))
	})

	t.Run("LoadMemory success toml", func(t *testing.T) {
		str := `
		name = "aaa"
		[db]
		list = ["db1", "db2"]
		db1 = "127.0.0.1"
		`
		err := LoadMemory(str, "toml")
		assert.Nil(t, err)
		assert.Equal(t, "aaa", GetString("name"))
		assert.Equal(t, 2, len(GetStringArray("db", "list")))
	})

	t.Run("LoadMemory unsupported", func(t *testing.T) {
		str := `{
			"name" : "aaa",
			"db" : {
				"list" : ["db1", "db2"],
				"db1" : "127.0.0.1"
			}
		}`
		err := LoadMemory(str, "ini")
		assert.Equal(t, ErrUnsupportedFileFormat, err)
	})

	t.Run("LoadMemory error", func(t *testing.T) {
		str := `
		name = "aaa"
		[db]
		list = ["db1", "db2"]
		db1 = "127.0.0.1"
		`
		err := LoadMemory(str, "json")
		assert.NotNil(t, err)
	})
}

// func loadConfig(file string) error {

// 	if _, err := os.Stat(file); err != nil {
// 		return err
// 	}

// 	config := viper.New()
// 	config.SetConfigFile(file)
// 	err := config.ReadInConfig()
// 	if err != nil {
// 		return err
// 	}

// 	allKeys := config.AllKeys()
// 	for _, one := range allKeys {
// 		//fmt.Println(one)
// 		viper.SetDefault(one, config.Get(one))
// 	}

// 	return nil
// }

// func BenchmarkGet(b *testing.B) {
// 	_ = LoadFile(configFilePath + "config.toml")
// 	_ = LoadFile(configFilePath + "config.json")

// 	_ = loadConfig(configFilePath + "config.toml")
// 	_ = loadConfig(configFilePath + "config.json")

// 	_ = config.LoadFile(configFilePath + "config.toml")
// 	_ = config.LoadFile(configFilePath + "config.json")

// 	beego.LoadAppConfig("toml", configFilePath+"config.toml")

// 	fmt.Println(viper.Get("db.server"))
// 	fmt.Println(config.Get("db", "server").String(""))
// 	fmt.Println(GetString("db", "server"))

// 	beego.LoadAppConfig("ini", configFilePath+"config.ini")
// 	fmt.Println(beego.AppConfig.String("db"))

// 	b.Run("Get             ", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			value := GetString("db", "server")
// 			_ = value
// 		}
// 	})

// 	b.Run("viper Get         ", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			value := viper.Get("db.server")
// 			_ = value
// 		}
// 	})

// 	b.Run("micro config Get    ", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			value := config.Get("db", "server").String("")
// 			_ = value
// 		}
// 	})

// 	b.Run("beego config Get    ", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			value := beego.AppConfig.String("db")
// 			_ = value
// 		}
// 	})

// 	b.Run("Parallel Get       ", func(b *testing.B) {
// 		b.RunParallel(func(pb *testing.PB) {
// 			for pb.Next() {
// 				value := GetString("db", "server")
// 				_ = value
// 			}
// 		})
// 	})

// 	b.Run("Parallel viper Get", func(b *testing.B) {
// 		b.RunParallel(func(pb *testing.PB) {
// 			for pb.Next() {
// 				value := viper.Get("db.server")
// 				_ = value
// 			}
// 		})
// 	})

// 	b.Run("Parallel micro Get", func(b *testing.B) {
// 		b.RunParallel(func(pb *testing.PB) {
// 			for pb.Next() {
// 				value := config.Get("db", "server").String("")
// 				_ = value
// 			}
// 		})
// 	})

// 	b.Run("Parallel beego Get", func(b *testing.B) {
// 		b.RunParallel(func(pb *testing.PB) {
// 			for pb.Next() {
// 				value := beego.AppConfig.String("db")
// 				_ = value
// 			}
// 		})
// 	})
// }
