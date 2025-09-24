package upgrade_url_version

import (
	"context"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeurlupgradestrategy"
	"upgradelink-admin-upgrade/ent/upgradeurlversion"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeUrlVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeUrlVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeUrlVersionLogic {
	return &DeleteUpgradeUrlVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeUrlVersionLogic) DeleteUpgradeUrlVersion(req *types.IDsReq) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		data, err := l.svcCtx.DB.UpgradeUrlVersion.Get(l.ctx, id)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if data.CompanyID != companyID {
			return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
		}

		// 判断当前版本是否已经未删除的策略，如果有策略则不能删除
		var predicates []predicate.UpgradeUrlUpgradeStrategy
		predicates = append(predicates, upgradeurlupgradestrategy.CompanyIDEQ(companyID))
		// 删除状态
		predicates = append(predicates, upgradeurlupgradestrategy.URLVersionID(id))
		predicates = append(predicates, upgradeurlupgradestrategy.IsDelEQ(0))
		strategyList, err := l.svcCtx.DB.UpgradeUrlUpgradeStrategy.Query().Where(predicates...).All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if len(strategyList) > 0 {
			return nil, errorx.NewCodeInvalidArgumentError("应用版本还存在对应的策略，不能删除")
		}
	}

	intDel := int32(1) // 删除状态
	err = l.svcCtx.DB.UpgradeUrlVersion.Update().
		Where(upgradeurlversion.IDIn(req.Ids...), upgradeurlversion.CompanyIDIn(companyID)).
		SetNotNilIsDel(&intDel).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
