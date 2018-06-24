package dbServer

import (
	"common"
	"github.com/gomodule/redigo/redis"
	"fmt"
)

type RedisApi struct {
	Config  *common.Config
	connect redis.Conn
}

func (redisApi *RedisApi) Connect() error {
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", redisApi.Config.Redis.Server, redisApi.Config.Redis.Port),
		redis.DialPassword(redisApi.Config.Redis.Password))
	redisApi.connect = c
	return err
}

func (redisApi *RedisApi) Set(k string, v interface{}) (string, error) {
	return redis.String(redisApi.connect.Do("SET", k, v))
}
func (redisApi *RedisApi) Get(k string) (string, error) {
	return redis.String(redisApi.connect.Do("GET", k))
}
