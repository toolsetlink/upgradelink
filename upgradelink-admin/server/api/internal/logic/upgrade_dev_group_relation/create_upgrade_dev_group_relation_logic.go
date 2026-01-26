package upgrade_dev_group_relation

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeDevGroupRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeDevGroupRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeDevGroupRelationLogic {
	return &CreateUpgradeDevGroupRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeDevGroupRelationLogic) CreateUpgradeDevGroupRelation(req *types.UpgradeDevGroupRelationInfo) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.UpgradeDevGroupRelation.Create().
		SetNotNilDevID(req.DevId).
		SetNotNilDevGroupID(req.DevGroupId).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
