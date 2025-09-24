package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyUrlInfoByKey = PREFIX + "URL_INFO:KEY:%v"
)

// GetUrlInfoByKey
// 通过 url key获取 url 信息
func (c *Ctx) GetUrlInfoByKey(ctx context.Context,
	key string) (*model.UpgradeUrl, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var urlInfo model.UpgradeUrl
		query := fmt.Sprintf("select * from upgrade_url where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &urlInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return urlInfo, nil
	})

	if v != nil {
		urlInfo := v.(model.UpgradeUrl)
		return &urlInfo, err
	}

	return nil, err
}
