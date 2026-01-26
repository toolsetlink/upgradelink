package resource

import (
	"context"
	"errors"
	"fmt"
	"log"
	"upgradelink-admin/server/api/internal/resource/model"
	"upgradelink-admin/server/api/internal/svc"
)

// GetPatchInfo 获取指定 apkId、highApkVersionId、lowApkVersionId 的 patch 信息
func GetPatchInfo(ctx context.Context, svcCtx *svc.ServiceContext, apkId int64, patchAlgo int64, highApkVersionId, lowApkVersionId int64) (*model.UpgradeApkPatch, error) {

	query := fmt.Sprintf("select * from upgrade_apk_patch where status = 9 and is_del = 0  and `apk_id` = ? and `patch_algo` = ? and `high_apk_version_id` = ? and `low_apk_version_id` = ? limit 1")
	queryContext, err := svcCtx.DB.QueryContext(ctx, query, apkId, patchAlgo, highApkVersionId, lowApkVersionId)
	if err != nil {
		return nil, err
	}

	var result []model.UpgradeApkPatch
	// 解析结构体
	for queryContext.Next() {
		var info model.UpgradeApkPatch
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

// GetApkPatchListByStatus
// 获取指定patch 列表
func GetApkPatchListByStatus(ctx context.Context, svcCtx *svc.ServiceContext, status int64) ([]model.UpgradeApkPatch, error) {
	query := fmt.Sprintf("select * from upgrade_apk_patch where `status` = ? and is_del = 0 ")
	queryContext, err := svcCtx.DB.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}

	var result []model.UpgradeApkPatch
	// 解析结构体
	for queryContext.Next() {
		var info model.UpgradeApkPatch
		if err := queryContext.Scan(&info); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	return result, nil
}

type AddPatchInfoReq struct {
	CompanyId        int64  `db:"company_id"`          // 公司ID
	ApkId            int64  `db:"apk_id"`              // 安卓应用ID
	HighApkVersionId int64  `db:"high_apk_version_id"` // 外键：apk_version.id
	LowApkVersionId  int64  `db:"low_apk_version_id"`  // 外键：apk_version.id
	PatchAlgo        int64  `db:"patch_algo"`          // 差分算法 0:默认值无; 1 HDiffPatch;2 bsdiff;
	Status           int64  `db:"status"`              // 处理状态：0:尚未进行差分处理; 1:正在处理差分; 2:差分过程错误; 3:差分过程超时; 4:差分包有问题; 5:差分处理成功; 6:差分包大于新版本全量包; 7:上传文件中; 8:上传文件失败; 9:处理完成
	CloudFileId      string `db:"cloud_file_id"`       // 云文件id
}

func AddPatchInfo(ctx context.Context, svcCtx *svc.ServiceContext, req AddPatchInfoReq) error {
	// 插入前判断下是否已经有相同的记录
	query1 := fmt.Sprintf("select * from upgrade_apk_patch where is_del = 0 and  `apk_id` = ? and `patch_algo` = ? and `high_apk_version_id` = ? and `low_apk_version_id` = ? limit 1")
	_, err := svcCtx.DB.QueryContext(ctx, query1, req.ApkId, req.PatchAlgo, req.HighApkVersionId, req.LowApkVersionId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		// 如果为空则插入数据
		query2 := fmt.Sprintf("insert into %s (company_id, apk_id, high_apk_version_id, low_apk_version_id, patch_algo) values (?, ?, ?, ?, ?)", "upgrade_apk_patch")
		_, err = svcCtx.DB.ExecContext(ctx, query2, req.CompanyId, req.ApkId, req.HighApkVersionId, req.LowApkVersionId, req.PatchAlgo)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

// UpdateApkPatchStatus 更新 patch 状态
func UpdateApkPatchStatus(ctx context.Context, svcCtx *svc.ServiceContext, id int64, status int64, cloudFileId, description string) error {
	query := fmt.Sprintf("update %s set `status` = ?, cloud_file_id = ?, description = ? where `id` = ?", "upgrade_apk_patch")
	_, err := svcCtx.DB.ExecContext(ctx, query, status, cloudFileId, description, id)
	if err != nil {
		return err
	}

	return nil
}
