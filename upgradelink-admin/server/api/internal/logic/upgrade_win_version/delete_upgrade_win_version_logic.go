package upgrade_win_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradewinupgradestrategy"
	"upgradelink-admin/server/api/internal/ent/upgradewinversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeWinVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeWinVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeWinVersionLogic {
	return &DeleteUpgradeWinVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeWinVersionLogic) DeleteUpgradeWinVersion(req *types.IDsReq) (*types.BaseMsgResp, error) {
	for _, id := range req.Ids {
		// 判断当前版本是否已经未删除的策略，如果有策略则不能删除
		var predicates []predicate.UpgradeWinUpgradeStrategy
		predicates = append(predicates, upgradewinupgradestrategy.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
		// 删除状态
		predicates = append(predicates, upgradewinupgradestrategy.WinVersionID(int(id)))
		predicates = append(predicates, upgradewinupgradestrategy.IsDelEQ(0))
		strategyList, err := l.svcCtx.DB.UpgradeWinUpgradeStrategy.Query().Where(predicates...).All(l.ctx)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}
		if len(strategyList) > 0 {
			return nil, http_error.NewCodeBadRequestError("应用版本还存在对应的策略，不能删除")
		}
	}

	intDelTrue := enum.IsDelTrue // 删除状态
	var Ids []int
	for _, id := range req.Ids {
		Ids = append(Ids, int(id))
	}
	err := l.svcCtx.DB.UpgradeWinVersion.Update().
		Where(upgradewinversion.IDIn(Ids...), upgradewinversion.CompanyIDIn(companyctx.GetCompanyIDFromCtx(l.ctx))).
		SetNotNilIsDel(&intDelTrue).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
