package upgrade_electron_upgrade_strategy

import (
	"context"
	"strconv"
	"strings"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/ent/upgradeelectronupgradestrategy"
	"upgradelink-admin-upgrade/ent/upgradeelectronupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUpgradeElectronUpgradeStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUpgradeElectronUpgradeStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUpgradeElectronUpgradeStrategyLogic {
	return &DeleteUpgradeElectronUpgradeStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUpgradeElectronUpgradeStrategyLogic) DeleteUpgradeElectronUpgradeStrategy(req *types.IDsReq) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	// 效验下请求数据是否属于当前操作者
	for _, id := range req.Ids {
		data, err := l.svcCtx.DB.UpgradeElectronUpgradeStrategy.Get(l.ctx, id)
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
			strategyData, err := l.svcCtx.DB.UpgradeElectronUpgradeStrategy.Get(l.ctx, req.Ids[i])
			if err != nil {
				return err
			}

			intDel := int32(1)
			// 删除相关的灰度策略及流量限制策略
			grayIds := make([]int, 0)
			grayIds, _ = splitStringToIntSlice(strategyData.GrayData)
			for i := 0; i < len(grayIds); i++ {
				err = l.svcCtx.DB.UpgradeElectronUpgradeStrategyGrayStrategy.Update().
					Where(upgradeelectronupgradestrategygraystrategy.IDEQ(grayIds[i])).
					SetNotNilIsDel(&intDel).
					Exec(l.ctx)
				if err != nil {
					return err
				}
			}

			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = splitStringToIntSlice(strategyData.FlowLimitData)
			for i := 0; i < len(flowLimitIds); i++ {
				err = l.svcCtx.DB.UpgradeElectronUpgradeStrategy.Update().
					Where(upgradeelectronupgradestrategy.IDEQ(flowLimitIds[i])).
					SetNotNilIsDel(&intDel).
					Exec(l.ctx)
				if err != nil {
					return err
				}
			}

			// 删除策略信息
			err = l.svcCtx.DB.UpgradeElectronUpgradeStrategy.Update().
				Where(upgradeelectronupgradestrategy.IDIn(req.Ids...), upgradeelectronupgradestrategy.CompanyIDIn(companyID)).
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
