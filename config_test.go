package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var configFilePath = "/Users/zacyuan/MyWork/config/examples/"

func loadConfig(file string) error {

	if _, err := os.Stat(file); err != nil {
		return err
	}

	config := viper.New()
	config.SetConfigFile(file)
	err := config.ReadInConfig()
	if err != nil {
		return err
	}

	allKeys := config.AllKeys()
	for _, one := range allKeys {
		//fmt.Println(one)
		viper.SetDefault(one, config.Get(one))
	}

	return nil
}

func TestConfig(t *testing.T) {
	cfg := New()
	err := cfg.LoadFile(configFilePath + "config.toml")
	assert.Nil(t, err)

	err = cfg.LoadFile(configFilePath + "config.json")
	assert.Nil(t, err)

	err = cfg.LoadFile(configFilePath + "config.yaml")
	assert.Nil(t, err)

	fmt.Println(cfg.GetString("app_desc"))
	fmt.Println(cfg.GetString("common", "config_json"))
	fmt.Println(cfg.GetString("common", "config_yaml"))
	fmt.Println(cfg.GetString("common", "config_toml"))

}

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
