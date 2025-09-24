package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyDevModelInfoById = PREFIX + "DEV_MODEL_INFO:ID:%v"
)

// GetDevModelInfoById
// 通过获取详细信息
func (c *Ctx) GetDevModelInfoById(ctx context.Context, id int64) (*model.UpgradeDevModel, error) {

	cacheKey := fmt.Sprintf(CacheKeyDevModelInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeDevModel
		query := fmt.Sprintf("select * from upgrade_dev_model where `id` = ?  limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeDevModel)
		return &info, err
	}

	return nil, err
}
