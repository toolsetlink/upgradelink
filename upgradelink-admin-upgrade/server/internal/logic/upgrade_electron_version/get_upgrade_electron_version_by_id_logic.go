package upgrade_electron_version

import (
	"context"

	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeElectronVersionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeElectronVersionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeElectronVersionByIdLogic {
	return &GetUpgradeElectronVersionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeElectronVersionByIdLogic) GetUpgradeElectronVersionById(req *types.IDReq) (*types.UpgradeElectronVersionInfoResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	data, err := l.svcCtx.DB.UpgradeElectronVersion.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	// 效验下请求数据是否属于当前操作者
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	electronData, err := l.svcCtx.DB.UpgradeElectron.Get(l.ctx, data.ElectronID)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeElectronVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},
		Data: types.RespUpgradeElectronVersionInfo{
			Id:                   &data.ID,
			ElectronName:         &electronData.Name,
			CloudFileId:          &data.CloudFileID,
			CloudFileName:        &electronData.Name,
			InstallCloudFileId:   &data.InstallCloudFileID,
			InstallCloudFileName: &electronData.Name,
			VersionName:          &data.VersionName,
			VersionCode:          &data.VersionCode,
			Platform:             &data.Platform,
			Arch:                 &data.Arch,
			Sha512:               &data.Sha512,
			InstallSha512:        &data.InstallSha512,
			Description:          &data.Description,
			IsDel:                &data.IsDel,
			CreateAt:             pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:             pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
