package svc

import (
	"time"

	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/config"
	"upgradelink-api/server/api/internal/middleware"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type ServiceContext struct {
	Config       config.Config
	Signature    rest.Middleware
	ReplayAttack rest.Middleware
	RateLimit    rest.Middleware
	CdnRateLimit rest.Middleware
	ResourceCtx  *resource.Ctx
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 1. 注册全局异常处理
	httpx.SetErrorHandlerCtx(http_handlers.ErrorHandler)

	// 2. 注册日志配置
	logc.MustSetup(c.Log)

	// 3. 注册 mysqlCli
	mysqlCli := sqlx.NewMysql(c.MysqlConfig)

	// 4. 注册 mysqlCli Cache
	mysqlCacheCli := sqlc.NewConn(mysqlCli, cache.CacheConf{
		{
			RedisConf: c.RedisConfig,
			Weight:    100,
		},
	}, cache.WithExpiry(time.Second*60))

	// 6. 注册 redisCli
	redisCli := redis.MustNewRedis(c.RedisConfig)

	// 注册 内存缓存
	localCache, _ := collection.NewCache(time.Second*60, collection.WithLimit(10000))

	// 注册 model
	UrlModel := model.NewUpgradeUrlModel(mysqlCli)

	resourceCtx := resource.NewCtx(mysqlCli, redisCli, mysqlCacheCli, localCache, UrlModel)

	return &ServiceContext{
		Config:       c,
		ResourceCtx:  resourceCtx, // 通用服务 Ctx
		Signature:    middleware.NewSignatureMiddleware(resourceCtx).Handle,
		ReplayAttack: middleware.NewReplayAttackMiddleware(resourceCtx).Handle,
		RateLimit:    middleware.NewRateLimitMiddleware(resourceCtx).Handle,
		CdnRateLimit: middleware.NewCdnRateLimitMiddleware(resourceCtx).Handle,
	}
}
