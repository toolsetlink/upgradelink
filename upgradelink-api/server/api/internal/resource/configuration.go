package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyConfigurationInfoByKey = PREFIX + "CONFIGURATION_INFO:KEY:%v"
)

// GetConfigurationInfoByKey
// 通过 Configuration key获取 Configuration 信息
func (c *Ctx) GetConfigurationInfoByKey(ctx context.Context,
	key string) (*model.UpgradeConfiguration, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var configurationInfo model.UpgradeConfiguration
		query := fmt.Sprintf("select * from upgrade_configuration where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &configurationInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return configurationInfo, nil
	})

	if v != nil {
		configurationInfo := v.(model.UpgradeConfiguration)
		return &configurationInfo, err
	}

	return nil, err
}
