package upgrade_configuration_upgrade_strategy

import (
	"context"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeConfigurationUpgradeStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeConfigurationUpgradeStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeConfigurationUpgradeStrategyLogic {
	return &DeleteUpgradeConfigurationUpgradeStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeConfigurationUpgradeStrategyLogic) DeleteUpgradeConfigurationUpgradeStrategy(req *types.IDsReq) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		data, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Get(l.ctx, id)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
		if data.CompanyID != companyID {
			return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
		}
	}

	// 开启事务
	if err := base.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		for i := 0; i < len(req.Ids); i++ {

			// 获取策略信息
			strategyData, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Get(l.ctx, req.Ids[i])
			if err != nil {
				return err
			}

			intDel := int32(1)
			// 删除相关的灰度策略及流量限制策略
			grayIds := make([]int, 0)
			grayIds, _ = utils.SplitStringToIntSlice(strategyData.GrayData)
			for i := 0; i < len(grayIds); i++ {
				err = l.svcCtx.DB.UpgradeConfigurationUpgradeStrategyGrayStrategy.Update().
					Where(upgradeconfigurationupgradestrategygraystrategy.IDEQ(grayIds[i])).
					SetNotNilIsDel(&intDel).
					Exec(l.ctx)
				if err != nil {
					return err
				}
			}

			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = utils.SplitStringToIntSlice(strategyData.FlowLimitData)
			for i := 0; i < len(flowLimitIds); i++ {
				err = l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Update().
					Where(upgradeconfigurationupgradestrategy.IDEQ(flowLimitIds[i])).
					SetNotNilIsDel(&intDel).
					Exec(l.ctx)
				if err != nil {
					return err
				}
			}

			// 删除策略信息
			err = l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Update().
				Where(upgradeconfigurationupgradestrategy.IDIn(req.Ids...), upgradeconfigurationupgradestrategy.CompanyIDIn(companyID)).
				SetNotNilIsDel(&intDel).
				Exec(l.ctx)
			if err != nil {
				return err
			}

		}

		return nil

	}); err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.DeleteSuccess}, nil
}
