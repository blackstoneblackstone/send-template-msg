package common

import (
	"sync"
	"path/filepath"
	"github.com/BurntSushi/toml"
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
		filePath, _ := filepath.Abs("../../config.toml")
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return &cfg
}
