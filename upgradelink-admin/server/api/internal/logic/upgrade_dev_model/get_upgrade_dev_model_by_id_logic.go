package upgrade_dev_model

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevModelByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevModelByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevModelByIdLogic {
	return &GetUpgradeDevModelByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevModelByIdLogic) GetUpgradeDevModelById(req *types.IDReq) (*types.UpgradeDevModelInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeDevModel.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeDevModelInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UpgradeDevModelInfo{
			Id:       &data.ID,
			Key:      &data.Key,
			Name:     &data.Name,
			IsDel:    &data.IsDel,
			CreateAt: pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt: pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
