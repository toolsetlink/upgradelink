package upgrade_configuration_upgrade_strategy

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategyflowlimitstrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeConfigurationUpgradeStrategyByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeConfigurationUpgradeStrategyByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeConfigurationUpgradeStrategyByIdLogic {
	return &GetUpgradeConfigurationUpgradeStrategyByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeConfigurationUpgradeStrategyByIdLogic) GetUpgradeConfigurationUpgradeStrategyById(req *types.IDReq) (*types.UpgradeConfigurationUpgradeStrategyInfoResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	data, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	// 效验下请求数据是否属于当前操作者
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	beginTime := data.BeginDatetime.Format("2006-01-02 15:04:05")
	endTime := data.EndDatetime.Format("2006-01-02 15:04:05")

	// 调用 splitStringToIntSlice 函数将字符串拆分成整数切片
	upgradeDevTypeOneData := make([]int, 0)
	upgradeDevTypeTwoData := make([]int, 0)
	if data.UpgradeDevType == 1 {
		upgradeDevTypeOneData, _ = utils.SplitStringToIntSlice(data.UpgradeDevData)
	} else if data.UpgradeDevType == 2 {
		upgradeDevTypeTwoData, _ = utils.SplitStringToIntSlice(data.UpgradeDevData)
	}

	upgradeVersionTypeOneData := make([]int, 0)
	if data.UpgradeVersionType == 1 {
		upgradeVersionTypeOneData, _ = utils.SplitStringToIntSlice(data.UpgradeVersionData)
	}

	// 处理灰度策略数据
	grayInfoList := make([]*types.ConfigurationGrayDataInfo, 0)
	if data.IsGray == 1 {
		grayIds := make([]int, 0)
		grayIds, _ = utils.SplitStringToIntSlice(data.GrayData)
		var predicates1 []predicate.UpgradeConfigurationUpgradeStrategyGrayStrategy
		predicates1 = append(predicates1, upgradeconfigurationupgradestrategygraystrategy.IDIn(grayIds...))
		grayData, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategyGrayStrategy.Query().Where(predicates1...).All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		for i := 0; i < len(grayData); i++ {
			grayBeginDatetime := grayData[i].BeginDatetime.Format("2006-01-02 15:04:05")
			grayEndDateTime := grayData[i].EndDatetime.Format("2006-01-02 15:04:05")
			grayInfo := &types.ConfigurationGrayDataInfo{
				Enable:        &grayData[i].Enable,
				BeginDatetime: &grayBeginDatetime,
				EndDatetime:   &grayEndDateTime,
				Limit:         &grayData[i].Limit,
			}
			grayInfoList = append(grayInfoList, grayInfo)
		}
	}

	// 处理流量限制数据
	flowLimitInfoList := make([]*types.ConfigurationFlowLimitDataInfo, 0)
	if data.IsFlowLimit == 1 {
		flowLimitIds := make([]int, 0)
		flowLimitIds, _ = utils.SplitStringToIntSlice(data.FlowLimitData)
		var predicates2 []predicate.UpgradeConfigurationUpgradeStrategyFlowLimitStrategy
		predicates2 = append(predicates2, upgradeconfigurationupgradestrategyflowlimitstrategy.IDIn(flowLimitIds...))
		flowLimitData, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategyFlowLimitStrategy.Query().Where(predicates2...).All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		for i := 0; i < len(flowLimitData); i++ {
			flowLimitBeginTime := flowLimitData[i].BeginTime
			flowLimitEndTime := flowLimitData[i].EndTime
			flowLimitInfo := &types.ConfigurationFlowLimitDataInfo{
				Enable:    &flowLimitData[i].Enable,
				Begintime: &flowLimitBeginTime,
				Endtime:   &flowLimitEndTime,
				Dimension: &flowLimitData[i].Dimension,
				Limit:     &flowLimitData[i].Limit,
			}
			flowLimitInfoList = append(flowLimitInfoList, flowLimitInfo)
		}
	}

	return &types.UpgradeConfigurationUpgradeStrategyInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},

		Data: types.RespUpgradeConfigurationUpgradeStrategyInfo{
			Id:                         &data.ID,
			Enable:                     &data.Enable,
			Name:                       &data.Name,
			Description:                &data.Description,
			ConfigurationId:            &data.ConfigurationID,
			ConfigurationVersionId:     &data.ConfigurationVersionID,
			UpgradeType:                &data.UpgradeType,
			PromptUpgradeContent:       &data.PromptUpgradeContent,
			UpgradeDevType:             &data.UpgradeDevType,
			UpgradeDevData:             &data.UpgradeDevData,
			UpgradeDevTypeZeroData:     &data.UpgradeDevData,
			UpgradeDevTypeOneData:      upgradeDevTypeOneData,
			UpgradeDevTypeTwoData:      upgradeDevTypeTwoData,
			UpgradeVersionType:         &data.UpgradeVersionType,
			UpgradeVersionData:         &data.UpgradeVersionData,
			UpgradeVersionTypeZeroData: &data.UpgradeVersionData,
			UpgradeVersionTypeOneData:  upgradeVersionTypeOneData,
			BeginDatetime:              &beginTime,
			EndDatetime:                &endTime,
			IsGray:                     &data.IsGray,
			GrayDataInfo:               grayInfoList,
			IsFlowLimit:                &data.IsFlowLimit,
			FlowLimitDataInfo:          flowLimitInfoList,
			CreateAt:                   pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:                   pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
