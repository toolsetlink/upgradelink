package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyConfigurationStrategyInfoByDevTypeAllAndConfigurationIdAndVersion = PREFIX + "CONFIGURATION_STRATEGY_INFO:DEV_TYPE:ALL:CONFIGURATION_ID:%v:CLIENT_VERSION_CODE:%v:"
	CacheKeyConfigurationStrategyListByConfigurationIdAndVersion              = PREFIX + "CONFIGURATION_STRATEGY_LIST:CONFIGURATION_ID:%v:CLIENT_VERSION_CODE:%v:"
)

// GetConfigurationStrategyInfoByConfigurationIdAndVersionAndDevTypeAll
// 获取大于 客户端的 versionCode 的版本 获取开启了 全部设备的策略
func (c *Ctx) GetConfigurationStrategyInfoByConfigurationIdAndVersionAndDevTypeAll(ctx context.Context, configurationId int64, clientVersionCode int64) (*model.UpgradeConfigurationUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationStrategyInfoByDevTypeAllAndConfigurationIdAndVersion, configurationId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var configurationStrategyInfo model.UpgradeConfigurationUpgradeStrategy
		query := fmt.Sprintf("select upgrade_configuration_upgrade_strategy.* from upgrade_configuration_upgrade_strategy " +
			"left join upgrade_configuration_version on upgrade_configuration_upgrade_strategy.configuration_version_id = upgrade_configuration_version.id " +
			"where upgrade_configuration_upgrade_strategy.configuration_id = ? " +
			"AND upgrade_configuration_version.version_code > ? " +
			"AND ? > upgrade_configuration_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_configuration_upgrade_strategy.end_datetime " +
			"AND upgrade_configuration_upgrade_strategy.upgrade_dev_type = 0 " +
			"AND upgrade_configuration_upgrade_strategy.enable = 1 " +
			"AND upgrade_configuration_upgrade_strategy.is_del = 0 " +
			"order by upgrade_configuration_version.version_code desc limit 1")

		//fmt.Println("query: ", query)
		err := c.mysqlConnCache.QueryRowCtx(ctx, &configurationStrategyInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, configurationId, clientVersionCode, now, now)
		})

		if err != nil {
			return nil, err
		}

		return configurationStrategyInfo, nil
	})

	if v != nil {
		configurationStrategyInfo := v.(model.UpgradeConfigurationUpgradeStrategy)
		return &configurationStrategyInfo, err
	}
	return nil, err
}

// GetConfigurationStrategyListByConfigurationIdAndVersion
// 获取大于 客户端的 versionCode 的版本 的全部策略 list
func (c *Ctx) GetConfigurationStrategyListByConfigurationIdAndVersion(ctx context.Context, configurationId int64, clientVersionCode int64) ([]*model.UpgradeConfigurationUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyConfigurationStrategyListByConfigurationIdAndVersion, configurationId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var configurationStrategyList []*model.UpgradeConfigurationUpgradeStrategy
		query := fmt.Sprintf("select upgrade_configuration_upgrade_strategy.* from upgrade_configuration_upgrade_strategy " +
			"left join upgrade_configuration_version on upgrade_configuration_upgrade_strategy.configuration_version_id = upgrade_configuration_version.id " +
			"where upgrade_configuration_upgrade_strategy.configuration_id = ? " +
			"AND upgrade_configuration_version.version_code > ? " +
			"AND ? >= upgrade_configuration_upgrade_strategy.begin_datetime " +
			"AND ? <= upgrade_configuration_upgrade_strategy.end_datetime " +
			"AND upgrade_configuration_upgrade_strategy.enable = 1 " +
			"AND upgrade_configuration_upgrade_strategy.is_del = 0 " +
			"order by upgrade_configuration_version.version_code desc")
		err := c.mysqlConn.QueryRowsCtx(context.Background(), &configurationStrategyList, query, configurationId, clientVersionCode, now, now)
		if err != nil {
			return nil, err
		}

		return configurationStrategyList, nil
	})

	if v != nil {
		configurationStrategyList := v.([]*model.UpgradeConfigurationUpgradeStrategy)
		return configurationStrategyList, err
	}

	return nil, err
}
