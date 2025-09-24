package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyTauriInfoByKey = PREFIX + "TAURI_INFO:KEY:%v"
)

func (c *Ctx) GetTauriInfoByKey(ctx context.Context,
	key string) (*model.UpgradeTauri, error) {

	cacheKey := fmt.Sprintf(CacheKeyTauriInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeTauri
		query := fmt.Sprintf("select * from upgrade_tauri where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeTauri)
		return &info, err
	}

	return nil, err
}
