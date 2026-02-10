package upgrade_electron

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/upgradeelectron"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeElectronLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeElectronLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeElectronLogic {
	return &DeleteUpgradeElectronLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeElectronLogic) DeleteUpgradeElectron(req *types.IDsReq) (*types.BaseMsgResp, error) {

	intDelTrue := enum.IsDelTrue
	var Ids []int
	for _, id := range req.Ids {
		Ids = append(Ids, int(id))
	}
	err := l.svcCtx.DB.UpgradeElectron.Update().
		Where(upgradeelectron.IDIn(Ids...)).
		SetNotNilIsDel(&intDelTrue).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
