package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyElectronStrategyInfoByDevTypeAllAndElectronIdAndVersion    = PREFIX + "ELECTRON_STRATEGY_INFO:DEV_TYPE:ALL:ELECTRON_ID:%v:CLIENT_VERSION_CODE:%v:"
	CacheKeyElectronStrategyListByElectronIdAndVersionAndTargetAndArch = PREFIX + "ELECTRON_STRATEGY_LIST:ELECTRON_ID:%v:CLIENT_VERSION_CODE:%v:platform:%v:arch:%v:"

	// 最新版本信息
	CacheKeyElectronStrategyLastInfoByElectronId = PREFIX + "ELECTRON_STRATEGY_LAST_INFO:ELECTRON_ID:%v"
)

// GetElectronStrategyInfoByElectronIdAndVersionAndDevTypeAll
// 获取大于 客户端的 versionCode 的版本 获取开启了 全部设备的策略
func (c *Ctx) GetElectronStrategyInfoByElectronIdAndVersionAndDevTypeAll(ctx context.Context, electronId int64, clientVersionCode int64) (*model.UpgradeElectronUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronStrategyInfoByDevTypeAllAndElectronIdAndVersion, electronId, clientVersionCode)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeElectronUpgradeStrategy
		query := fmt.Sprintf("select upgrade_electron_upgrade_strategy.* from upgrade_electron_upgrade_strategy " +
			"left join upgrade_electron_version on upgrade_electron_upgrade_strategy.electron_version_id = upgrade_electron_version.id " +
			"where upgrade_electron_upgrade_strategy.electron_id = ? " +
			"AND upgrade_electron_version.version_code > ? " +
			"AND ? > upgrade_electron_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_electron_upgrade_strategy.end_datetime " +
			"AND upgrade_electron_upgrade_strategy.upgrade_dev_type = 0 " +
			"AND upgrade_electron_upgrade_strategy.enable = 1 " +
			"AND upgrade_electron_upgrade_strategy.is_del = 0 " +
			"order by upgrade_electron_version.version_code desc limit 1")

		//fmt.Println("query: ", query)
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, electronId, clientVersionCode, now, now)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeElectronUpgradeStrategy)
		return &info, err
	}
	return nil, err
}

// GetElectronStrategyListByElectronIdAndVersionAndPlatformAndArch
// 获取大于 客户端的 versionCode 的版本 的全部策略 list
func (c *Ctx) GetElectronStrategyListByElectronIdAndVersionAndPlatformAndArch(ctx context.Context, electronId int64, clientVersionCode int64, platform string, arch string) ([]*model.UpgradeElectronUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronStrategyListByElectronIdAndVersionAndTargetAndArch, electronId, clientVersionCode, platform, arch)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var list []*model.UpgradeElectronUpgradeStrategy
		query := fmt.Sprintf("select upgrade_electron_upgrade_strategy.* from upgrade_electron_upgrade_strategy " +
			"left join upgrade_electron_version on upgrade_electron_upgrade_strategy.electron_version_id = upgrade_electron_version.id " +
			"where upgrade_electron_upgrade_strategy.electron_id = ? " +
			"AND upgrade_electron_version.platform = ? " +
			"AND upgrade_electron_version.arch = ? " +
			"AND upgrade_electron_version.version_code > ? " +
			"AND ? > upgrade_electron_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_electron_upgrade_strategy.end_datetime " +
			"AND upgrade_electron_upgrade_strategy.enable = 1 " +
			"AND upgrade_electron_upgrade_strategy.is_del = 0 " +
			"order by upgrade_electron_version.version_code desc")
		err := c.mysqlConn.QueryRowsCtx(ctx, &list, query, electronId, platform, arch, clientVersionCode, now, now)
		if err != nil {
			return nil, err
		}

		return list, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeElectronUpgradeStrategy)
		return list, err
	}

	return nil, err
}

