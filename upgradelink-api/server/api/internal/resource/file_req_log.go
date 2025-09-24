package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"
)

type AddUpgradeFileReqLog struct {
	Id          int64  `db:"id"`            // ID
	CompanyId   int64  `db:"company_id"`    // 公司ID
	FileId      int64  `db:"file_id"`       // 应用ID
	VersionCode int64  `db:"version_code"`  // 应用版本号
	DevModelKey string `db:"dev_model_key"` // 设备机型唯一标识
	DevKey      string `db:"dev_key"`       // 设备唯一标识
}

func (c *Ctx) AddFileReqLog(ctx context.Context, addFileReqLogReq AddUpgradeFileReqLog) (*model.UpgradeUrlReqLog, error) {

	r, err := c.mysqlConn.ExecCtx(ctx, "insert into upgrade_file_req_log (company_id, file_id, version_code, dev_model_key, dev_key) values (?, ?,?,?,?)",
		addFileReqLogReq.CompanyId, addFileReqLogReq.FileId, addFileReqLogReq.VersionCode, addFileReqLogReq.DevModelKey, addFileReqLogReq.DevKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.RowsAffected())

	return nil, err
}
