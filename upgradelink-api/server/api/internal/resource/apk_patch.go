package resource

import (
	"context"
	"errors"
	"fmt"
	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CachePatchInfoByApkIdAndHighApkVersionIdAndLowApkVersionId = PREFIX + "APK_PATCH_INFO:APK_ID:%v:HIGH_APK_VERSION_ID:%v:LOW_APK_VERSION_ID:%v"
	CachePatchInfoByApkIdAndCloudFileId                        = PREFIX + "APK_PATCH_INFO:APK_ID:%v:CLOUD_FILE_ID:%v"
)

// GetPatchInfo 获取指定 apkId、highApkVersionId、lowApkVersionId 的 patch 信息
func (c *Ctx) GetPatchInfo(ctx context.Context, apkId int64, patchAlgo int64, highApkVersionId, lowApkVersionId int64) (*model.UpgradeApkPatch, error) {

	cacheKey := fmt.Sprintf(CachePatchInfoByApkIdAndHighApkVersionIdAndLowApkVersionId, apkId, highApkVersionId, lowApkVersionId)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeApkPatch
		query := fmt.Sprintf("select * from upgrade_apk_patch where status = 9 and is_del = 0  and `apk_id` = ? and `patch_algo` = ? and `high_apk_version_id` = ? and `low_apk_version_id` = ? limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, apkId, patchAlgo, highApkVersionId, lowApkVersionId)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeApkPatch)
		return &info, err
	}

	return nil, err
}

// GetPatchInfoByCloudFileId  获取指定 apkId、cloudFileId 的 patch 信息
func (c *Ctx) GetPatchInfoByCloudFileId(ctx context.Context, apkId int64, cloudFileId string) (*model.UpgradeApkPatch, error) {

	cacheKey := fmt.Sprintf(CachePatchInfoByApkIdAndCloudFileId, apkId, cloudFileId)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeApkPatch
		query := fmt.Sprintf("select * from upgrade_apk_patch where status = 9 and is_del = 0  and `apk_id` = ?  and `cloud_file_id` = ? limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, apkId, cloudFileId)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeApkPatch)
		return &info, err
	}

	return nil, err
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

func (c *Ctx) AddPatchInfo(ctx context.Context, req AddPatchInfoReq) error {
	// 插入前判断下是否已经有相同的记录
	var info model.UpgradeApkPatch
	query1 := fmt.Sprintf("select * from upgrade_apk_patch where is_del = 0 and `apk_id` = ? and `patch_algo` = ? and `high_apk_version_id` = ? and `low_apk_version_id` = ? limit 1")
	err := c.mysqlConn.QueryRowCtx(ctx, &info, query1, req.ApkId, req.PatchAlgo, req.HighApkVersionId, req.LowApkVersionId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		// 如果为空则插入数据
		query2 := fmt.Sprintf("insert into %s (company_id, apk_id, high_apk_version_id, low_apk_version_id, patch_algo) values (?, ?, ?, ?, ?)", "upgrade_apk_patch")
		ret, err := c.mysqlConn.ExecCtx(ctx, query2, req.CompanyId, req.ApkId, req.HighApkVersionId, req.LowApkVersionId, req.PatchAlgo)
		fmt.Println(ret.RowsAffected())
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
