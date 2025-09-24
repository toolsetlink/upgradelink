package upgrade_electron_version

import (
	"context"
	"time"
	"upgradelink-admin-upgrade/server/internal/utils"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeelectronversion"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeElectronVersionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeElectronVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeElectronVersionListLogic {
	return &GetUpgradeElectronVersionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeElectronVersionListLogic) GetUpgradeElectronVersionList(req *types.UpgradeElectronVersionListReq) (*types.UpgradeElectronVersionListResp, error) {
	var predicates []predicate.UpgradeElectronVersion

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	predicates = append(predicates, upgradeelectronversion.CompanyIDEQ(companyID))

	// 删除状态
	predicates = append(predicates, upgradeelectronversion.IsDelEQ(0))

	if req.ElectronId != nil {
		predicates = append(predicates, upgradeelectronversion.ElectronIDEQ(*req.ElectronId))
	}

	if req.VersionName != nil {
		predicates = append(predicates, upgradeelectronversion.VersionNameContains(*req.VersionName))
	}
	if req.VersionCode != nil {
		predicates = append(predicates, upgradeelectronversion.VersionCodeEQ(*req.VersionCode))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradeelectronversion.DescriptionContains(*req.Description))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradeelectronversion.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradeelectronversion.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradeelectronversion.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeElectronVersion.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeElectronVersionListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		electronData, err := l.svcCtx.DB.UpgradeElectron.Get(l.ctx, v.ElectronID)
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
			types.RespUpgradeElectronVersionInfo{
				Id:                     &v.ID,
				ElectronId:             &v.ElectronID,
				ElectronName:           &electronData.Name,
				CloudFileId:            &v.CloudFileID,
				CloudFileName:          &cloudFileData.Name,
				InstallCloudFileId:     &v.InstallCloudFileID,
				InstallCloudFileName:   &InstallCloudFileData.Name,
				VersionName:            &v.VersionName,
				VersionCode:            &v.VersionCode,
				InstallVersionFileSize: &installFileSize,
				UpgradeVersionFileSize: &upgradeFileSize,
				Platform:               &v.Platform,
				Arch:                   &v.Arch,
				Description:            &v.Description,
				Sha512:                 &v.Sha512,
				InstallSha512:          &v.InstallSha512,
				IsDel:                  &v.IsDel,
				CreateAt:               pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:               pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
