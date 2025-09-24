package upgrade_dev_group_relation

import (
	"context"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"


	"github.com/suyuan32/simple-admin-common/msg/errormsg"
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: errormsg.UpdateSuccess}, nil
}
