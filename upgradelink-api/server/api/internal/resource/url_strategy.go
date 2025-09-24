package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyUrlStrategyInfoByDevTypeAllAndUrlIdAndVersion = PREFIX + "URL_STRATEGY_INFO:DEV_TYPE:ALL:URL_ID:%v:CLIENT_VERSION_CODE:%v:"
	CacheKeyUrlStrategyListByUrlIdAndVersion              = PREFIX + "URL_STRATEGY_LIST:URL_ID:%v:CLIENT_VERSION_CODE:%v:"
)

// GetUrlStrategyInfoByUrlIdAndVersionAndDevTypeAll
// 获取大于 客户端的 versionCode 的版本 获取开启了 全部设备的策略
func (c *Ctx) GetUrlStrategyInfoByUrlIdAndVersionAndDevTypeAll(ctx context.Context, urlId int64, clientVersionCode int64) (*model.UpgradeUrlUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlStrategyInfoByDevTypeAllAndUrlIdAndVersion, urlId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var urlStrategyInfo model.UpgradeUrlUpgradeStrategy
		query := fmt.Sprintf("select upgrade_url_upgrade_strategy.* from upgrade_url_upgrade_strategy " +
			"left join upgrade_url_version on upgrade_url_upgrade_strategy.url_version_id = upgrade_url_version.id " +
			"where upgrade_url_upgrade_strategy.url_id = ? " +
			"AND upgrade_url_version.version_code > ? " +
			"AND ? > upgrade_url_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_url_upgrade_strategy.end_datetime " +
			"AND upgrade_url_upgrade_strategy.upgrade_dev_type = 0 " +
			"AND upgrade_url_upgrade_strategy.enable = 1 " +
			"AND upgrade_url_upgrade_strategy.is_del = 0 " +
			"order by upgrade_url_version.version_code desc limit 1")

		//fmt.Println("query: ", query)
		err := c.mysqlConnCache.QueryRowCtx(ctx, &urlStrategyInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, urlId, clientVersionCode, now, now)
		})

		if err != nil {
			return nil, err
		}

		return urlStrategyInfo, nil
	})

	if v != nil {
		urlStrategyInfo := v.(model.UpgradeUrlUpgradeStrategy)
		return &urlStrategyInfo, err
	}
	return nil, err
}

// GetUrlStrategyListByUrlIdAndVersion
// 获取大于 客户端的 versionCode 的版本 的全部策略 list
func (c *Ctx) GetUrlStrategyListByUrlIdAndVersion(ctx context.Context, urlId int64, clientVersionCode int64) ([]*model.UpgradeUrlUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlStrategyListByUrlIdAndVersion, urlId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var urlStrategyList []*model.UpgradeUrlUpgradeStrategy
		query := fmt.Sprintf("select upgrade_url_upgrade_strategy.* from upgrade_url_upgrade_strategy " +
			"left join upgrade_url_version on upgrade_url_upgrade_strategy.url_version_id = upgrade_url_version.id " +
			"where upgrade_url_upgrade_strategy.url_id = ? " +
			"AND upgrade_url_version.version_code > ? " +
			"AND ? >= upgrade_url_upgrade_strategy.begin_datetime " +
			"AND ? <= upgrade_url_upgrade_strategy.end_datetime " +
			"AND upgrade_url_upgrade_strategy.enable = 1 " +
			"AND upgrade_url_upgrade_strategy.is_del = 0 " +
			"order by upgrade_url_version.version_code desc")
		err := c.mysqlConn.QueryRowsCtx(context.Background(), &urlStrategyList, query, urlId, clientVersionCode, now, now)
		if err != nil {
			return nil, err
		}

		return urlStrategyList, nil
	})

	if v != nil {
		urlStrategyList := v.([]*model.UpgradeUrlUpgradeStrategy)
		return urlStrategyList, err
	}

	return nil, err
}
