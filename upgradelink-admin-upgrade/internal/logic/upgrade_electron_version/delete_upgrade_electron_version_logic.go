package upgrade_electron_version

import (
	"context"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeelectronupgradestrategy"
	"upgradelink-admin-upgrade/ent/upgradeelectronversion"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeElectronVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeElectronVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeElectronVersionLogic {
	return &DeleteUpgradeElectronVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeElectronVersionLogic) DeleteUpgradeElectronVersion(req *types.IDsReq) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		data, err := l.svcCtx.DB.UpgradeElectronVersion.Get(l.ctx, id)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if data.CompanyID != companyID {
			return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
		}

		// 判断当前版本是否已经未删除的策略，如果有策略则不能删除
		var predicates []predicate.UpgradeElectronUpgradeStrategy
		predicates = append(predicates, upgradeelectronupgradestrategy.CompanyIDEQ(companyID))
		// 删除状态
		predicates = append(predicates, upgradeelectronupgradestrategy.ElectronVersionID(id))
		predicates = append(predicates, upgradeelectronupgradestrategy.IsDelEQ(0))
		strategyList, err := l.svcCtx.DB.UpgradeElectronUpgradeStrategy.Query().Where(predicates...).All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if len(strategyList) > 0 {
			return nil, errorx.NewCodeInvalidArgumentError("应用版本还存在对应的策略，不能删除")
		}
	}

	intDel := int32(1) // 删除状态
	err = l.svcCtx.DB.UpgradeElectronVersion.Update().
		Where(upgradeelectronversion.IDIn(req.Ids...), upgradeelectronversion.CompanyIDIn(companyID)).
		SetNotNilIsDel(&intDel).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
