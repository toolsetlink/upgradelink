package resource

import (
	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// Ctx 通用能力上下文
type Ctx struct {
	mysqlConn      sqlx.SqlConn
	mysqlConnCache sqlc.CachedConn
	redisConn      *redis.Redis
	localCache     *collection.Cache

	UrlModel model.UpgradeUrlModel
}

// NewCtx 创建上下文
func NewCtx(_mysql sqlx.SqlConn, _redis *redis.Redis, _mysqlCache sqlc.CachedConn, _localCache *collection.Cache, urlModel model.UpgradeUrlModel) *Ctx {
	return &Ctx{
		mysqlConn:      _mysql,
		mysqlConnCache: _mysqlCache,
		redisConn:      _redis,
		localCache:     _localCache,
		UrlModel:       urlModel,
	}
}

const PREFIX = "UPGRADE-API:"
