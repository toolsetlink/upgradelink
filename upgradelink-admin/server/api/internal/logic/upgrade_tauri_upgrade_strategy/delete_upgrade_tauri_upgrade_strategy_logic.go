package upgrade_tauri_upgrade_strategy

import (
	"context"
	"strconv"
	"strings"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/entx"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/upgradetauriupgradestrategy"
	"upgradelink-admin/server/api/internal/ent/upgradetauriupgradestrategygraystrategy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeTauriUpgradeStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeTauriUpgradeStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeTauriUpgradeStrategyLogic {
	return &DeleteUpgradeTauriUpgradeStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeTauriUpgradeStrategyLogic) DeleteUpgradeTauriUpgradeStrategy(req *types.IDsReq) (*types.BaseMsgResp, error) {

	// 开启事务
	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		for i := 0; i < len(req.Ids); i++ {

			// 获取策略信息
			strategyData, err := l.svcCtx.DB.UpgradeTauriUpgradeStrategy.Get(l.ctx, int(req.Ids[i]))
			if err != nil {
				return err
			}

			intDelTrue := enum.IsDelTrue
			// 删除相关的灰度策略及流量限制策略
			grayIds := make([]int, 0)
			grayIds, _ = splitStringToIntSlice(strategyData.GrayData)
			for i := 0; i < len(grayIds); i++ {
				err = l.svcCtx.DB.UpgradeTauriUpgradeStrategyGrayStrategy.Update().
					Where(upgradetauriupgradestrategygraystrategy.IDEQ(grayIds[i])).
					SetNotNilIsDel(&intDelTrue).
					Exec(l.ctx)
				if err != nil {
					return err
				}
			}

			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = splitStringToIntSlice(strategyData.FlowLimitData)
			for i := 0; i < len(flowLimitIds); i++ {
				err = l.svcCtx.DB.UpgradeTauriUpgradeStrategy.Update().
					Where(upgradetauriupgradestrategy.IDEQ(flowLimitIds[i])).
					SetNotNilIsDel(&intDelTrue).
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
			err = l.svcCtx.DB.UpgradeTauriUpgradeStrategy.Update().
				Where(upgradetauriupgradestrategy.IDIn(Ids...)).
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
