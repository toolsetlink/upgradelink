package upgrade_configuration_upgrade_strategy

import (
	"context"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/entx"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/upgradeconfigurationupgradestrategy"
	"upgradelink-admin/server/api/internal/ent/upgradeconfigurationupgradestrategygraystrategy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

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

	// 开启事务
	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		for i := 0; i < len(req.Ids); i++ {

			// 获取策略信息
			strategyData, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Get(l.ctx, int(req.Ids[i]))
			if err != nil {
				return err
			}

			intDel := int32(1)
			// 删除相关的灰度策略及流量限制策略
			grayIds := make([]int, 0)
			grayIds, _ = common.SplitStringToIntSlice(strategyData.GrayData)
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
			flowLimitIds, _ = common.SplitStringToIntSlice(strategyData.FlowLimitData)
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
			var Ids []int
			for _, id := range req.Ids {
				Ids = append(Ids, int(id))
			}
			err = l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Update().
				Where(upgradeconfigurationupgradestrategy.IDIn(Ids...)).
				SetNotNilIsDel(&intDel).
				Exec(l.ctx)
			if err != nil {
				return err
			}

		}

		return nil

	}); err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
