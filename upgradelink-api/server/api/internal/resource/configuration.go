package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyConfigurationInfoByKey             = PREFIX + "CONFIGURATION_INFO:KEY:%v"
	CacheKeyConfigurationInfoByCompanyIdAndKey = PREFIX + "CONFIGURATION_INFO:COMPANY_ID:%v:KEY:%v"
)

// GetConfigurationInfoByKey
// 获取应用信息
func (c *Ctx) GetConfigurationInfoByKey(ctx context.Context,
	key string) (*model.UpgradeConfiguration, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeConfiguration
		query := fmt.Sprintf("select * from upgrade_configuration where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeConfiguration)
		return &info, err
	}

	return nil, err
}

// GetConfigurationInfoByCompanyIdAndKey
// 获取应用信息
func (c *Ctx) GetConfigurationInfoByCompanyIdAndKey(ctx context.Context,
	companyId int64, key string) (*model.UpgradeConfiguration, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationInfoByCompanyIdAndKey, companyId, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeConfiguration
		query := fmt.Sprintf("select * from upgrade_configuration where `company_id` = ? and `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, companyId, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeConfiguration)
		return &info, err
	}

	return nil, err
}
