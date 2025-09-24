package upgrade_url_upgrade_strategy

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeurlupgradestrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeurlupgradestrategyflowlimitstrategy"
	"upgradelink-admin-upgrade/server/ent/upgradeurlupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeUrlUpgradeStrategyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeUrlUpgradeStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeUrlUpgradeStrategyListLogic {
	return &GetUpgradeUrlUpgradeStrategyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeUrlUpgradeStrategyListLogic) GetUpgradeUrlUpgradeStrategyList(req *types.UpgradeUrlUpgradeStrategyListReq) (*types.UpgradeUrlUpgradeStrategyListResp, error) {
	var predicates []predicate.UpgradeUrlUpgradeStrategy

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradeurlupgradestrategy.CompanyIDEQ(companyID))

	// 删除状态
	predicates = append(predicates, upgradeurlupgradestrategy.IsDelEQ(0))

	if req.UrlId != nil {
		predicates = append(predicates, upgradeurlupgradestrategy.URLIDEQ(*req.UrlId))
	}
	if req.Enable != nil {
		predicates = append(predicates, upgradeurlupgradestrategy.EnableEQ(*req.Enable))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradeurlupgradestrategy.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.UpgradeUrlUpgradeStrategy.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeUrlUpgradeStrategyListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		urlData, err := l.svcCtx.DB.UpgradeUrl.Get(l.ctx, v.URLID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		versionData, err := l.svcCtx.DB.UpgradeUrlVersion.Get(l.ctx, v.URLVersionID)
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
		grayInfoList := make([]*types.GrayDataInfo, 0)
		if v.IsGray == 1 {
			grayIds := make([]int, 0)
			grayIds, _ = utils.SplitStringToIntSlice(v.GrayData)
			var predicates1 []predicate.UpgradeUrlUpgradeStrategyGrayStrategy
			predicates1 = append(predicates1, upgradeurlupgradestrategygraystrategy.IDIn(grayIds...))
			grayData, err := l.svcCtx.DB.UpgradeUrlUpgradeStrategyGrayStrategy.Query().Where(predicates1...).All(l.ctx)
			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
			}

			for i := 0; i < len(grayData); i++ {
				grayBeginDatetime := grayData[i].BeginDatetime.Format("2006-01-02 15:04:05")
				grayEndDateTime := grayData[i].EndDatetime.Format("2006-01-02 15:04:05")
				grayInfo := &types.GrayDataInfo{
					Enable:        &grayData[i].Enable,
					BeginDatetime: &grayBeginDatetime,
					EndDatetime:   &grayEndDateTime,
					Limit:         &grayData[i].Limit,
				}
				grayInfoList = append(grayInfoList, grayInfo)
			}
		}

		// 处理流量限制数据
		flowLimitInfoList := make([]*types.FlowLimitDataInfo, 0)
		if v.IsFlowLimit == 1 {
			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = utils.SplitStringToIntSlice(v.FlowLimitData)
			var predicates2 []predicate.UpgradeUrlUpgradeStrategyFlowLimitStrategy
			predicates2 = append(predicates2, upgradeurlupgradestrategyflowlimitstrategy.IDIn(flowLimitIds...))
			flowLimitData, err := l.svcCtx.DB.UpgradeUrlUpgradeStrategyFlowLimitStrategy.Query().Where(predicates2...).All(l.ctx)
			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
			}

			for i := 0; i < len(flowLimitData); i++ {
				flowLimitBeginTime := flowLimitData[i].BeginTime
				flowLimitEndTime := flowLimitData[i].EndTime
				flowLimitInfo := &types.FlowLimitDataInfo{
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
			types.RespUpgradeUrlUpgradeStrategyInfo{
				Id:                         &v.ID,
				Enable:                     &v.Enable,
				Name:                       &v.Name,
				Description:                &v.Description,
				UrlId:                      &v.URLID,
				UrlName:                    &urlData.Name,
				UrlVersionId:               &v.URLVersionID,
				UrlVersionName:             &versionData.VersionName,
				UrlVersionCode:             &versionData.VersionCode,
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
