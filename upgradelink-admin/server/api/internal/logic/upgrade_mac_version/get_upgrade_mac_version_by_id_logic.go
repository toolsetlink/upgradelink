package upgrade_mac_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeMacVersionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeMacVersionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeMacVersionByIdLogic {
	return &GetUpgradeMacVersionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeMacVersionByIdLogic) GetUpgradeMacVersionById(req *types.IDReq) (*types.UpgradeMacVersionInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeMacVersion.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	macData, err := l.svcCtx.DB.UpgradeMac.Get(l.ctx, data.MACID)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeMacVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.RespUpgradeMacVersionInfo{
			Id:            &data.ID,
			MacName:       &macData.Name,
			CloudFileId:   &data.CloudFileID,
			CloudFileName: &macData.Name,
			VersionName:   &data.VersionName,
			VersionCode:   &data.VersionCode,
			Arch:          &data.Arch,
			Description:   &data.Description,
			IsDel:         &data.IsDel,
			CreateAt:      pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:      pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
