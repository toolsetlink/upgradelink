package svc

import (
	"github.com/redis/go-redis/v9"
	"upgradelink-admin-core/server/rpc/ent"
	"upgradelink-admin-core/server/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/logx"

	_ "upgradelink-admin-core/server/rpc/ent/runtime"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Error), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
	)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  c.RedisConf.MustNewUniversalRedis(),
	}
}
