package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyFileInfoByKey = PREFIX + "FILE_INFO:KEY:%v"
)

func (c *Ctx) GetFileInfoByKey(ctx context.Context,
	key string) (*model.UpgradeFile, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var urlInfo model.UpgradeFile
		query := fmt.Sprintf("select * from upgrade_file where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &urlInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return urlInfo, nil
	})

	if v != nil {
		urlInfo := v.(model.UpgradeFile)
		return &urlInfo, err
	}

	return nil, err
}
