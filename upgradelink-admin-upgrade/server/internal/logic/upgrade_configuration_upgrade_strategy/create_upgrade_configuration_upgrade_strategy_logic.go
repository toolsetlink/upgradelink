package upgrade_configuration_upgrade_strategy

import (
	"context"
	"strconv"
	"strings"

	"upgradelink-admin-upgrade/server/ent"
	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeconfigurationupgradestrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeConfigurationUpgradeStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeConfigurationUpgradeStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeConfigurationUpgradeStrategyLogic {
	return &CreateUpgradeConfigurationUpgradeStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeConfigurationUpgradeStrategyLogic) CreateUpgradeConfigurationUpgradeStrategy(req *types.UpgradeConfigurationUpgradeStrategyInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeConfigurationUpgradeStrategy(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 开启事务
	if err := base.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		isDel := int32(0)
		// 创建相关灰度及流控策略
		updateGrayIds := make([]int, 0)
		if *req.IsGray == 1 {
			for i := 0; i < len(req.GrayDataInfo); i++ {
				grayBeginTime, _ := utils.StringToTime(*req.GrayDataInfo[i].BeginDatetime)
				grayEndTime, _ := utils.StringToTime(*req.GrayDataInfo[i].EndDatetime)

				grayInfo, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategyGrayStrategy.Create().
					SetEnable(*req.GrayDataInfo[i].Enable).
					SetNotNilBeginDatetime(grayBeginTime).
					SetNotNilEndDatetime(grayEndTime).
					SetNotNilLimit(req.GrayDataInfo[i].Limit).
					SetNotNilIsDel(&isDel).
					SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
					SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
					Save(l.ctx)
				if err != nil {
					return err
				}

				updateGrayIds = append(updateGrayIds, grayInfo.ID)
			}
		}

		updateFlowLimitIds := make([]int, 0)
		if *req.IsFlowLimit == 1 {
			for i := 0; i < len(req.FlowLimitDataInfo); i++ {
				flowLimitInfo, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategyFlowLimitStrategy.Create().
					SetEnable(*req.FlowLimitDataInfo[i].Enable).
					SetBeginTime(*req.FlowLimitDataInfo[i].Begintime).
					SetEndTime(*req.FlowLimitDataInfo[i].Endtime).
					SetNotNilDimension(req.FlowLimitDataInfo[i].Dimension).
					SetNotNilLimit(req.FlowLimitDataInfo[i].Limit).
					SetNotNilIsDel(&isDel).
					SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
					SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
					Save(l.ctx)
				if err != nil {
					return err
				}
				updateFlowLimitIds = append(updateFlowLimitIds, flowLimitInfo.ID)
			}
		}

		beginTime, _ := utils.StringToTime(*req.BeginDatetime)
		endTime, _ := utils.StringToTime(*req.EndDatetime)

		// 整理请求参数
		if *req.UpgradeDevType == 0 {
			kongStr := ""
			req.UpgradeDevData = &kongStr
		} else if *req.UpgradeDevType == 1 {
			strs := make([]string, len(req.UpgradeDevTypeOneData))
			for i, num := range req.UpgradeDevTypeOneData {
				strs[i] = strconv.Itoa(num)
			}
			upgradeDevData := strings.Join(strs, ",")
			req.UpgradeDevData = &upgradeDevData
		} else if *req.UpgradeDevType == 2 {
			strs := make([]string, len(req.UpgradeDevTypeTwoData))
			for i, num := range req.UpgradeDevTypeTwoData {
				strs[i] = strconv.Itoa(num)
			}
			upgradeDevData := strings.Join(strs, ",")
			req.UpgradeDevData = &upgradeDevData
		}

		if *req.UpgradeVersionType == 0 {
			kongStr := ""
			req.UpgradeVersionData = &kongStr
		} else if *req.UpgradeVersionType == 1 {
			strs := make([]string, len(req.UpgradeVersionTypeOneData))
			for i, num := range req.UpgradeVersionTypeOneData {
				strs[i] = strconv.Itoa(num)
			}
			upgradeVersionData := strings.Join(strs, ",")
			req.UpgradeVersionData = &upgradeVersionData
		}

		updateGrayIdsStr := utils.IntSliceToString(updateGrayIds)
		updateFlowLimitIdsStr := utils.IntSliceToString(updateFlowLimitIds)

		_, err = l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Create().
			SetNotNilCompanyID(&companyID).
			SetNotNilEnable(req.Enable).
			SetNotNilName(req.Name).
			SetNotNilDescription(req.Description).
			SetNotNilConfigurationID(req.ConfigurationId).
			SetNotNilConfigurationVersionID(req.ConfigurationVersionId).
			SetNotNilUpgradeType(req.UpgradeType).
			SetNotNilPromptUpgradeContent(req.PromptUpgradeContent).
			SetNotNilUpgradeDevType(req.UpgradeDevType).
			SetNotNilUpgradeDevData(req.UpgradeDevData).
			SetNotNilUpgradeVersionType(req.UpgradeVersionType).
			SetNotNilUpgradeVersionData(req.UpgradeVersionData).
			SetNotNilBeginDatetime(beginTime).
			SetNotNilEndDatetime(endTime).
			SetNotNilIsGray(req.IsGray).
			SetNotNilGrayData(&updateGrayIdsStr).
			SetNotNilIsFlowLimit(req.IsFlowLimit).
			SetNotNilFlowLimitData(&updateFlowLimitIdsStr).
			SetNotNilIsDel(&isDel).
			SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
			SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
			Save(l.ctx)
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeConfigurationUpgradeStrategyLogic) CheckCreateUpgradeConfigurationUpgradeStrategy(req *types.UpgradeConfigurationUpgradeStrategyInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeConfigurationUpgradeStrategy
	predicates = append(predicates, upgradeconfigurationupgradestrategy.Name(*req.Name))
	predicates = append(predicates, upgradeconfigurationupgradestrategy.IsDelEQ(0))
	predicates = append(predicates, upgradeconfigurationupgradestrategy.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeConfigurationUpgradeStrategy.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("任务名称重复")
	}

	// 判断任务的开始时间是否大于结束时间
	if !utils.IsStartTimeBeforeEndTime(*req.BeginDatetime, *req.EndDatetime) {
		return errorx.NewCodeInvalidArgumentError("任务开始时间大于或等于结束时间")
	}

	// 判断灰度策略的开始时间是否大于结束时间
	if *req.IsGray == 1 {
		for i := 0; i < len(req.GrayDataInfo); i++ {
			// 判断为启用状态的数据
			if *req.GrayDataInfo[i].Enable == 1 {
				if !utils.IsStartTimeBeforeEndTime(*req.GrayDataInfo[i].BeginDatetime, *req.GrayDataInfo[i].EndDatetime) {
					return errorx.NewCodeInvalidArgumentError("灰度策略存在开始时间小于或等于结束时间")
				}
			}

		}

		// 判断开始时间与结束时间是否存在交集
		for i := 0; i < len(req.GrayDataInfo); i++ {
			// 判断为启用状态的数据
			if *req.GrayDataInfo[i].Enable == 1 {
				for j := i + 1; j < len(req.GrayDataInfo); j++ {
					// 判断为启用状态的数据
					if *req.GrayDataInfo[j].Enable == 1 {
						if utils.IsStartTimeBeforeEndTime(*req.GrayDataInfo[i].BeginDatetime, *req.GrayDataInfo[j].EndDatetime) &&
							utils.IsStartTimeBeforeEndTime(*req.GrayDataInfo[j].BeginDatetime, *req.GrayDataInfo[i].EndDatetime) {
							return errorx.NewCodeInvalidArgumentError("灰度策略存在开始时间与结束时间存在交集")
						}
					}
				}
			}
		}
	}

	// 判断流控策略时间是否冲突
	if *req.IsFlowLimit == 1 {
		for i := 0; i < len(req.FlowLimitDataInfo); i++ {
			// 判断为启用状态的数据
			if *req.FlowLimitDataInfo[i].Enable == 1 {
				if !utils.IsStartTimeBeforeEndTime("2023-12-01 "+*req.FlowLimitDataInfo[i].Begintime, "2023-12-01 "+*req.FlowLimitDataInfo[i].Endtime) {
					return errorx.NewCodeInvalidArgumentError("流控策略存在开始时间小于或等于结束时间")
				}
			}
		}

		// 判断开始时间与结束时间是否存在交集
		for i := 0; i < len(req.FlowLimitDataInfo); i++ {
			// 判断为启用状态的数据
			if *req.FlowLimitDataInfo[i].Enable == 1 {
				for j := i + 1; j < len(req.FlowLimitDataInfo); j++ {
					// 判断为启用状态的数据
					if *req.FlowLimitDataInfo[j].Enable == 1 {
						if utils.IsStartTimeBeforeEndTime("2023-12-01 "+*req.FlowLimitDataInfo[i].Begintime, "2023-12-01 "+*req.FlowLimitDataInfo[j].Endtime) &&
							utils.IsStartTimeBeforeEndTime("2023-12-01 "+*req.FlowLimitDataInfo[j].Begintime, "2023-12-01 "+*req.FlowLimitDataInfo[i].Endtime) {
							return errorx.NewCodeInvalidArgumentError("流控策略存在开始时间与结束时间存在交集")
						}
					}
				}
			}
		}
	}

	return nil
}
