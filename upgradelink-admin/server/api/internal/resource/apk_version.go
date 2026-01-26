package resource

import (
	"context"
	"fmt"
	"log"
	"upgradelink-admin/server/api/internal/resource/model"
	"upgradelink-admin/server/api/internal/svc"
)

// GetApkVersionListByApkIdAndVersionCode
// 获取大于 versionCode 的版本列表
func GetApkVersionListByApkIdAndVersionCode(ctx context.Context, svcCtx *svc.ServiceContext, apkId int64, versionCode int64) ([]model.UpgradeApkVersion, error) {

	query := fmt.Sprintf("select * from upgrade_apk_version where `apk_id` = ? and is_del = 0 and `version_code` > ? order by version_code desc ")
	queryContext, err := svcCtx.DB.QueryContext(ctx, query, apkId, versionCode)
	if err != nil {
		return nil, err
	}

	var result []model.UpgradeApkVersion
	// 解析结构体
	for queryContext.Next() {
		var info model.UpgradeApkVersion
		if err := queryContext.Scan(&info); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil

}

// GetApkVersionLastInfoByApkId
// 获取最新版本的信息
func GetApkVersionLastInfoByApkId(ctx context.Context, svcCtx *svc.ServiceContext, apkId int64) (*model.UpgradeApkVersion, error) {

	query := fmt.Sprintf("select * from upgrade_apk_version where `apk_id` = ? and is_del = 0 order by version_code desc limit 1")
	queryContext, err := svcCtx.DB.QueryContext(ctx, query, apkId)
	if err != nil {
		return nil, err
	}

	var result []model.UpgradeApkVersion
	// 解析结构体
	for queryContext.Next() {
		var info model.UpgradeApkVersion
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

// 获取固定版本的信息
func GetApkVersionInfoByApkIdAndVersionCode(ctx context.Context, svcCtx *svc.ServiceContext, apkId int64, versionCode int64) (*model.UpgradeApkVersion, error) {

	query := fmt.Sprintf("select * from upgrade_apk_version where `apk_id` = ? and is_del = 0 and `version_code` = ? LIMIT 1")
	queryContext, err := svcCtx.DB.QueryContext(ctx, query, apkId, versionCode)
	if err != nil {
		return nil, err
	}

	var result []model.UpgradeApkVersion
	// 解析结构体
	for queryContext.Next() {
		var info model.UpgradeApkVersion
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

// GetApkVersionInfoById
// 通过 apk version id 获取信息
func GetApkVersionInfoById(ctx context.Context, svcCtx *svc.ServiceContext,
	id int64) (*model.UpgradeApkVersion, error) {

	query := fmt.Sprintf("select * from upgrade_apk_version where `id` = ?  limit 1")
	queryContext, err := svcCtx.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var result []model.UpgradeApkVersion
	// 解析结构体
	for queryContext.Next() {
		var info model.UpgradeApkVersion
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
