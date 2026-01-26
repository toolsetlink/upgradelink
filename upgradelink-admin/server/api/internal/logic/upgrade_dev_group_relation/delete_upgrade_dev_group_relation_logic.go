package upgrade_dev_group_relation

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/upgradedevgrouprelation"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

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
	var Ids []int
	for _, id := range req.Ids {
		Ids = append(Ids, int(id))
	}
	_, err := l.svcCtx.DB.UpgradeDevGroupRelation.Delete().Where(upgradedevgrouprelation.IDIn(Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
