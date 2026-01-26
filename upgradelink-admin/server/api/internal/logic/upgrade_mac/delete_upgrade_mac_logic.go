package upgrade_mac

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/ent/upgrademac"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeMacLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeMacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeMacLogic {
	return &DeleteUpgradeMacLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeMacLogic) DeleteUpgradeMac(req *types.IDsReq) (*types.BaseMsgResp, error) {

	intDel := int32(1)
	var Ids []int
	for _, id := range req.Ids {
		Ids = append(Ids, int(id))
	}
	err := l.svcCtx.DB.UpgradeMac.Update().
		Where(upgrademac.IDIn(Ids...), upgrademac.CompanyIDIn(companyctx.GetCompanyIDFromCtx(l.ctx))).
		SetNotNilIsDel(&intDel).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
