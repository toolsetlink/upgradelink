package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"
)

type AddAppReportLogReq struct {
	Id                  int64     `db:"id"`                    // ID
	CompanyId           int64     `db:"company_id"`            // 公司ID
	EventType           string    `db:"eventType"`             // 事件类型
	AppKey              string    `db:"app_key"`               // 应用Key
	VersionCode         int64     `db:"version_code"`          // 当前应用版本号
	DevModelKey         string    `db:"dev_model_key"`         // 设备机型唯一标识
	DevKey              string    `db:"dev_key"`               // 设备唯一标识
	Timestamp           time.Time `db:"timestamp"`             // 事件发生时间
	LaunchTime          time.Time `db:"launch_time"`           // 应用启动事件-应用启动时间
	Code                int64     `db:"code"`                  // 事件-状态码
	DownloadVersionCode int64     `db:"download_version_code"` // 应用升级-下载事件-应用版本号
	UpgradeVersionCode  int64     `db:"upgrade_version_code"`  // 应用升级-升级事件-应用版本号
	CreateAt            time.Time `db:"create_at"`             // 创建时间
}

func (c *Ctx) AddAppReportLog(ctx context.Context, req AddAppReportLogReq) (*model.UpgradeAppReportLog, error) {

	query := fmt.Sprintf("insert into %s (company_id, event_type, app_key, version_code, dev_model_key, dev_key, timestamp, launch_time, code, download_version_code, upgrade_version_code) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_app_report_log")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.EventType, req.AppKey, req.VersionCode, req.DevModelKey, req.DevKey, req.Timestamp, req.LaunchTime, req.Code, req.DownloadVersionCode, req.UpgradeVersionCode)
	fmt.Println(ret.RowsAffected())

	return nil, err
}
