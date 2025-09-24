package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"
)

type AddUpgradeUrlReqLog struct {
	Id          int64  `db:"id"`            // ID
	CompanyId   int64  `db:"company_id"`    // 公司ID
	UrlId       int64  `db:"url_id"`        // url应用ID
	VersionCode int64  `db:"version_code"`  // 应用版本号
	DevModelKey string `db:"dev_model_key"` // 设备机型唯一标识
	DevKey      string `db:"dev_key"`       // 设备唯一标识
}

func (c *Ctx) AddUrlReqLog(ctx context.Context, addUrlReqLogReq AddUpgradeUrlReqLog) (*model.UpgradeUrlReqLog, error) {

	r, err := c.mysqlConn.ExecCtx(context.Background(), "insert into upgrade_url_req_log (company_id, url_id,version_code,dev_model_key, dev_key) values (?, ?,?,?,?)",
		addUrlReqLogReq.CompanyId, addUrlReqLogReq.UrlId, addUrlReqLogReq.VersionCode, addUrlReqLogReq.DevModelKey, addUrlReqLogReq.DevKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.RowsAffected())

	return nil, err
}
