package upgrade_configuration_upgrade_strategy

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategyflowlimitstrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeConfigurationUpgradeStrategyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeConfigurationUpgradeStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeConfigurationUpgradeStrategyListLogic {
	return &GetUpgradeConfigurationUpgradeStrategyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeConfigurationUpgradeStrategyListLogic) GetUpgradeConfigurationUpgradeStrategyList(req *types.UpgradeConfigurationUpgradeStrategyListReq) (*types.UpgradeConfigurationUpgradeStrategyListResp, error) {
	var predicates []predicate.UpgradeConfigurationUpgradeStrategy

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradeconfigurationupgradestrategy.CompanyIDEQ(companyID))

	// 删除状态
	predicates = append(predicates, upgradeconfigurationupgradestrategy.IsDelEQ(0))

	if req.ConfigurationId != nil {
		predicates = append(predicates, upgradeconfigurationupgradestrategy.ConfigurationIDEQ(*req.ConfigurationId))
	}
	if req.Enable != nil {
		predicates = append(predicates, upgradeconfigurationupgradestrategy.EnableEQ(*req.Enable))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradeconfigurationupgradestrategy.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeConfigurationUpgradeStrategyListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		configurationData, err := l.svcCtx.DB.UpgradeConfiguration.Get(l.ctx, v.ConfigurationID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		versionData, err := l.svcCtx.DB.UpgradeConfigurationVersion.Get(l.ctx, v.ConfigurationVersionID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		beginTime := v.BeginDatetime.Format("2006-01-02 15:04:05")
		endTime := v.EndDatetime.Format("2006-01-02 15:04:05")

		// 调用 splitStringToIntSlice 函数将字符串拆分成整数切片
		upgradeDevTypeOneData := make([]int, 0)
		upgradeDevTypeTwoData := make([]int, 0)
		if v.UpgradeDevType == 1 {
			upgradeDevTypeOneData, _ = utils.SplitStringToIntSlice(v.UpgradeDevData)
		} else if v.UpgradeDevType == 2 {
			upgradeDevTypeTwoData, _ = utils.SplitStringToIntSlice(v.UpgradeDevData)
		}

		upgradeVersionTypeOneData := make([]int, 0)
		if v.UpgradeVersionType == 1 {
			upgradeVersionTypeOneData, _ = utils.SplitStringToIntSlice(v.UpgradeVersionData)
		}

		// 处理灰度策略数据
		grayInfoList := make([]*types.ConfigurationGrayDataInfo, 0)
		if v.IsGray == 1 {
			grayIds := make([]int, 0)
			grayIds, _ = utils.SplitStringToIntSlice(v.GrayData)
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
		if v.IsFlowLimit == 1 {
			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = utils.SplitStringToIntSlice(v.FlowLimitData)
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

		resp.Data.Data = append(resp.Data.Data,
			types.RespUpgradeConfigurationUpgradeStrategyInfo{
				Id:                         &v.ID,
				Enable:                     &v.Enable,
				Name:                       &v.Name,
				Description:                &v.Description,
				ConfigurationId:            &v.ConfigurationID,
				ConfigurationName:          &configurationData.Name,
				ConfigurationVersionId:     &v.ConfigurationVersionID,
				ConfigurationVersionName:   &versionData.VersionName,
				ConfigurationVersionCode:   &versionData.VersionCode,
				UpgradeType:                &v.UpgradeType,
				PromptUpgradeContent:       &v.PromptUpgradeContent,
				UpgradeDevType:             &v.UpgradeDevType,
				UpgradeDevData:             &v.UpgradeDevData,
				UpgradeDevTypeZeroData:     &v.UpgradeDevData,
				UpgradeDevTypeOneData:      upgradeDevTypeOneData,
				UpgradeDevTypeTwoData:      upgradeDevTypeTwoData,
				UpgradeVersionType:         &v.UpgradeVersionType,
				UpgradeVersionData:         &v.UpgradeVersionData,
				UpgradeVersionTypeZeroData: &v.UpgradeVersionData,
				UpgradeVersionTypeOneData:  upgradeVersionTypeOneData,
				IsGray:                     &v.IsGray,
				GrayDataInfo:               grayInfoList,
				IsFlowLimit:                &v.IsFlowLimit,
				FlowLimitDataInfo:          flowLimitInfoList,
				BeginDatetime:              &beginTime,
				EndDatetime:                &endTime,
				CreateAt:                   pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:                   pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
