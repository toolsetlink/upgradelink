package upgrade_configuration_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeConfigurationVersionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeConfigurationVersionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeConfigurationVersionByIdLogic {
	return &GetUpgradeConfigurationVersionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeConfigurationVersionByIdLogic) GetUpgradeConfigurationVersionById(req *types.IDReq) (*types.UpgradeConfigurationVersionInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeConfigurationVersion.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeConfigurationVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UpgradeConfigurationVersionInfo{
			Id:              &data.ID,
			ConfigurationId: &data.ConfigurationID,
			Content:         &data.Content,
			VersionName:     &data.VersionName,
			VersionCode:     &data.VersionCode,
			Description:     &data.Description,
			IsDel:           &data.IsDel,
			CreateAt:        pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:        pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
