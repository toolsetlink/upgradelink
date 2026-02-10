package upgrade_lnx

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/upgradelnx"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeLnxLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeLnxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeLnxLogic {
	return &DeleteUpgradeLnxLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeLnxLogic) DeleteUpgradeLnx(req *types.IDsReq) (*types.BaseMsgResp, error) {

	intDelTrue := enum.IsDelTrue
	var Ids []int
	for _, id := range req.Ids {
		Ids = append(Ids, int(id))
	}
	err := l.svcCtx.DB.UpgradeLnx.Update().
		Where(upgradelnx.IDIn(Ids...)).
		SetNotNilIsDel(&intDelTrue).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
