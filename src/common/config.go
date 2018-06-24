package common

import (
	"sync"
	"path/filepath"
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	DB    database `toml:"database"`
	Redis redis
}
type database struct {
	Port, PoolNumber               int
	Server, Db, Username, Password string
}

type redis struct {
	Port             int
	Server, Password string
}

func GetConfig() *Config {
	cfg := Config{}
	once := sync.Once{}
	once.Do(func() {
		filePath, err := filepath.Abs("./config.toml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return &cfg
}
