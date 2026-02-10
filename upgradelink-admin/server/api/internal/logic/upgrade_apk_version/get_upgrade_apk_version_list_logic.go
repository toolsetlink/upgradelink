package upgrade_apk_version

import (
	"context"
	"fmt"
	"time"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/common/utils/cdn"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeapkversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeApkVersionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeApkVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeApkVersionListLogic {
	return &GetUpgradeApkVersionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeApkVersionListLogic) GetUpgradeApkVersionList(req *types.UpgradeApkVersionListReq) (*types.UpgradeApkVersionListResp, error) {
	var predicates []predicate.UpgradeApkVersion
	predicates = append(predicates, upgradeapkversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradeapkversion.IsDelEQ(0))

	if req.ApkId != nil {
		predicates = append(predicates, upgradeapkversion.ApkIDEQ(*req.ApkId))
	}
	if req.CloudFileId != nil {
		predicates = append(predicates, upgradeapkversion.CloudFileIDContains(*req.CloudFileId))
	}
	if req.VersionName != nil {
		predicates = append(predicates, upgradeapkversion.VersionNameContains(*req.VersionName))
	}
	if req.VersionCode != nil {
		predicates = append(predicates, upgradeapkversion.VersionCodeEQ(*req.VersionCode))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradeapkversion.DescriptionContains(*req.Description))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradeapkversion.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradeapkversion.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradeapkversion.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeApkVersion.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeApkVersionListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		fileData, err := l.svcCtx.DB.UpgradeApk.Get(l.ctx, v.ApkID)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		cloudFileData, err := l.svcCtx.DB.FmsCloudFile.Get(l.ctx, v.CloudFileID)
		if err != nil {
			fmt.Println("err: ", err)
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		// 文件大小
		fileSize := common.BytesToMBString(cloudFileData.Size)

		cloudFilePath := ""
		// cdn 连接
		if l.svcCtx.Config.UploadConf.CdnUrl != "" {
			cloudFilePath, err = cdn.ReturnCdnUrl(l.svcCtx.Config.UploadConf.CdnUrl, cloudFileData.URL)
			if err != nil {
				return nil, db_error.DefaultEntError(l.Logger, err, req)
			}
		}

		resp.Data.Data = append(resp.Data.Data,
			types.RespUpgradeApkVersionInfo{
				Id:              &v.ID,
				ApkId:           &v.ApkID,
				ApkName:         &fileData.Name,
				CloudFileId:     &v.CloudFileID,
				CloudFileName:   &cloudFileData.Name,
				CloudFilePath:   &cloudFilePath,
				VersionName:     &v.VersionName,
				VersionCode:     &v.VersionCode,
				VersionFileSize: &fileSize,
				Description:     &v.Description,
				IsDel:           &v.IsDel,
				CreateAt:        pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:        pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
