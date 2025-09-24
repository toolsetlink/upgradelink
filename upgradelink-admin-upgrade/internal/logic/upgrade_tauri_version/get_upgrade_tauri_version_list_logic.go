package upgrade_tauri_version

import (
	"context"
	"upgradelink-admin-upgrade/internal/utils"
	"time"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradetauriversion"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
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

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	predicates = append(predicates, upgradetauriversion.CompanyIDEQ(companyID))

	// 删除状态
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeTauriVersionListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		tauriData, err := l.svcCtx.DB.UpgradeTauri.Get(l.ctx, v.TauriID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		cloudFileData, err := l.svcCtx.DB.FmsCloudFile.Get(l.ctx, v.CloudFileID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		if v.InstallCloudFileID == "" {
			v.InstallCloudFileID = v.CloudFileID
		}
		InstallCloudFileData, err := l.svcCtx.DB.FmsCloudFile.Get(l.ctx, v.InstallCloudFileID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		// 安装包文件大小
		installFileSize := utils.BytesToMBString(cloudFileData.Size)
		// 升级包文件大小
		upgradeFileSize := utils.BytesToMBString(InstallCloudFileData.Size)

		resp.Data.Data = append(resp.Data.Data,
			types.RespUpgradeTauriVersionInfo{
				Id:                     &v.ID,
				TauriId:                &v.TauriID,
				TauriName:              &tauriData.Name,
				CloudFileId:            &v.CloudFileID,
				CloudFileName:          &cloudFileData.Name,
				InstallCloudFileId:     &v.InstallCloudFileID,
				InstallCloudFileName:   &InstallCloudFileData.Name,
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
