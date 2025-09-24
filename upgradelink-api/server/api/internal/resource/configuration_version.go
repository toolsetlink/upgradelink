package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyConfigurationVersionListByConfigurationIdAndVersionCode = PREFIX + "CONFIGURATION_VERSION_LIST:CONFIGURATION_ID:%v:VERSION_CODE:%v"
	CacheKeyConfigurationVersionInfoById                            = PREFIX + "CONFIGURATION_VERSION_INFO:ID:%v"
	CacheKeyConfigurationVersionLastInfoByConfigurationId           = PREFIX + "CONFIGURATION_VERSION_LAST_INFO:CONFIGURATION_ID:%v"
	CacheKeyConfigurationVersionInfoByConfigurationIdAndVersionCode = PREFIX + "CONFIGURATION_VERSION_INFO:CONFIGURATION_ID:%v:VERSION_CODE:%v"
)

// GetConfigurationVersionListByConfigurationIdAndVersionCode
// 获取大于 versionCode 的版本列表
func (c *Ctx) GetConfigurationVersionListByConfigurationIdAndVersionCode(ctx context.Context, configurationId int64, versionCode int64) ([]*model.UpgradeConfigurationVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationVersionListByConfigurationIdAndVersionCode, configurationId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var ConfigurationVersionList []*model.UpgradeConfigurationVersion
		query := fmt.Sprintf("select * from upgrade_configuration_version where `configuration_id` = ? and is_del = 0 and `version_code` > ? order by version_code desc ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &ConfigurationVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &ConfigurationVersionList, query, configurationId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return ConfigurationVersionList, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeConfigurationVersion)
		return list, err
	}

	return nil, err
}

// GetConfigurationVersionLastInfoByConfigurationId
// 获取最新版本的信息
func (c *Ctx) GetConfigurationVersionLastInfoByConfigurationId(ctx context.Context, configurationId int64) (*model.UpgradeConfigurationVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationVersionLastInfoByConfigurationId, configurationId)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeConfigurationVersion
		query := fmt.Sprintf("select * from upgrade_configuration_version where `configuration_id` = ? and is_del = 0 order by version_code desc limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, configurationId)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeConfigurationVersion)
		return &info, err
	}

	return nil, err
}

// 获取固定版本的信息
func (c *Ctx) GetConfigurationVersionInfoByConfigurationIdAndVersionCode(ctx context.Context, configurationId int64, versionCode int64) (*model.UpgradeConfigurationVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationVersionInfoByConfigurationIdAndVersionCode, configurationId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeConfigurationVersion
		query := fmt.Sprintf("select * from upgrade_configuration_version where `configuration_id` = ? and is_del = 0 and `version_code` = ? LIMIT 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, configurationId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeConfigurationVersion)
		return &info, err
	}

	return nil, err
}

// GetConfigurationVersionInfoById
// 通过 Configuration version id 获取信息
func (c *Ctx) GetConfigurationVersionInfoById(ctx context.Context,
	id int64) (*model.UpgradeConfigurationVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationVersionInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeConfigurationVersion
		query := fmt.Sprintf("select * from upgrade_configuration_version where `id` = ?  limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})
		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeConfigurationVersion)
		return &info, err
	}

	return nil, err
}
