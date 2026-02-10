package upgrade_tauri_version

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
	"upgradelink-admin/server/api/internal/ent/upgradetauriversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeTauriVersionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeTauriVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeTauriVersionListLogic {
	return &GetUpgradeTauriVersionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeTauriVersionListLogic) GetUpgradeTauriVersionList(req *types.UpgradeTauriVersionListReq) (*types.UpgradeTauriVersionListResp, error) {
	var predicates []predicate.UpgradeTauriVersion
	predicates = append(predicates, upgradetauriversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradetauriversion.IsDelEQ(0))

	if req.TauriId != nil {
		predicates = append(predicates, upgradetauriversion.TauriIDEQ(*req.TauriId))
	}
	if req.CloudFileId != nil {
		predicates = append(predicates, upgradetauriversion.CloudFileIDContains(*req.CloudFileId))
	}
	if req.VersionName != nil {
		predicates = append(predicates, upgradetauriversion.VersionNameContains(*req.VersionName))
	}
	if req.VersionCode != nil {
		predicates = append(predicates, upgradetauriversion.VersionCodeEQ(*req.VersionCode))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradetauriversion.DescriptionContains(*req.Description))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradetauriversion.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradetauriversion.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradetauriversion.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeTauriVersion.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		fmt.Println("err1: ", err)
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeTauriVersionListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		tauriData, err := l.svcCtx.DB.UpgradeTauri.Get(l.ctx, v.TauriID)
		if err != nil {
			fmt.Println("err2: ", err)
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		cloudFileData, err := l.svcCtx.DB.FmsCloudFile.Get(l.ctx, v.CloudFileID)
		if err != nil {
			fmt.Println("err3: ", err)
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		if v.InstallCloudFileID == "" {
			v.InstallCloudFileID = v.CloudFileID
		}
		InstallCloudFileData, err := l.svcCtx.DB.FmsCloudFile.Get(l.ctx, v.InstallCloudFileID)
		if err != nil {
			fmt.Println("err4: ", err)
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		// 安装包文件大小
		installFileSize := common.BytesToMBString(cloudFileData.Size)
		// 升级包文件大小
		upgradeFileSize := common.BytesToMBString(InstallCloudFileData.Size)

		installCloudFilePath := ""
		// cdn 连接
		if l.svcCtx.Config.UploadConf.CdnUrl != "" {
			installCloudFilePath, err = cdn.ReturnCdnUrl(l.svcCtx.Config.UploadConf.CdnUrl, InstallCloudFileData.URL)
			if err != nil {
				return nil, db_error.DefaultEntError(l.Logger, err, req)
			}
		}

		resp.Data.Data = append(resp.Data.Data,
			types.RespUpgradeTauriVersionInfo{
				Id:                     &v.ID,
				TauriId:                &v.TauriID,
				TauriName:              &tauriData.Name,
				CloudFileId:            &v.CloudFileID,
				CloudFileName:          &cloudFileData.Name,
				InstallCloudFileId:     &v.InstallCloudFileID,
				InstallCloudFileName:   &InstallCloudFileData.Name,
				InstallCloudFilePath:   &installCloudFilePath,
				VersionName:            &v.VersionName,
				VersionCode:            &v.VersionCode,
				InstallVersionFileSize: &installFileSize,
				UpgradeVersionFileSize: &upgradeFileSize,
				Target:                 &v.Target,
				Arch:                   &v.Arch,
				Description:            &v.Description,
				Signature:              &v.Signature,
				IsDel:                  &v.IsDel,
				CreateAt:               pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:               pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
