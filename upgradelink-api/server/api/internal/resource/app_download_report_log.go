package resource

import (
	"context"
	"fmt"
	"time"
)

type AddAppDownloadReportLogReq struct {
	Id                  int64     `db:"id"`                     // ID
	CompanyId           int64     `db:"company_id"`             // 公司ID
	Timestamp           time.Time `db:"timestamp"`              // 事件发生时间
	AppKey              string    `db:"app_key"`                // 应用Key
	AppType             string    `db:"app_type"`               // 应用类型
	AppVersionId        int64     `db:"app_version_id"`         // 应用版本ID
	AppVersionCode      int64     `db:"app_version_code"`       // 应用版本号
	AppVersionPlatform  string    `db:"app_version_platform"`   // 应用版本操作系统:linux、darwin、windows
	AppVersionTarget    string    `db:"app_version_target"`     // 应用版本操作系统:linux、darwin、windows
	AppVersionArch      string    `db:"app_version_arch"`       // 应用版本机器架构:x86_64、i686、aarch64、armv7
	DownloadCloudFileId string    `db:"download_cloud_file_id"` // 下载云文件ID
	CreateAt            time.Time `db:"create_at"`              // 创建时间
}

func (c *Ctx) AddAppDownloadReportLog(ctx context.Context, req AddAppDownloadReportLogReq) (int64, error) {

	query := fmt.Sprintf("insert into %s (company_id, timestamp, app_key, app_type, app_version_id, app_version_code, app_version_platform, app_version_target, app_version_arch, download_cloud_file_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_app_download_report_log")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.Timestamp, req.AppKey, req.AppType, req.AppVersionId, req.AppVersionCode, req.AppVersionPlatform, req.AppVersionTarget, req.AppVersionArch, req.DownloadCloudFileId)
	fmt.Println(ret.RowsAffected())
	if err != nil {
		return 0, err
	}

	lastId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}
