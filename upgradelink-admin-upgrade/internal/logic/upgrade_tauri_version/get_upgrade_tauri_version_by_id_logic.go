package upgrade_tauri_version

import (
	"context"

	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
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
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	data, err := l.svcCtx.DB.UpgradeTauriVersion.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	// 效验下请求数据是否属于当前操作者
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	tauriData, err := l.svcCtx.DB.UpgradeTauri.Get(l.ctx, data.TauriID)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeTauriVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
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
