package upgrade_tauri_upgrade_strategy

import (
	"context"
	"time"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradetauriupgradestrategy"
	"upgradelink-admin-upgrade/server/ent/upgradetauriupgradestrategyflowlimitstrategy"
	"upgradelink-admin-upgrade/server/ent/upgradetauriupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeTauriUpgradeStrategyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeTauriUpgradeStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeTauriUpgradeStrategyListLogic {
	return &GetUpgradeTauriUpgradeStrategyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeTauriUpgradeStrategyListLogic) GetUpgradeTauriUpgradeStrategyList(req *types.UpgradeTauriUpgradeStrategyListReq) (*types.UpgradeTauriUpgradeStrategyListResp, error) {
	var predicates []predicate.UpgradeTauriUpgradeStrategy

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradetauriupgradestrategy.CompanyIDEQ(companyID))
	// 删除状态
	predicates = append(predicates, upgradetauriupgradestrategy.IsDelEQ(0))

	if req.Enable != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.EnableEQ(*req.Enable))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.NameContains(*req.Name))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.DescriptionContains(*req.Description))
	}
	if req.TauriId != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.TauriIDEQ(*req.TauriId))
	}
	if req.TauriVersionId != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.TauriVersionIDEQ(*req.TauriVersionId))
	}
	if req.UpgradeType != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.UpgradeTypeEQ(*req.UpgradeType))
	}
	if req.PromptUpgradeContent != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.PromptUpgradeContentContains(*req.PromptUpgradeContent))
	}
	if req.UpgradeDevType != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.UpgradeDevTypeEQ(*req.UpgradeDevType))
	}
	if req.UpgradeDevData != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.UpgradeDevDataContains(*req.UpgradeDevData))
	}
	//if req.BeginDatetime != nil {
	//	predicates = append(predicates, upgradetauriupgradestrategy.BeginDatetimeGTE(time.UnixMilli(*req.BeginDatetime)))
	//}
	//if req.EndDatetime != nil {
	//	predicates = append(predicates, upgradetauriupgradestrategy.EndDatetimeGTE(time.UnixMilli(*req.EndDatetime)))
	//}
	if req.IsDel != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradetauriupgradestrategy.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeTauriUpgradeStrategy.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeTauriUpgradeStrategyListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		tauriData, err := l.svcCtx.DB.UpgradeTauri.Get(l.ctx, v.TauriID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		versionData, err := l.svcCtx.DB.UpgradeTauriVersion.Get(l.ctx, v.TauriVersionID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		beginTime := v.BeginDatetime.Format("2006-01-02 15:04:05")
		endTime := v.EndDatetime.Format("2006-01-02 15:04:05")

		// 调用 splitStringToIntSlice 函数将字符串拆分成整数切片
		upgradeDevTypeOneData := make([]int, 0)
		upgradeDevTypeTwoData := make([]int, 0)
		if v.UpgradeDevType == 1 {
			upgradeDevTypeOneData, _ = splitStringToIntSlice(v.UpgradeDevData)
		} else if v.UpgradeDevType == 2 {
			upgradeDevTypeTwoData, _ = splitStringToIntSlice(v.UpgradeDevData)
		}

		upgradeVersionTypeOneData := make([]int, 0)
		if v.UpgradeVersionType == 1 {
			upgradeVersionTypeOneData, _ = splitStringToIntSlice(v.UpgradeVersionData)
		}
		// 处理灰度策略数据
		grayInfoList := make([]*types.TauriGrayDataInfo, 0)
		if v.IsGray == 1 {
			grayIds := make([]int, 0)
			grayIds, _ = splitStringToIntSlice(v.GrayData)
			var predicates1 []predicate.UpgradeTauriUpgradeStrategyGrayStrategy
			predicates1 = append(predicates1, upgradetauriupgradestrategygraystrategy.IDIn(grayIds...))
			grayData, err := l.svcCtx.DB.UpgradeTauriUpgradeStrategyGrayStrategy.Query().Where(predicates1...).All(l.ctx)
			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
			}

			for i := 0; i < len(grayData); i++ {
				grayBeginDatetime := grayData[i].BeginDatetime.Format("2006-01-02 15:04:05")
				grayEndDateTime := grayData[i].EndDatetime.Format("2006-01-02 15:04:05")
				grayInfo := &types.TauriGrayDataInfo{
					Enable:        &grayData[i].Enable,
					BeginDatetime: &grayBeginDatetime,
					EndDatetime:   &grayEndDateTime,
					Limit:         &grayData[i].Limit,
				}
				grayInfoList = append(grayInfoList, grayInfo)
			}
		}

		// 处理流量限制数据
		flowLimitInfoList := make([]*types.TauriFlowLimitDataInfo, 0)
		if v.IsFlowLimit == 1 {
			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = splitStringToIntSlice(v.FlowLimitData)
			var predicates2 []predicate.UpgradeTauriUpgradeStrategyFlowLimitStrategy
			predicates2 = append(predicates2, upgradetauriupgradestrategyflowlimitstrategy.IDIn(flowLimitIds...))
			flowLimitData, err := l.svcCtx.DB.UpgradeTauriUpgradeStrategyFlowLimitStrategy.Query().Where(predicates2...).All(l.ctx)
			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
			}

			for i := 0; i < len(flowLimitData); i++ {
				flowLimitBeginTime := flowLimitData[i].BeginTime
				flowLimitEndTime := flowLimitData[i].EndTime
				flowLimitInfo := &types.TauriFlowLimitDataInfo{
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
			types.RespUpgradeTauriUpgradeStrategyInfo{
				Id:                         &v.ID,
				Enable:                     &v.Enable,
				Name:                       &v.Name,
				Description:                &v.Description,
				TauriId:                    &v.TauriID,
				TauriName:                  &tauriData.Name,
				TauriVersionId:             &v.TauriVersionID,
				TauriVersionName:           &versionData.VersionName,
				TauriVersionCode:           &versionData.VersionCode,
				TauriTarget:                &versionData.Target,
				TauriArch:                  &versionData.Arch,
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
				IsDel:                      &v.IsDel,
				CreateAt:                   pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:                   pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
