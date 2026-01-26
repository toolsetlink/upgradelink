package upgrade_file_upgrade_strategy

import (
	"context"
	"strconv"
	"strings"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/entx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradefileupgradestrategy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeFileUpgradeStrategyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeFileUpgradeStrategyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeFileUpgradeStrategyLogic {
	return &CreateUpgradeFileUpgradeStrategyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeFileUpgradeStrategyLogic) CreateUpgradeFileUpgradeStrategy(req *types.UpgradeFileUpgradeStrategyInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgradeFileUpgradeStrategy(req)
	if err != nil {
		return nil, err
	}

	// 开启事务
	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		isDel := int32(0)

		// 创建相关灰度及流控策略
		updateGrayIds := make([]int, 0)
		if *req.IsGray == 1 {
			for i := 0; i < len(req.GrayDataInfo); i++ {
				grayBeginTime, _ := common.StringToTime(*req.GrayDataInfo[i].BeginDatetime)
				grayEndTime, _ := common.StringToTime(*req.GrayDataInfo[i].EndDatetime)

				grayInfo, err := l.svcCtx.DB.UpgradeFileUpgradeStrategyGrayStrategy.Create().
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

				flowLimitInfo, err := l.svcCtx.DB.UpgradeFileUpgradeStrategyFlowLimitStrategy.Create().
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

		beginTime, _ := common.StringToTime(*req.BeginDatetime)
		endTime, _ := common.StringToTime(*req.EndDatetime)

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

		updateGrayIdsStr := common.IntSliceToString(updateGrayIds)
		updateFlowLimitIdsStr := common.IntSliceToString(updateFlowLimitIds)

		_, err = l.svcCtx.DB.UpgradeFileUpgradeStrategy.Create().
			SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
			SetNotNilEnable(req.Enable).
			SetNotNilName(req.Name).
			SetNotNilDescription(req.Description).
			SetNotNilFileID(req.FileId).
			SetNotNilFileVersionID(req.FileVersionId).
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
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeFileUpgradeStrategyLogic) CheckCreateUpgradeFileUpgradeStrategy(req *types.UpgradeFileUpgradeStrategyInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeFileUpgradeStrategy
	predicates = append(predicates, upgradefileupgradestrategy.Name(*req.Name))
	predicates = append(predicates, upgradefileupgradestrategy.IsDelEQ(0))
	predicates = append(predicates, upgradefileupgradestrategy.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeFileUpgradeStrategy.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("任务名称重复")
	}

	// 判断任务的开始时间是否大于结束时间
	if !common.IsStartTimeBeforeEndTime(*req.BeginDatetime, *req.EndDatetime) {
		return http_error.NewCodeBadRequestError("任务开始时间大于或等于结束时间")
	}

	// 判断灰度策略的开始时间是否大于结束时间
	if *req.IsGray == 1 {
		for i := 0; i < len(req.GrayDataInfo); i++ {
			// 判断为启用状态的数据
			if *req.GrayDataInfo[i].Enable == 1 {
				if !common.IsStartTimeBeforeEndTime(*req.GrayDataInfo[i].BeginDatetime, *req.GrayDataInfo[i].EndDatetime) {
					return http_error.NewCodeBadRequestError("灰度策略存在开始时间小于或等于结束时间")
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
						if common.IsStartTimeBeforeEndTime(*req.GrayDataInfo[i].BeginDatetime, *req.GrayDataInfo[j].EndDatetime) &&
							common.IsStartTimeBeforeEndTime(*req.GrayDataInfo[j].BeginDatetime, *req.GrayDataInfo[i].EndDatetime) {
							return http_error.NewCodeBadRequestError("灰度策略存在开始时间与结束时间存在交集")
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
				if !common.IsStartTimeBeforeEndTime("2023-12-01 "+*req.FlowLimitDataInfo[i].Begintime, "2023-12-01 "+*req.FlowLimitDataInfo[i].Endtime) {
					return http_error.NewCodeBadRequestError("流控策略存在开始时间小于或等于结束时间")
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
						if common.IsStartTimeBeforeEndTime("2023-12-01 "+*req.FlowLimitDataInfo[i].Begintime, "2023-12-01 "+*req.FlowLimitDataInfo[j].Endtime) &&
							common.IsStartTimeBeforeEndTime("2023-12-01 "+*req.FlowLimitDataInfo[j].Begintime, "2023-12-01 "+*req.FlowLimitDataInfo[i].Endtime) {
							return http_error.NewCodeBadRequestError("流控策略存在开始时间与结束时间存在交集")
						}
					}
				}
			}
		}
	}

	return nil
}
