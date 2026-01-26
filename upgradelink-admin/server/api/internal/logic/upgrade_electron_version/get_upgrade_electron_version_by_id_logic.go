package upgrade_electron_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

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

	data, err := l.svcCtx.DB.UpgradeElectronVersion.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	electronData, err := l.svcCtx.DB.UpgradeElectron.Get(l.ctx, data.ElectronID)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeElectronVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
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
