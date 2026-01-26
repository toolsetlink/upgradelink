package upgrade_mac_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgrademacupgradestrategy"
	"upgradelink-admin/server/api/internal/ent/upgrademacversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeMacVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeMacVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeMacVersionLogic {
	return &DeleteUpgradeMacVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeMacVersionLogic) DeleteUpgradeMacVersion(req *types.IDsReq) (*types.BaseMsgResp, error) {

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		// 判断当前版本是否已经未删除的策略，如果有策略则不能删除
		var predicates []predicate.UpgradeMacUpgradeStrategy
		predicates = append(predicates, upgrademacupgradestrategy.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
		// 删除状态
		predicates = append(predicates, upgrademacupgradestrategy.MACVersionID(int(id)))
		predicates = append(predicates, upgrademacupgradestrategy.IsDelEQ(0))
		strategyList, err := l.svcCtx.DB.UpgradeMacUpgradeStrategy.Query().Where(predicates...).All(l.ctx)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}
		if len(strategyList) > 0 {
			return nil, http_error.NewCodeBadRequestError("应用版本还存在对应的策略，不能删除")
		}
	}

	intDel := int32(1) // 删除状态
	var Ids []int
	for _, id := range req.Ids {
		Ids = append(Ids, int(id))
	}

	err := l.svcCtx.DB.UpgradeMacVersion.Update().
		Where(upgrademacversion.IDIn(Ids...)).
		SetNotNilIsDel(&intDel).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
