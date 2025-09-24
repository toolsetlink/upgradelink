package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"
)

type AddAppUpgradeGetStrategyReportLogReq struct {
	Id                  int64     `db:"id"`                    // ID
	CompanyId           int64     `db:"company_id"`            // 公司ID
	Timestamp           time.Time `db:"timestamp"`             // 事件发生时间
	AppKey              string    `db:"app_key"`               // 应用Key
	AppType             string    `db:"app_type"`              // 应用类型
	AppVersionId        int64     `db:"app_version_id"`        // 应用版本ID
	AppVersionCode      int64     `db:"app_version_code"`      // 应用版本号
	DevModelKey         string    `db:"dev_model_key"`         // 设备机型唯一标识
	DevKey              string    `db:"dev_key"`               // 设备唯一标识
	AppVersionPlatform  string    `db:"app_version_platform"`  // 操作系统
	AppVersionTarget    string    `db:"app_version_target"`    // 操作系统
	AppVersionArch      string    `db:"app_version_arch"`      // 机器架构
	StrategyVersionId   int64     `db:"strategy_version_id"`   // 升级策略应用版本号
	StrategyVersionCode int64     `db:"strategy_version_code"` // 升级策略应用版本Id
	StrategyId          int64     `db:"strategy_id"`           // 升级策略Id
	CreateAt            time.Time `db:"create_at"`             // 创建时间
}

func (c *Ctx) AddAppUpgradeGetStrategyReportLog(ctx context.Context, req AddAppUpgradeGetStrategyReportLogReq) (*model.UpgradeAppUpgradeGetStrategyReportLog, error) {

	query := fmt.Sprintf("insert into %s (company_id, timestamp, app_key, app_type, app_version_id, app_version_code, dev_model_key, dev_key, app_version_platform, app_version_target, app_version_arch, strategy_version_id, strategy_version_code, strategy_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_app_upgrade_get_strategy_report_log")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.Timestamp, req.AppKey, req.AppType, req.AppVersionId, req.AppVersionCode, req.DevModelKey, req.DevKey, req.AppVersionPlatform, req.AppVersionTarget, req.AppVersionArch, req.StrategyVersionId, req.StrategyVersionCode, req.StrategyId)

	fmt.Println(ret.RowsAffected())

	return nil, err
}
