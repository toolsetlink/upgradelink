package upgrade_lnx_upgrade_strategy

import (
	"context"
	"time"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradelnxupgradestrategy"
	"upgradelink-admin/server/api/internal/ent/upgradelnxupgradestrategyflowlimitstrategy"
	"upgradelink-admin/server/api/internal/ent/upgradelnxupgradestrategygraystrategy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeLnxUpgradeStrategyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeLnxUpgradeStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeLnxUpgradeStrategyListLogic {
	return &GetUpgradeLnxUpgradeStrategyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeLnxUpgradeStrategyListLogic) GetUpgradeLnxUpgradeStrategyList(req *types.UpgradeLnxUpgradeStrategyListReq) (*types.UpgradeLnxUpgradeStrategyListResp, error) {
	var predicates []predicate.UpgradeLnxUpgradeStrategy
	predicates = append(predicates, upgradelnxupgradestrategy.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradelnxupgradestrategy.IsDelEQ(0))

	if req.Enable != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.EnableEQ(*req.Enable))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.NameContains(*req.Name))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.DescriptionContains(*req.Description))
	}
	if req.LnxId != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.LnxIDEQ(*req.LnxId))
	}
	if req.LnxVersionId != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.LnxVersionIDEQ(*req.LnxVersionId))
	}
	if req.UpgradeType != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.UpgradeTypeEQ(*req.UpgradeType))
	}
	if req.PromptUpgradeContent != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.PromptUpgradeContentContains(*req.PromptUpgradeContent))
	}
	if req.UpgradeDevType != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.UpgradeDevTypeEQ(*req.UpgradeDevType))
	}
	if req.UpgradeDevData != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.UpgradeDevDataContains(*req.UpgradeDevData))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradelnxupgradestrategy.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeLnxUpgradeStrategy.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeLnxUpgradeStrategyListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		lnxData, err := l.svcCtx.DB.UpgradeLnx.Get(l.ctx, v.LnxID)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		versionData, err := l.svcCtx.DB.UpgradeLnxVersion.Get(l.ctx, v.LnxVersionID)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
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
		grayInfoList := make([]*types.LnxGrayDataInfo, 0)
		if v.IsGray == 1 {
			grayIds := make([]int, 0)
			grayIds, _ = splitStringToIntSlice(v.GrayData)
			var predicates1 []predicate.UpgradeLnxUpgradeStrategyGrayStrategy
			predicates1 = append(predicates1, upgradelnxupgradestrategygraystrategy.IDIn(grayIds...))
			grayData, err := l.svcCtx.DB.UpgradeLnxUpgradeStrategyGrayStrategy.Query().Where(predicates1...).All(l.ctx)
			if err != nil {
				return nil, db_error.DefaultEntError(l.Logger, err, req)
			}

			for i := 0; i < len(grayData); i++ {
				grayBeginDatetime := grayData[i].BeginDatetime.Format("2006-01-02 15:04:05")
				grayEndDateTime := grayData[i].EndDatetime.Format("2006-01-02 15:04:05")
				grayInfo := &types.LnxGrayDataInfo{
					Enable:        &grayData[i].Enable,
					BeginDatetime: &grayBeginDatetime,
					EndDatetime:   &grayEndDateTime,
					Limit:         &grayData[i].Limit,
				}
				grayInfoList = append(grayInfoList, grayInfo)
			}
		}

		// 处理流量限制数据
		flowLimitInfoList := make([]*types.LnxFlowLimitDataInfo, 0)
		if v.IsFlowLimit == 1 {
			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = splitStringToIntSlice(v.FlowLimitData)
			var predicates2 []predicate.UpgradeLnxUpgradeStrategyFlowLimitStrategy
			predicates2 = append(predicates2, upgradelnxupgradestrategyflowlimitstrategy.IDIn(flowLimitIds...))
			flowLimitData, err := l.svcCtx.DB.UpgradeLnxUpgradeStrategyFlowLimitStrategy.Query().Where(predicates2...).All(l.ctx)
			if err != nil {
				return nil, db_error.DefaultEntError(l.Logger, err, req)
			}

			for i := 0; i < len(flowLimitData); i++ {
				flowLimitBeginTime := flowLimitData[i].BeginTime
				flowLimitEndTime := flowLimitData[i].EndTime
				flowLimitInfo := &types.LnxFlowLimitDataInfo{
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
			types.RespUpgradeLnxUpgradeStrategyInfo{
				Id:                         &v.ID,
				Enable:                     &v.Enable,
				Name:                       &v.Name,
				Description:                &v.Description,
				LnxId:                      &v.LnxID,
				LnxName:                    &lnxData.Name,
				LnxVersionId:               &v.LnxVersionID,
				LnxVersionName:             &versionData.VersionName,
				LnxVersionCode:             &versionData.VersionCode,
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
