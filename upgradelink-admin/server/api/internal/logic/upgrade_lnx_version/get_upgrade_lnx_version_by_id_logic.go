package upgrade_lnx_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeLnxVersionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeLnxVersionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeLnxVersionByIdLogic {
	return &GetUpgradeLnxVersionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeLnxVersionByIdLogic) GetUpgradeLnxVersionById(req *types.IDReq) (*types.UpgradeLnxVersionInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeLnxVersion.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	lnxData, err := l.svcCtx.DB.UpgradeLnx.Get(l.ctx, data.LnxID)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeLnxVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.RespUpgradeLnxVersionInfo{
			Id:            &data.ID,
			LnxName:       &lnxData.Name,
			CloudFileId:   &data.CloudFileID,
			CloudFileName: &lnxData.Name,
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
