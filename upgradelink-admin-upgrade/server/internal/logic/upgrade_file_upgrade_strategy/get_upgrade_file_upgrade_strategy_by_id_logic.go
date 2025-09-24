package upgrade_file_upgrade_strategy

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradefileupgradestrategyflowlimitstrategy"
	"upgradelink-admin-upgrade/server/ent/upgradefileupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeFileUpgradeStrategyByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeFileUpgradeStrategyByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeFileUpgradeStrategyByIdLogic {
	return &GetUpgradeFileUpgradeStrategyByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeFileUpgradeStrategyByIdLogic) GetUpgradeFileUpgradeStrategyById(req *types.IDReq) (*types.UpgradeFileUpgradeStrategyInfoResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	data, err := l.svcCtx.DB.UpgradeFileUpgradeStrategy.Get(l.ctx, req.Id)
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
		upgradeDevTypeOneData, _ = splitStringToIntSlice(data.UpgradeDevData)
	} else if data.UpgradeDevType == 2 {
		upgradeDevTypeTwoData, _ = splitStringToIntSlice(data.UpgradeDevData)
	}

	upgradeVersionTypeOneData := make([]int, 0)
	if data.UpgradeVersionType == 1 {
		upgradeVersionTypeOneData, _ = splitStringToIntSlice(data.UpgradeVersionData)
	}
	// 处理灰度策略数据
	grayInfoList := make([]*types.FileGrayDataInfo, 0)
	if data.IsGray == 1 {
		grayIds := make([]int, 0)
		grayIds, _ = splitStringToIntSlice(data.GrayData)
		var predicates1 []predicate.UpgradeFileUpgradeStrategyGrayStrategy
		predicates1 = append(predicates1, upgradefileupgradestrategygraystrategy.IDIn(grayIds...))
		grayData, err := l.svcCtx.DB.UpgradeFileUpgradeStrategyGrayStrategy.Query().Where(predicates1...).All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		for i := 0; i < len(grayData); i++ {
			grayBeginDatetime := grayData[i].BeginDatetime.Format("2006-01-02 15:04:05")
			grayEndDateTime := grayData[i].EndDatetime.Format("2006-01-02 15:04:05")
			grayInfo := &types.FileGrayDataInfo{
				Enable:        &grayData[i].Enable,
				BeginDatetime: &grayBeginDatetime,
				EndDatetime:   &grayEndDateTime,
				Limit:         &grayData[i].Limit,
			}
			grayInfoList = append(grayInfoList, grayInfo)
		}
	}

	// 处理流量限制数据
	flowLimitInfoList := make([]*types.FileFlowLimitDataInfo, 0)
	if data.IsFlowLimit == 1 {
		flowLimitIds := make([]int, 0)
		flowLimitIds, _ = splitStringToIntSlice(data.FlowLimitData)
		var predicates2 []predicate.UpgradeFileUpgradeStrategyFlowLimitStrategy
		predicates2 = append(predicates2, upgradefileupgradestrategyflowlimitstrategy.IDIn(flowLimitIds...))
		flowLimitData, err := l.svcCtx.DB.UpgradeFileUpgradeStrategyFlowLimitStrategy.Query().Where(predicates2...).All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		for i := 0; i < len(flowLimitData); i++ {
			flowLimitBeginTime := flowLimitData[i].BeginTime
			flowLimitEndTime := flowLimitData[i].EndTime
			flowLimitInfo := &types.FileFlowLimitDataInfo{
				Enable:    &flowLimitData[i].Enable,
				Begintime: &flowLimitBeginTime,
				Endtime:   &flowLimitEndTime,
				Dimension: &flowLimitData[i].Dimension,
				Limit:     &flowLimitData[i].Limit,
			}
			flowLimitInfoList = append(flowLimitInfoList, flowLimitInfo)
		}
	}

	return &types.UpgradeFileUpgradeStrategyInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},
		Data: types.RespUpgradeFileUpgradeStrategyInfo{
			Id:                         &data.ID,
			Enable:                     &data.Enable,
			Name:                       &data.Name,
			Description:                &data.Description,
			FileId:                     &data.FileID,
			FileVersionId:              &data.FileVersionID,
			UpgradeType:                &data.UpgradeType,
			PromptUpgradeContent:       &data.PromptUpgradeContent,
			UpgradeDevType:             &data.UpgradeDevType,
			UpgradeDevData:             &data.UpgradeDevData,
			UpgradeDevTypeZeroData:     &data.UpgradeDevData,
			UpgradeDevTypeOneData:      upgradeDevTypeOneData,
			UpgradeDevTypeTwoData:      upgradeDevTypeTwoData,
			UpgradeVersionData:         &data.UpgradeVersionData,
			UpgradeVersionTypeZeroData: &data.UpgradeVersionData,
			UpgradeVersionTypeOneData:  upgradeVersionTypeOneData,
			BeginDatetime:              &beginTime,
			EndDatetime:                &endTime,
			IsGray:                     &data.IsGray,
			GrayDataInfo:               grayInfoList,
			IsFlowLimit:                &data.IsFlowLimit,
			FlowLimitDataInfo:          flowLimitInfoList,
			IsDel:                      &data.IsDel,
			CreateAt:                   pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:                   pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
