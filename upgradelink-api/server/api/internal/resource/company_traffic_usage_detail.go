package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"
)

type AddCompanyTrafficUsageDetailReq struct {
	Id                     int64     `db:"id"`                        // ID
	CompanyId              int64     `db:"company_id"`                // 公司ID
	AppDownloadReportId    int64     `db:"app_download_report_id"`    // 应用下载记录ID
	CompanyTrafficPacketId int64     `db:"company_traffic_packet_id"` // 流量包 id
	UsedSize               int64     `db:"used_size"`                 // 使用流量
	UsedTime               time.Time `db:"used_time"`                 // 使用时间
	CreateAt               time.Time `db:"create_at"`                 // 创建时间
	UpdateAt               time.Time `db:"update_at"`                 // 修改时间
}

func (c *Ctx) AddCompanyTrafficUsageDetail(ctx context.Context, req AddCompanyTrafficUsageDetailReq) (*model.UpgradeCompanyTrafficUsageDetail, error) {

	query := fmt.Sprintf("insert into %s (company_id, app_download_report_id, company_traffic_packet_id, used_size, used_time) values (?, ?, ?, ?, ?)", "upgrade_company_traffic_usage_detail")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.AppDownloadReportId, req.CompanyTrafficPacketId, req.UsedSize, req.UsedTime)
	fmt.Println(ret.RowsAffected())

	return nil, err
}
