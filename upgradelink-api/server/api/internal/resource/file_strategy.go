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

type AddFileStrategyReq struct {
	CompanyId            int64     `db:"company_id"`             // 公司ID
	Enable               int64     `db:"enable"`                 // 是否生效；可通过此控制策略是否生效0：失效；1：生效
	Name                 string    `db:"name"`                   // 任务名称
	Description          string    `db:"description"`            // 任务描述信息
	FileId               int64     `db:"file_id"`                // 文件应用ID
	FileVersionId        int64     `db:"file_version_id"`        // file_version_id; 外键file_version.id
	BeginDatetime        time.Time `db:"begin_datetime"`         // 升级任务开始时间
	EndDatetime          time.Time `db:"end_datetime"`           // 升级任务结束时间
	UpgradeType          int64     `db:"upgrade_type"`           // 升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级
	PromptUpgradeContent string    `db:"prompt_upgrade_content"` // 提示升级描述内容
}

func (c *Ctx) AddFileStrategy(ctx context.Context, req AddFileStrategyReq) (int64, error) {
	query := fmt.Sprintf("insert into %s (company_id, enable, name, description, file_id, file_version_id, begin_datetime, end_datetime, upgrade_type, prompt_upgrade_content, upgrade_dev_type, upgrade_dev_data, upgrade_version_type, upgrade_version_data, is_gray, gray_data, is_flow_limit, flow_limit_data, is_del) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 0, '', 0, '', 0, '', 0, '', 0)", "upgrade_file_upgrade_strategy")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.Enable, req.Name, req.Description, req.FileId, req.FileVersionId, req.BeginDatetime, req.EndDatetime, req.UpgradeType, req.PromptUpgradeContent)
	if err != nil {
		return 0, err
	}
	id, _ := ret.LastInsertId()
	return id, err
}
