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
	AccessKey    rest.Middleware
	ReplayAttack rest.Middleware
	RateLimit    rest.Middleware
	CdnRateLimit rest.Middleware
	Rule         rest.Middleware
	ResourceCtx  *resource.Ctx
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 缓存配置常量定义
	// Redis 作为 MySQL 查询缓存的权重，用于多缓存源场景
	const redisCacheWeight = 100
	// MySQL 查询结果缓存过期时间
	const mysqlCacheExpiry = time.Second * 60
	// 本地内存缓存过期时间
	const localCacheExpiry = time.Second * 60
	// 本地内存缓存最大条目数限制
	const localCacheMaxItems = 10000

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
			Weight:    redisCacheWeight,
		},
	}, cache.WithExpiry(mysqlCacheExpiry))

	// 6. 注册 redisCli
	redisCli := redis.MustNewRedis(c.RedisConfig)

	// 注册 内存缓存
	localCache, _ := collection.NewCache(localCacheExpiry, collection.WithLimit(localCacheMaxItems))

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
