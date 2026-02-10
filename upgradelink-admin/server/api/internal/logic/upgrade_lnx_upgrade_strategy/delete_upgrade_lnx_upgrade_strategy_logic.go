package upgrade_lnx_upgrade_strategy

import (
	"context"
	"strconv"
	"strings"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/entx"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/upgradelnxupgradestrategy"
	"upgradelink-admin/server/api/internal/ent/upgradelnxupgradestrategygraystrategy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeLnxUpgradeStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeLnxUpgradeStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeLnxUpgradeStrategyLogic {
	return &DeleteUpgradeLnxUpgradeStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeLnxUpgradeStrategyLogic) DeleteUpgradeLnxUpgradeStrategy(req *types.IDsReq) (*types.BaseMsgResp, error) {
	// 开启事务
	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		for i := 0; i < len(req.Ids); i++ {

			// 获取策略信息
			strategyData, err := l.svcCtx.DB.UpgradeLnxUpgradeStrategy.Get(l.ctx, int(req.Ids[i]))
			if err != nil {
				return err
			}

			intDelTrue := enum.IsDelTrue
			// 删除相关的灰度策略及流量限制策略
			grayIds := make([]int, 0)
			grayIds, _ = splitStringToIntSlice(strategyData.GrayData)
			for i := 0; i < len(grayIds); i++ {
				err = l.svcCtx.DB.UpgradeLnxUpgradeStrategyGrayStrategy.Update().
					Where(upgradelnxupgradestrategygraystrategy.IDEQ(grayIds[i])).
					SetNotNilIsDel(&intDelTrue).
					Exec(l.ctx)
				if err != nil {
					return err
				}
			}

			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = splitStringToIntSlice(strategyData.FlowLimitData)
			for i := 0; i < len(flowLimitIds); i++ {
				err = l.svcCtx.DB.UpgradeLnxUpgradeStrategy.Update().
					Where(upgradelnxupgradestrategy.IDEQ(flowLimitIds[i])).
					SetNotNilIsDel(&intDelTrue).
					Exec(l.ctx)
				if err != nil {
					return err
				}
			}

			var Ids []int
			for _, id := range req.Ids {
				Ids = append(Ids, int(id))
			}
			// 删除策略信息
			err = l.svcCtx.DB.UpgradeLnxUpgradeStrategy.Update().
				Where(upgradelnxupgradestrategy.IDIn(Ids...)).
				SetNotNilIsDel(&intDelTrue).
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

func splitStringToIntSlice(s string) ([]int, error) {
	// 使用 strings.Split 函数将字符串按逗号拆分成字符串切片
	strs := strings.Split(s, ",")
	// 创建一个整数切片，用于存储转换后的整数
	ints := make([]int, len(strs))
	// 遍历字符串切片，将每个字符串元素转换为整数
	for i, str := range strs {
		// 使用 strconv.Atoi 函数将字符串转换为整数
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		// 将转换后的整数添加到整数切片中
		ints[i] = num
	}
	return ints, nil
}
