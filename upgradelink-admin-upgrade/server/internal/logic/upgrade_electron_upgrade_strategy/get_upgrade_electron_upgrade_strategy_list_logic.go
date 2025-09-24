package upgrade_electron_upgrade_strategy

import (
	"context"
	"time"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeelectronupgradestrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeelectronupgradestrategyflowlimitstrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeelectronupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeElectronUpgradeStrategyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeElectronUpgradeStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeElectronUpgradeStrategyListLogic {
	return &GetUpgradeElectronUpgradeStrategyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeElectronUpgradeStrategyListLogic) GetUpgradeElectronUpgradeStrategyList(req *types.UpgradeElectronUpgradeStrategyListReq) (*types.UpgradeElectronUpgradeStrategyListResp, error) {
	var predicates []predicate.UpgradeElectronUpgradeStrategy

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradeelectronupgradestrategy.CompanyIDEQ(companyID))
	// 删除状态
	predicates = append(predicates, upgradeelectronupgradestrategy.IsDelEQ(0))

	if req.Enable != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.EnableEQ(*req.Enable))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.NameContains(*req.Name))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.DescriptionContains(*req.Description))
	}
	if req.ElectronId != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.ElectronIDEQ(*req.ElectronId))
	}
	if req.ElectronVersionId != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.ElectronVersionIDEQ(*req.ElectronVersionId))
	}
	if req.UpgradeType != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.UpgradeTypeEQ(*req.UpgradeType))
	}
	if req.PromptUpgradeContent != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.PromptUpgradeContentContains(*req.PromptUpgradeContent))
	}
	if req.UpgradeDevType != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.UpgradeDevTypeEQ(*req.UpgradeDevType))
	}
	if req.UpgradeDevData != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.UpgradeDevDataContains(*req.UpgradeDevData))
	}
	//if req.BeginDatetime != nil {
	//	predicates = append(predicates, upgradeelectronupgradestrategy.BeginDatetimeGTE(time.UnixMilli(*req.BeginDatetime)))
	//}
	//if req.EndDatetime != nil {
	//	predicates = append(predicates, upgradeelectronupgradestrategy.EndDatetimeGTE(time.UnixMilli(*req.EndDatetime)))
	//}
	if req.IsDel != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradeelectronupgradestrategy.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeElectronUpgradeStrategy.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeElectronUpgradeStrategyListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		electronData, err := l.svcCtx.DB.UpgradeElectron.Get(l.ctx, v.ElectronID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		versionData, err := l.svcCtx.DB.UpgradeElectronVersion.Get(l.ctx, v.ElectronVersionID)
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
		grayInfoList := make([]*types.ElectronGrayDataInfo, 0)
		if v.IsGray == 1 {
			grayIds := make([]int, 0)
			grayIds, _ = splitStringToIntSlice(v.GrayData)
			var predicates1 []predicate.UpgradeElectronUpgradeStrategyGrayStrategy
			predicates1 = append(predicates1, upgradeelectronupgradestrategygraystrategy.IDIn(grayIds...))
			grayData, err := l.svcCtx.DB.UpgradeElectronUpgradeStrategyGrayStrategy.Query().Where(predicates1...).All(l.ctx)
			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
			}

			for i := 0; i < len(grayData); i++ {
				grayBeginDatetime := grayData[i].BeginDatetime.Format("2006-01-02 15:04:05")
				grayEndDateTime := grayData[i].EndDatetime.Format("2006-01-02 15:04:05")
				grayInfo := &types.ElectronGrayDataInfo{
					Enable:        &grayData[i].Enable,
					BeginDatetime: &grayBeginDatetime,
					EndDatetime:   &grayEndDateTime,
					Limit:         &grayData[i].Limit,
				}
				grayInfoList = append(grayInfoList, grayInfo)
			}
		}

		// 处理流量限制数据
		flowLimitInfoList := make([]*types.ElectronFlowLimitDataInfo, 0)
		if v.IsFlowLimit == 1 {
			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = splitStringToIntSlice(v.FlowLimitData)
			var predicates2 []predicate.UpgradeElectronUpgradeStrategyFlowLimitStrategy
			predicates2 = append(predicates2, upgradeelectronupgradestrategyflowlimitstrategy.IDIn(flowLimitIds...))
			flowLimitData, err := l.svcCtx.DB.UpgradeElectronUpgradeStrategyFlowLimitStrategy.Query().Where(predicates2...).All(l.ctx)
			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
			}

			for i := 0; i < len(flowLimitData); i++ {
				flowLimitBeginTime := flowLimitData[i].BeginTime
				flowLimitEndTime := flowLimitData[i].EndTime
				flowLimitInfo := &types.ElectronFlowLimitDataInfo{
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
			types.RespUpgradeElectronUpgradeStrategyInfo{
				Id:                         &v.ID,
				Enable:                     &v.Enable,
				Name:                       &v.Name,
				Description:                &v.Description,
				ElectronId:                 &v.ElectronID,
				ElectronName:               &electronData.Name,
				ElectronVersionId:          &v.ElectronVersionID,
				ElectronVersionName:        &versionData.VersionName,
				ElectronVersionCode:        &versionData.VersionCode,
				ElectronPlatform:           &versionData.Platform,
				ElectronArch:               &versionData.Arch,
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
