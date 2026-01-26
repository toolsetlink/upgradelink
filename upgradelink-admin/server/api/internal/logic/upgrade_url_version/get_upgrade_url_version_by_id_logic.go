package upgrade_url_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeUrlVersionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeUrlVersionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeUrlVersionByIdLogic {
	return &GetUpgradeUrlVersionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeUrlVersionByIdLogic) GetUpgradeUrlVersionById(req *types.IDReq) (*types.UpgradeUrlVersionInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeUrlVersion.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeUrlVersionInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UpgradeUrlVersionInfo{
			Id:          &data.ID,
			UrlId:       &data.URLID,
			UrlPath:     &data.URLPath,
			VersionName: &data.VersionName,
			VersionCode: &data.VersionCode,
			Description: &data.Description,
			IsDel:       &data.IsDel,
			CreateAt:    pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:    pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
