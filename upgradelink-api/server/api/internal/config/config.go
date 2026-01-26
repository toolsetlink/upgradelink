package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MysqlConf string          `json:",optional,inherit"`
	RedisConf redis.RedisConf `json:",optional,inherit"`
}
