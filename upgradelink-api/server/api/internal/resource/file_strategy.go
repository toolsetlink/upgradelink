package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyFileStrategyInfoByDevTypeAllAndFileIdAndVersion = PREFIX + "FILE_STRATEGY_INFO:DEV_TYPE:ALL:FILE_ID:%v:CLIENT_VERSION_CODE:%v:"
	CacheKeyFileStrategyListByFileIdAndVersion              = PREFIX + "FILE_STRATEGY_LIST:FILE_ID:%v:CLIENT_VERSION_CODE:%v:"

	// 最新版本信息
	CacheKeyFileStrategyLastInfoByFileId = PREFIX + "FILE_STRATEGY_LAST_INFO:FILE_ID:%v"
)

// GetFileStrategyInfoByFileIdAndVersionAndDevTypeAll
// 获取大于 客户端的 versionCode 的版本 获取开启了 全部设备的策略
func (c *Ctx) GetFileStrategyInfoByFileIdAndVersionAndDevTypeAll(ctx context.Context, fileId int64, clientVersionCode int64) (*model.UpgradeFileUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileStrategyInfoByDevTypeAllAndFileIdAndVersion, fileId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var fileStrategyInfo model.UpgradeFileUpgradeStrategy
		query := fmt.Sprintf("select upgrade_file_upgrade_strategy.* from upgrade_file_upgrade_strategy " +
			"left join upgrade_file_version on upgrade_file_upgrade_strategy.file_version_id = upgrade_file_version.id " +
			"where upgrade_file_upgrade_strategy.file_id = ? " +
			"AND upgrade_file_version.version_code > ? " +
			"AND ? > upgrade_file_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_file_upgrade_strategy.end_datetime " +
			"AND upgrade_file_upgrade_strategy.upgrade_dev_type = 0 " +
			"AND upgrade_file_upgrade_strategy.enable = 1 " +
			"AND upgrade_file_upgrade_strategy.is_del = 0 " +
			"order by upgrade_file_version.version_code desc limit 1")

		//fmt.Println("query: ", query)
		err := c.mysqlConnCache.QueryRowCtx(ctx, &fileStrategyInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, fileId, clientVersionCode, now, now)
		})

		if err != nil {
			return nil, err
		}

		return fileStrategyInfo, nil
	})

	if v != nil {
		fileStrategyInfo := v.(model.UpgradeFileUpgradeStrategy)
		return &fileStrategyInfo, err
	}
	return nil, err
}

// GetFileStrategyListByFileIdAndVersion
// 获取大于 客户端的 versionCode 的版本 的全部策略 list
func (c *Ctx) GetFileStrategyListByFileIdAndVersion(ctx context.Context, fileId int64, clientVersionCode int64) ([]*model.UpgradeFileUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileStrategyListByFileIdAndVersion, fileId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var fileStrategyList []*model.UpgradeFileUpgradeStrategy
		query := fmt.Sprintf("select upgrade_file_upgrade_strategy.* from upgrade_file_upgrade_strategy " +
			"left join upgrade_file_version on upgrade_file_upgrade_strategy.file_version_id = upgrade_file_version.id " +
			"where upgrade_file_upgrade_strategy.file_id = ? " +
			"AND upgrade_file_version.version_code > ? " +
			"AND ? > upgrade_file_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_file_upgrade_strategy.end_datetime " +
			"AND upgrade_file_upgrade_strategy.enable = 1 " +
			"AND upgrade_file_upgrade_strategy.is_del = 0 " +
			"order by upgrade_file_version.version_code desc")
		fmt.Println("query: ", query)
		err := c.mysqlConn.QueryRowsCtx(context.Background(), &fileStrategyList, query, fileId, clientVersionCode, now, now)
		if err != nil {
			return nil, err
		}

		fmt.Println("GetFileStrategyListByFileIdAndVersion fileStrategyList: ", len(fileStrategyList))
		return fileStrategyList, nil
	})

	if v != nil {
		fileStrategyList := v.([]*model.UpgradeFileUpgradeStrategy)
		return fileStrategyList, err
	}

	return nil, err
}

// GetLastFileStrategyInfoByFileIdAndVersion
// 获取最新版本客户端的的版本
// 未使用
func (c *Ctx) GetLastFileStrategyInfoByFileIdAndVersion(ctx context.Context, fileId int64) (*model.UpgradeFileUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileStrategyLastInfoByFileId, fileId)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var fileStrategyInfo model.UpgradeFileUpgradeStrategy
		query := fmt.Sprintf("select upgrade_file_upgrade_strategy.* from upgrade_file_upgrade_strategy " +
			"left join upgrade_file_version on upgrade_file_upgrade_strategy.file_version_id = upgrade_file_version.id " +
			"where upgrade_file_upgrade_strategy.file_id = ? " +
			"AND ? > upgrade_file_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_file_upgrade_strategy.end_datetime " +
			"AND upgrade_file_upgrade_strategy.upgrade_dev_type = 0 " +
			"AND upgrade_file_upgrade_strategy.enable = 1 " +
			"AND upgrade_file_upgrade_strategy.is_del = 0 " +
			"order by upgrade_file_version.version_code desc limit 1")

		//fmt.Println("query: ", query)
		err := c.mysqlConnCache.QueryRowCtx(ctx, &fileStrategyInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, fileId, now, now)
		})

		if err != nil {
			return nil, err
		}

		return fileStrategyInfo, nil
	})

	if v != nil {
		fileStrategyInfo := v.(model.UpgradeFileUpgradeStrategy)
		return &fileStrategyInfo, err
	}
	return nil, err
}
