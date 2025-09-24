package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"
)

type AddAppUpgradeDownloadReportLogReq struct {
	Id                  int64     `db:"id"`                    // ID
	CompanyId           int64     `db:"company_id"`            // 公司ID
	Timestamp           time.Time `db:"timestamp"`             // 事件发生时间
	AppKey              string    `db:"app_key"`               // 应用Key
	AppType             string    `db:"app_type"`              // 应用类型
	AppVersionId        int64     `db:"app_version_id"`        // 应用版本ID
	AppVersionCode      int64     `db:"app_version_code"`      // 应用版本号
	DevModelKey         string    `db:"dev_model_key"`         // 设备机型唯一标识
	DevKey              string    `db:"dev_key"`               // 设备唯一标识
	AppVersionTarget    string    `db:"app_version_target"`    // 操作系统
	AppVersionArch      string    `db:"app_version_arch"`      // 机器架构
	DownloadVersionCode int64     `db:"download_version_code"` // 应用下载版本号
	Code                int64     `db:"code"`                  // 事件-状态码
	CreateAt            time.Time `db:"create_at"`             // 创建时间
}

func (c *Ctx) AddAppUpgradeDownloadReportLog(ctx context.Context, req AddAppUpgradeDownloadReportLogReq) (*model.UpgradeAppUpgradeDownloadReportLog, error) {

	query := fmt.Sprintf("insert into %s (company_id, timestamp, app_key, app_type, app_version_id, app_version_code, dev_model_key, dev_key, app_version_target, app_version_arch, download_version_code, code) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_app_upgrade_download_report_log")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.Timestamp, req.AppKey, req.AppType, req.AppVersionId, req.AppVersionCode, req.DevModelKey, req.DevKey, req.AppVersionTarget, req.AppVersionArch, req.DownloadVersionCode, req.Code)

	fmt.Println(ret.RowsAffected())

	return nil, err
}
