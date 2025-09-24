package upgrade_dev_group_relation

import (
	"context"

	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}