// GetLastElectronStrategyInfoByElectronIdAndVersion
// 获取最新版本客户端的的版本
// 未使用
func (c *Ctx) GetLastElectronStrategyInfoByElectronIdAndVersion(ctx context.Context, electronId int64) (*model.UpgradeElectronUpgradeStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronStrategyLastInfoByElectronId, electronId)

	now := time.Now()

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeElectronUpgradeStrategy
		query := fmt.Sprintf("select upgrade_electron_upgrade_strategy.* from upgrade_electron_upgrade_strategy " +
			"left join upgrade_electron_version on upgrade_electron_upgrade_strategy.electron_version_id = upgrade_electron_version.id " +
			"where upgrade_electron_upgrade_strategy.electron_id = ? " +
			"AND ? > upgrade_electron_upgrade_strategy.begin_datetime " +
			"AND ? < upgrade_electron_upgrade_strategy.end_datetime " +
			"AND upgrade_electron_upgrade_strategy.upgrade_dev_type = 0 " +
			"AND upgrade_electron_upgrade_strategy.enable = 1 " +
			"AND upgrade_electron_upgrade_strategy.is_del = 0 " +
			"order by upgrade_electron_version.version_code desc limit 1")

		//fmt.Println("query: ", query)
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, electronId, now, now)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeElectronUpgradeStrategy)
		return &info, err
	}
	return nil, err
}

type AddElectronStrategyReq struct {
	Id                   int64     `db:"id"`                     // ID
	CompanyId            int64     `db:"company_id"`             // 公司ID
	Enable               int64     `db:"enable"`                 // 是否生效；可通过此控制策略是否生效0：失效；1：生效
	Name                 string    `db:"name"`                   // 任务名称
	Description          string    `db:"description"`            // 任务描述信息
	ElectronId           int64     `db:"electron_id"`            // Electron应用ID
	ElectronVersionId    int64     `db:"electron_version_id"`    // electron_version_id; 外键electron_version.id
	BeginDatetime        time.Time `db:"begin_datetime"`         // 升级任务开始时间
	EndDatetime          time.Time `db:"end_datetime"`           // 升级任务结束时间
	UpgradeType          int64     `db:"upgrade_type"`           // 升级方式：0：未知方式；1：提示升级；2：静默升级；3: 强制升级
	PromptUpgradeContent string    `db:"prompt_upgrade_content"` // 提示升级描述内容
	//UpgradeDevType       int64     `db:"upgrade_dev_type"`       // 指定升级的设备范围：0：全部设备；1：指定设备分组；2：指定机型
	//UpgradeDevData       string    `db:"upgrade_dev_data"`       // 升级设备数据：0.当为全部设备时，此字段为空；；1.当指定设备分组时，此字段存储设备分组id；2.当指定设备机型时，此字段存储选中的设备机型id;
	//UpgradeVersionType   int64     `db:"upgrade_version_type"`   // 指定升级的应用版本：0：全部版本；1：指定版本
	//UpgradeVersionData   string    `db:"upgrade_version_data"`   // 升级设备数据：0.当为全部版本时，此字段为空；；1.当指定应用版本时，此字段存储应用版本id;
	//IsGray               int64     `db:"is_gray"`                // 是否开启灰度 0：不开启；1：开启
	//GrayData             string    `db:"gray_data"`              // 灰度策略id数据
	//IsFlowLimit          int64     `db:"is_flow_limit"`          // 是否开启频控 0：不开启；1：开启
	//FlowLimitData        string    `db:"flow_limit_data"`        // 频控策略id数据
	//IsDel                int64     `db:"is_del"`                 // 是否删除 0：正常；1：已删除
	CreateAt time.Time `db:"create_at"` // 创建时间
	UpdateAt time.Time `db:"update_at"` // 修改时间
}

func (c *Ctx) AddElectronStrategy(ctx context.Context, req AddElectronStrategyReq) (int64, error) {

	query := fmt.Sprintf("insert into %s (company_id, enable, name, description, electron_id, electron_version_id, begin_datetime, end_datetime, upgrade_type, prompt_upgrade_content) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_electron_upgrade_strategy")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.Enable, req.Name, req.Description, req.ElectronId, req.ElectronVersionId, req.BeginDatetime, req.EndDatetime, req.UpgradeType, req.PromptUpgradeContent)
	if err != nil {
		return 0, err
	}
	id, _ := ret.LastInsertId()
	return id, err
}

// DelElectronStrategyByVersionId
// 删除 electron 版本
func (c *Ctx) DelElectronStrategyByVersionId(ctx context.Context, versionId int64) (int64, error) {

	query := fmt.Sprintf("update %s set `is_del` = 1 where `electron_version_id` = ?", "upgrade_electron_upgrade_strategy")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, versionId)
	if err != nil {
		return 0, err
	}
	id, _ := ret.LastInsertId()
	return id, err
}
