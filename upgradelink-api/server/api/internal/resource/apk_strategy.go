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

type AddApkStrategyReq struct {
	CompanyId            int64     `db:"company_id"`             // 公司ID
	Enable               int64     `db:"enable"`                 // 是否生效；可通过此控制策略是否生效0：失效；1：生效
	Name                 string    `db:"name"`                   // 任务名称
	Description          string    `db:"description"`            // 任务描述信息
	ApkId                int64     `db:"apk_id"`                 // 安卓应用ID
	ApkVersionId         int64     `db:"apk_version_id"`         // apk_version_id; 外键apk_version.id
	BeginDatetime        time.Time `db:"begin_datetime"`         // 升级任务开始时间
	EndDatetime          time.Time `db:"end_datetime"`           // 升级任务结束时间
	UpgradeType          int64     `db:"upgrade_type"`           // 升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级
	PromptUpgradeContent string    `db:"prompt_upgrade_content"` // 提示升级描述内容
}

func (c *Ctx) AddApkStrategy(ctx context.Context, req AddApkStrategyReq) (int64, error) {
	query := fmt.Sprintf("insert into %s (company_id, enable, name, description, apk_id, apk_version_id, begin_datetime, end_datetime, upgrade_type, prompt_upgrade_content, upgrade_dev_type, upgrade_dev_data, upgrade_version_type, upgrade_version_data, is_gray, gray_data, is_flow_limit, flow_limit_data, is_del) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 0, '', 0, '', 0, '', 0, '', 0)", "upgrade_apk_upgrade_strategy")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.Enable, req.Name, req.Description, req.ApkId, req.ApkVersionId, req.BeginDatetime, req.EndDatetime, req.UpgradeType, req.PromptUpgradeContent)
	if err != nil {
		return 0, err
	}
	id, _ := ret.LastInsertId()
	return id, err
}
