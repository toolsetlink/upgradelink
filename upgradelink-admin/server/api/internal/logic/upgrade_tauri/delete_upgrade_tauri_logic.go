package upgrade_tauri

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/upgradetauri"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeTauriLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeTauriLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeTauriLogic {
	return &DeleteUpgradeTauriLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeTauriLogic) DeleteUpgradeTauri(req *types.IDsReq) (*types.BaseMsgResp, error) {

	intDel := int32(1)

	var Ids []int
	for _, id := range req.Ids {
		Ids = append(Ids, int(id))
	}

	err := l.svcCtx.DB.UpgradeTauri.Update().
		Where(upgradetauri.IDIn(Ids...)).
		SetNotNilIsDel(&intDel).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
