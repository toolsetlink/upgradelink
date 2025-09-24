package upgrade_configuration_version

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationversion"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeConfigurationVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeConfigurationVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeConfigurationVersionLogic {
	return &DeleteUpgradeConfigurationVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeConfigurationVersionLogic) DeleteUpgradeConfigurationVersion(req *types.IDsReq) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		data, err := l.svcCtx.DB.UpgradeConfigurationVersion.Get(l.ctx, id)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if data.CompanyID != companyID {
			return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
		}

		// 判断当前版本是否已经未删除的策略，如果有策略则不能删除
		var predicates []predicate.UpgradeConfigurationUpgradeStrategy
		predicates = append(predicates, upgradeconfigurationupgradestrategy.CompanyIDEQ(companyID))
		// 删除状态
		predicates = append(predicates, upgradeconfigurationupgradestrategy.ConfigurationVersionID(id))
		predicates = append(predicates, upgradeconfigurationupgradestrategy.IsDelEQ(0))
		strategyList, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Query().Where(predicates...).All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if len(strategyList) > 0 {
			return nil, errorx.NewCodeInvalidArgumentError("版本还存在对应的策略，不能删除")
		}
	}

	intDel := int32(1) // 删除状态
	err = l.svcCtx.DB.UpgradeConfigurationVersion.Update().
		Where(upgradeconfigurationversion.IDIn(req.Ids...), upgradeconfigurationversion.CompanyIDIn(companyID)).
		SetNotNilIsDel(&intDel).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
