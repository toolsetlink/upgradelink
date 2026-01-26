package upgrade_dev_group_relation

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevGroupRelationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevGroupRelationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevGroupRelationByIdLogic {
	return &GetUpgradeDevGroupRelationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevGroupRelationByIdLogic) GetUpgradeDevGroupRelationById(req *types.IDReq) (*types.UpgradeDevGroupRelationInfoResp, error) {
	data, err := l.svcCtx.DB.UpgradeDevGroupRelation.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeDevGroupRelationInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UpgradeDevGroupRelationInfo{
			Id:         &data.ID,
			DevId:      &data.DevID,
			DevGroupId: &data.DevGroupID,
		},
	}, nil
}
