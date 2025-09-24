package upgrade_dev_group_relation

import (
	"context"

	"upgradelink-admin-upgrade/ent/upgradedevgrouprelation"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeDevGroupRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeDevGroupRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeDevGroupRelationLogic {
	return &DeleteUpgradeDevGroupRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeDevGroupRelationLogic) DeleteUpgradeDevGroupRelation(req *types.IDsReq) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.UpgradeDevGroupRelation.Delete().Where(upgradedevgrouprelation.IDIn(req.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
