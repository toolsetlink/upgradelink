package upgrade_dev_group_relation

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeDevGroupRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeDevGroupRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeDevGroupRelationLogic {
	return &UpdateUpgradeDevGroupRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeDevGroupRelationLogic) UpdateUpgradeDevGroupRelation(req *types.UpgradeDevGroupRelationInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.UpgradeDevGroupRelation.UpdateOneID(*req.Id).
		SetNotNilDevID(req.DevId).
		SetNotNilDevGroupID(req.DevGroupId).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
