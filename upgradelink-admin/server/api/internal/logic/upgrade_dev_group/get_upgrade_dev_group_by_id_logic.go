package upgrade_dev_group

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevGroupByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevGroupByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevGroupByIdLogic {
	return &GetUpgradeDevGroupByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevGroupByIdLogic) GetUpgradeDevGroupById(req *types.IDReq) (*types.UpgradeDevGroupInfoResp, error) {

	data, err := l.svcCtx.DB.UpgradeDevGroup.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeDevGroupInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UpgradeDevGroupInfo{
			Id:       &data.ID,
			Name:     &data.Name,
			IsDel:    &data.IsDel,
			CreateAt: pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt: pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
