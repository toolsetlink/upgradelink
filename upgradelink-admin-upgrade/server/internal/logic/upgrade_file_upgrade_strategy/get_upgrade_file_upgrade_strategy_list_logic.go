package upgrade_file_upgrade_strategy

import (
	"context"
	"strconv"
	"strings"
	"time"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradefileupgradestrategy"
	"upgradelink-admin-upgrade/server/ent/upgradefileupgradestrategyflowlimitstrategy"
	"upgradelink-admin-upgrade/server/ent/upgradefileupgradestrategygraystrategy"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeFileUpgradeStrategyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeFileUpgradeStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeFileUpgradeStrategyListLogic {
	return &GetUpgradeFileUpgradeStrategyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeFileUpgradeStrategyListLogic) GetUpgradeFileUpgradeStrategyList(req *types.UpgradeFileUpgradeStrategyListReq) (*types.UpgradeFileUpgradeStrategyListResp, error) {
	var predicates []predicate.UpgradeFileUpgradeStrategy

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradefileupgradestrategy.CompanyIDEQ(companyID))
	// 删除状态
	predicates = append(predicates, upgradefileupgradestrategy.IsDelEQ(0))

	if req.Enable != nil {
		predicates = append(predicates, upgradefileupgradestrategy.EnableEQ(*req.Enable))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradefileupgradestrategy.NameContains(*req.Name))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradefileupgradestrategy.DescriptionContains(*req.Description))
	}
	if req.FileId != nil {
		predicates = append(predicates, upgradefileupgradestrategy.FileIDEQ(*req.FileId))
	}
	if req.FileVersionId != nil {
		predicates = append(predicates, upgradefileupgradestrategy.FileVersionIDEQ(*req.FileVersionId))
	}
	if req.UpgradeType != nil {
		predicates = append(predicates, upgradefileupgradestrategy.UpgradeTypeEQ(*req.UpgradeType))
	}
	if req.PromptUpgradeContent != nil {
		predicates = append(predicates, upgradefileupgradestrategy.PromptUpgradeContentContains(*req.PromptUpgradeContent))
	}
	if req.UpgradeDevType != nil {
		predicates = append(predicates, upgradefileupgradestrategy.UpgradeDevTypeEQ(*req.UpgradeDevType))
	}
	if req.UpgradeDevData != nil {
		predicates = append(predicates, upgradefileupgradestrategy.UpgradeDevDataContains(*req.UpgradeDevData))
	}
	//if req.BeginDatetime != nil {
	//	predicates = append(predicates, upgradefileupgradestrategy.BeginDatetimeGTE(time.UnixMilli(*req.BeginDatetime)))
	//}
	//if req.EndDatetime != nil {
	//	predicates = append(predicates, upgradefileupgradestrategy.EndDatetimeGTE(time.UnixMilli(*req.EndDatetime)))
	//}
	if req.IsDel != nil {
		predicates = append(predicates, upgradefileupgradestrategy.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradefileupgradestrategy.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradefileupgradestrategy.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeFileUpgradeStrategy.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeFileUpgradeStrategyListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		fileData, err := l.svcCtx.DB.UpgradeFile.Get(l.ctx, v.FileID)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		versionData, err := l.svcCtx.DB.UpgradeFileVersion.Get(l.ctx, v.FileVersionID)
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
		grayInfoList := make([]*types.FileGrayDataInfo, 0)
		if v.IsGray == 1 {
			grayIds := make([]int, 0)
			grayIds, _ = splitStringToIntSlice(v.GrayData)
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
		if v.IsFlowLimit == 1 {
			flowLimitIds := make([]int, 0)
			flowLimitIds, _ = splitStringToIntSlice(v.FlowLimitData)
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

		resp.Data.Data = append(resp.Data.Data,
			types.RespUpgradeFileUpgradeStrategyInfo{
				Id:                         &v.ID,
				Enable:                     &v.Enable,
				Name:                       &v.Name,
				Description:                &v.Description,
				FileId:                     &v.FileID,
				FileName:                   &fileData.Name,
				FileVersionId:              &v.FileVersionID,
				FileVersionName:            &versionData.VersionName,
				FileVersionCode:            &versionData.VersionCode,
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
