package common

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"runtime"
	"sync"
)

type Config struct {
	DB     database `toml:"database"`
	Redis  redis
	Wechat wechat
}
type database struct {
	Port, PoolNumber               int
	Server, Db, Username, Password string
}

type redis struct {
	Port             int
	Server, Password string
}

type wechat struct {
	AccessTokenPath string
}

func GetConfig() *Config {
	cfg := Config{}
	once := sync.Once{}
	once.Do(func() {
		fileName := "./config.toml"
		if runtime.GOOS == "linux" {
			fileName = "/opt/server/go/send-template-msg/config.toml"
		}
		filePath, _ := filepath.Abs(fileName)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return &cfg
}
