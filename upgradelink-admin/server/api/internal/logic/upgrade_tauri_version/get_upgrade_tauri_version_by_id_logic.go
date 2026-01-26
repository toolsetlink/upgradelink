package upgrade_tauri_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeTauriVersionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeTauriVersionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeTauriVersionByIdLogic {
	return &GetUpgradeTauriVersionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeTauriVersionByIdLogic) GetUpgradeTauriVersionById(req *types.IDReq) (*types.UpgradeTauriVersionInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeTauriVersion.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	tauriData, err := l.svcCtx.DB.UpgradeTauri.Get(l.ctx, data.TauriID)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeTauriVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.RespUpgradeTauriVersionInfo{
			Id:                   &data.ID,
			TauriName:            &tauriData.Name,
			CloudFileId:          &data.CloudFileID,
			CloudFileName:        &tauriData.Name,
			InstallCloudFileId:   &data.InstallCloudFileID,
			InstallCloudFileName: &tauriData.Name,
			VersionName:          &data.VersionName,
			VersionCode:          &data.VersionCode,
			Target:               &data.Target,
			Arch:                 &data.Arch,
			Signature:            &data.Signature,
			Description:          &data.Description,
			IsDel:                &data.IsDel,
			CreateAt:             pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:             pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
