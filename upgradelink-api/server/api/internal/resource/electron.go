package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyElectronInfoByKey = PREFIX + "ELECTRON_INFO:KEY:%v"
)

func (c *Ctx) GetElectronInfoByKey(ctx context.Context,
	key string) (*model.UpgradeElectron, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeElectron
		query := fmt.Sprintf("select * from upgrade_electron where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeElectron)
		return &info, err
	}

	return nil, err
}
