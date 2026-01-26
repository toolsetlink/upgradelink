package resource

import (
	"context"
	"fmt"
	"log"
	"upgradelink-admin/server/api/internal/resource/model"
	"upgradelink-admin/server/api/internal/svc"
)

func GetApkInfoByKey(ctx context.Context, svcCtx *svc.ServiceContext,
	key string) (*model.UpgradeApk, error) {

	query := fmt.Sprintf("select * from upgrade_apk where `key` = ? and is_del = 0 limit 1")
	queryContext, err := svcCtx.DB.QueryContext(ctx, query, key)
	if err != nil {
		return nil, err
	}

	var result []model.UpgradeApk
	// 解析结构体
	for queryContext.Next() {
		var info model.UpgradeApk
		if err := queryContext.Scan(&info); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}
