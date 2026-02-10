package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MysqlConf  string          `json:",optional,inherit"`
	RedisConf  redis.RedisConf `json:",optional,inherit"`
	UploadConf UploadConf
}

type UploadConf struct {
	Bucket    string `json:",optional"`
	SecretID  string `json:",optional"`
	SecretKey string `json:",optional"`
	Endpoint  string `json:",optional"`
	Folder    string `json:",optional"`
	Region    string `json:",optional"`
	CdnUrl    string `json:",optional"`
}
