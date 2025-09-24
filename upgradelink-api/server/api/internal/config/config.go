package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MysqlConfig string          `json:",optional,inherit"`
	RedisConfig redis.RedisConf `json:",optional,inherit"`
}

const (
	// UIDField uid header key
	UIDField = "UID"
	// SsoIDField sso id header key
	SsoIDField = "SsoID"
	// SignatureField  sign header key
	SignatureField = "Signature"
	// TkField tk header key
	TkField = "Tk"
	// SignatureExpired sign 过期时间
	SignatureExpired = 50000 // 5分钟
)
