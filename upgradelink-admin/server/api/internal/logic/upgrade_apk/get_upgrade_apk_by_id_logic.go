package upgrade_apk

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeApkByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeApkByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeApkByIdLogic {
	return &GetUpgradeApkByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeApkByIdLogic) GetUpgradeApkById(req *types.IDReq) (*types.UpgradeApkInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeApk.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeApkInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UpgradeApkInfo{
			Id:          &data.ID,
			Key:         &data.Key,
			Name:        &data.Name,
			PackageName: &data.PackageName,
			Description: &data.Description,
			IsDel:       &data.IsDel,
			CreateAt:    pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:    pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
