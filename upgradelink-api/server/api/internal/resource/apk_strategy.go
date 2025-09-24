package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyApkStrategyInfoByDevTypeAllAndApkIdAndVersion = PREFIX + "APK_STRATEGY_INFO:DEV_TYPE:ALL:APK_ID:%v:CLIENT_VERSION_CODE:%v:"
	CacheKeyApkStrategyListByApkIdAndVersion              = PREFIX + "APK_STRATEGY_LIST:APK_ID:%v:CLIENT_VERSION_CODE:%v:"

	// CacheKeyApkStrategyLastInfoByApkId 最新版本信息
	CacheKeyApkStrategyLastInfoByApkId = PREFIX + "APK_STRATEGY_LAST_INFO:APK_ID:%v"
)

// GetApkStrategyInfoByApkIdAndVersionAndDevTypeAll
// 获取大于 客户端的 versionCode 的版本 获取开启了 全部设备的策略
func (c *Ctx) GetApkStrategyInfoByApkIdAndVersionAndDevTypeAll(ctx context.Context, apkId int64, clientVersionCode int64) (*model.UpgradeApkUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkStrategyInfoByDevTypeAllAndApkIdAndVersion, apkId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var apkStrategyInfo model.UpgradeApkUpgradeStrategy
		query := fmt.Sprintf("select upgrade_apk_upgrade_strategy.* from upgrade_apk_upgrade_strategy " +
			"left join upgrade_apk_version on upgrade_apk_upgrade_strategy.apk_version_id = upgrade_apk_version.id " +
			"where upgrade_apk_upgrade_strategy.apk_id = ? " +
			"AND upgrade_apk_version.version_code > ? " +
			"AND ? > upgrade_apk_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_apk_upgrade_strategy.end_datetime " +
			"AND upgrade_apk_upgrade_strategy.upgrade_dev_type = 0 " +
			"AND upgrade_apk_upgrade_strategy.enable = 1 " +
			"AND upgrade_apk_upgrade_strategy.is_del = 0 " +
			"order by upgrade_apk_version.version_code desc limit 1")

		//fmt.Println("query: ", query)
		err := c.mysqlConnCache.QueryRowCtx(ctx, &apkStrategyInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, apkId, clientVersionCode, now, now)
		})

		if err != nil {
			return nil, err
		}

		return apkStrategyInfo, nil
	})

	if v != nil {
		apkStrategyInfo := v.(model.UpgradeApkUpgradeStrategy)
		return &apkStrategyInfo, err
	}
	return nil, err
}

// GetApkStrategyListByApkIdAndVersion
// 获取大于 客户端的 versionCode 的版本 的全部策略 list
func (c *Ctx) GetApkStrategyListByApkIdAndVersion(ctx context.Context, apkId int64, clientVersionCode int64) ([]*model.UpgradeApkUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkStrategyListByApkIdAndVersion, apkId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var apkStrategyList []*model.UpgradeApkUpgradeStrategy
		query := fmt.Sprintf("select upgrade_apk_upgrade_strategy.* from upgrade_apk_upgrade_strategy " +
			"left join upgrade_apk_version on upgrade_apk_upgrade_strategy.apk_version_id = upgrade_apk_version.id " +
			"where upgrade_apk_upgrade_strategy.apk_id = ? " +
			"AND upgrade_apk_version.version_code > ? " +
			"AND ? > upgrade_apk_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_apk_upgrade_strategy.end_datetime " +
			"AND upgrade_apk_upgrade_strategy.enable = 1 " +
			"AND upgrade_apk_upgrade_strategy.is_del = 0 " +
			"order by upgrade_apk_version.version_code desc")
		fmt.Println("query: ", query)
		err := c.mysqlConn.QueryRowsCtx(context.Background(), &apkStrategyList, query, apkId, clientVersionCode, now, now)
		if err != nil {
			return nil, err
		}

		fmt.Println("GetApkStrategyListByApkIdAndVersion apkStrategyList: ", len(apkStrategyList))
		return apkStrategyList, nil
	})

	if v != nil {
		apkStrategyList := v.([]*model.UpgradeApkUpgradeStrategy)
		return apkStrategyList, err
	}

	return nil, err
}

// GetLastApkStrategyInfoByApkIdAndVersion
// 获取最新版本客户端的的版本
// 未使用
func (c *Ctx) GetLastApkStrategyInfoByApkIdAndVersion(ctx context.Context, apkId int64) (*model.UpgradeApkUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkStrategyLastInfoByApkId, apkId)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var apkStrategyInfo model.UpgradeApkUpgradeStrategy
		query := fmt.Sprintf("select upgrade_apk_upgrade_strategy.* from upgrade_apk_upgrade_strategy " +
			"left join upgrade_apk_version on upgrade_apk_upgrade_strategy.apk_version_id = upgrade_apk_version.id " +
			"where upgrade_apk_upgrade_strategy.apk_id = ? " +
			"AND ? > upgrade_apk_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_apk_upgrade_strategy.end_datetime " +
			"AND upgrade_apk_upgrade_strategy.upgrade_dev_type = 0 " +
			"AND upgrade_apk_upgrade_strategy.enable = 1 " +
			"AND upgrade_apk_upgrade_strategy.is_del = 0 " +
			"order by upgrade_apk_version.version_code desc limit 1")

		//fmt.Println("query: ", query)
		err := c.mysqlConnCache.QueryRowCtx(ctx, &apkStrategyInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, apkId, now, now)
		})

		if err != nil {
			return nil, err
		}

		return apkStrategyInfo, nil
	})

	if v != nil {
		apkStrategyInfo := v.(model.UpgradeApkUpgradeStrategy)
		return &apkStrategyInfo, err
	}
	return nil, err
}
