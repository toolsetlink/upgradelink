package upgrade_dev_group

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/entx"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/upgradedevgroup"
	"upgradelink-admin/server/api/internal/ent/upgradedevgrouprelation"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeDevGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeDevGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeDevGroupLogic {
	return &DeleteUpgradeDevGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeDevGroupLogic) DeleteUpgradeDevGroup(req *types.IDsReq) (*types.BaseMsgResp, error) {

	// 开启事务
	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		var Ids []int
		for _, id := range req.Ids {
			Ids = append(Ids, int(id))
		}

		intDel := int32(1)
		err := l.svcCtx.DB.UpgradeDevGroup.Update().
			Where(upgradedevgroup.IDIn(Ids...)).
			SetNotNilIsDel(&intDel).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		// 删除关联关系
		err = l.svcCtx.DB.UpgradeDevGroupRelation.Update().
			Where(upgradedevgrouprelation.DevGroupIDIn(Ids...)).
			SetNotNilIsDel(&intDel).
			Exec(l.ctx)
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
