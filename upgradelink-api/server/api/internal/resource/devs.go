package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyDevInfoByKey = PREFIX + "DEV_INFO:KEY:%v"
)

// GetDevInfoByKey
// 通过设备 key 获取设备信息
func (c *Ctx) GetDevInfoByKey(ctx context.Context, key string) (*model.UpgradeDevs, error) {

	cacheKey := fmt.Sprintf(CacheKeyDevInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeDevs
		query := fmt.Sprintf("select * from upgrade_devs where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeDevs)
		return &info, err
	}

	return nil, err
}
