package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyUrlInfoByKey             = PREFIX + "URL_INFO:KEY:%v"
	CacheKeyUrlInfoByCompanyIdAndKey = PREFIX + "URL_INFO:COMPANY_ID:%v:KEY:%v"
)

// GetUrlInfoByKey
// 获取应用信息
func (c *Ctx) GetUrlInfoByKey(ctx context.Context,
	key string) (*model.UpgradeUrl, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeUrl
		query := fmt.Sprintf("select * from upgrade_url where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeUrl)
		return &info, err
	}

	return nil, err
}

// GetUrlInfoByCompanyIdAndKey
// 获取应用信息
func (c *Ctx) GetUrlInfoByCompanyIdAndKey(ctx context.Context,
	companyId int64, key string) (*model.UpgradeUrl, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlInfoByCompanyIdAndKey, companyId, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeUrl
		query := fmt.Sprintf("select * from upgrade_url where `company_id` = ? and `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, companyId, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeUrl)
		return &info, err
	}

	return nil, err
}
