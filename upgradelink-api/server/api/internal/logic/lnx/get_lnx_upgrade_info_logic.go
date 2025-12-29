package lnx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLnxUpgradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLnxUpgradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLnxUpgradeInfoLogic {
	return &GetLnxUpgradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLnxUpgradeInfoLogic) GetLnxUpgradeInfo(req *types.GetLnxUpgradeInfoReq) (resp *types.GetLnxUpgradeInfoResp, err error) {

	// 请求参数效验
	if req.LnxKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, common.ErrLnx1Msg, common.ErrLnx1Docs)
	}
	if req.VersionCode < 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, common.ErrLnx1Msg, common.ErrLnx1Docs)
	}

	var res types.GetLnxUpgradeInfoResp

	// 通过唯一标识 获取到对应的应用信息
	lnxInfo, err := l.svcCtx.ResourceCtx.GetLnxInfoByKey(l.ctx, req.LnxKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, common.ErrLnx2Msg, common.ErrLnx2Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	// 查询应用版本表，判断是否有大于当前版本的  没有的话则代表当前就是最高版本
	_, err = l.svcCtx.ResourceCtx.GetLnxVersionListByLnxIdAndArchAndVersionCode(l.ctx, lnxInfo.Id, req.Arch, req.VersionCode)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		res.Code = 0
		res.Msg = common.AlreadyLatestVersionMsg
		res.Docs = common.AlreadyLatestVersionDocs
		return &res, nil
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	// 推出可使用的最高版本的升级策略
	lnxStrategyInfo, err := l.ReturnUpgradeStrategyInfo(lnxInfo.Id, req.Arch, req.VersionCode, req.AppointVersionCode, req.DevModelKey, req.DevKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		res.Code = 0
		res.Msg = common.AlreadyLatestVersionMsg
		res.Docs = common.AlreadyLatestVersionDocs
		return &res, nil
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	// 到这个地方的时候 说明前置条件已经都通过了，在这个位置再去判断 策略的频控配置是否符合
	flowLimitOk, err := l.CheckUpgradeStrategyFlowLimit(lnxStrategyInfo)
	if err != nil {
		return nil, err
	}
	if !flowLimitOk {
		// 被频控拦住 返回 429
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrStrategyTooManyReq, common.Err42901Msg, common.Err42901Docs)
	}

	// 通过升级版本 id 查询出对应版本信息 获取文件下载地址
	lnxVersionInfo, err := l.svcCtx.ResourceCtx.GetLnxVersionInfoById(l.ctx, lnxStrategyInfo.LnxVersionId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	// 通过文件信息
	cloudFileInfo, err := l.svcCtx.ResourceCtx.GetCloudFileInfoById(l.ctx, lnxVersionInfo.CloudFileId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}

	urlPath := ""
	// 获取云文件地址
	urlPath = cloudFileInfo.Url

	// 插入获取日志上报
	timestamp, err := common.ParseRFC3339ToTime(time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs)
	}
	// 获取应用版本 id
	appVersionId, err := l.svcCtx.ResourceCtx.GetAppVersionIdByReport(l.ctx, resource.GetAppVersionIdByReportReq{
		AppKey:           lnxInfo.Key,
		AppVersionCode:   req.VersionCode,
		DevModelKey:      req.DevModelKey,
		DevKey:           req.DevKey,
		AppVersionTarget: "",
		AppVersionArch:   "",
	})
	_, err = l.svcCtx.ResourceCtx.AddAppUpgradeGetStrategyReportLog(l.ctx, resource.AddAppUpgradeGetStrategyReportLogReq{
		CompanyId:           lnxInfo.CompanyId,
		AppKey:              lnxInfo.Key,
		AppType:             "lnx",
		Timestamp:           *timestamp,
		AppVersionId:        appVersionId,
		AppVersionCode:      req.VersionCode,
		DevModelKey:         req.DevModelKey,
		DevKey:              req.DevKey,
		AppVersionTarget:    "",
		AppVersionArch:      "",
		StrategyVersionId:   lnxStrategyInfo.LnxVersionId,
		StrategyVersionCode: lnxVersionInfo.VersionCode,
		StrategyId:          lnxStrategyInfo.Id,
	})

	res.Code = 200
	res.Msg = common.NewVersionMsg
	res.Data = types.GetLnxUpgradeInfoRespData{
		LnxKey:               lnxInfo.Key,
		PackageName:          lnxInfo.PackageName,
		VersionName:          lnxVersionInfo.VersionName,
		VersionCode:          lnxVersionInfo.VersionCode,
		UrlPath:              urlPath,
		UrlFileSize:          cloudFileInfo.Size,
		UrlFileMd5:           cloudFileInfo.Md5,
		UpgradeType:          lnxStrategyInfo.UpgradeType,
		PromptUpgradeContent: lnxStrategyInfo.PromptUpgradeContent,
	}
	return &res, nil
}

// ReturnUpgradeStrategyInfo
// 根据设备信息 获取对应可用的策略
func (l *GetLnxUpgradeInfoLogic) ReturnUpgradeStrategyInfo(lnxId int64, arch string, versionCode int64, appointVersionCode int64, devModelKey, devKey string) (resp *model.UpgradeLnxUpgradeStrategy, err error) {

	// 获取到大于请求版本的所有策略信息
	lnxStrategyList, err := l.svcCtx.ResourceCtx.GetLnxStrategyListByLnxIdAndArchAndVersion(l.ctx, lnxId, arch, versionCode)
	if err != nil {
		return nil, err
	}

	// 如果获取不到 则代表当前已经为最新版本
	if len(lnxStrategyList) == 0 {
		return nil, model.ErrNotFound
	}

	// 判断是否设置了期望版本， 如果设置了期望版本则直接获取期望版本的策略
	if appointVersionCode > 0 {
		for i := 0; i < len(lnxStrategyList); i++ {

			// 通过 Version id 获取到对应的版本信息
			lnxVersionInfo, err := l.svcCtx.ResourceCtx.GetLnxVersionInfoById(l.ctx, lnxStrategyList[i].LnxVersionId)
			if err != nil {
				return nil, err
			}

			if lnxVersionInfo.VersionCode == appointVersionCode {
				// 判断期望版本的策略的 设备相关策略是否符合
				strategyOk, err := l.CheckUpgradeStrategy(lnxStrategyList[i], devModelKey, devKey)
				if err != nil {
					return nil, err
				}

				// 如果符合条件则，返回期望版本的策略
				if strategyOk {
					// 判断是否符合灰度策略
					grayOk, err := l.CheckUpgradeStrategyGray(lnxStrategyList[i])
					if err != nil {
						return nil, err
					}
					if grayOk {
						return lnxStrategyList[i], nil
					}
				}

				break
			}
		}

		return nil, model.ErrNotFound
	}

	// 未设置期望版本 则获取当前版本的策略
	for i := 0; i < len(lnxStrategyList); i++ {
		// 判断是否符合条件
		strategyOk, err := l.CheckUpgradeStrategy(lnxStrategyList[i], devModelKey, devKey)
		if err != nil {
			return nil, err
		}

		// 如果符合条件则，返回期望版本的策略
		if strategyOk {
			// 判断是否符合灰度策略
			grayOk, err := l.CheckUpgradeStrategyGray(lnxStrategyList[i])
			if err != nil {
				return nil, err
			}
			if grayOk {
				return lnxStrategyList[i], nil
			}
		}
	}

	return nil, model.ErrNotFound

}

// CheckUpgradeStrategy
// 效验此版本是否符合规则
func (l *GetLnxUpgradeInfoLogic) CheckUpgradeStrategy(strategyInfo *model.UpgradeLnxUpgradeStrategy, devModelKey, devKey string) (ok bool, err error) {
	// 如果策略设备分组为 0 则代表全部设备
	if strategyInfo.UpgradeDevType == 0 {
		// 0: 全部设备
		return true, nil
	}

	// 效验设备唯一标识是否符合策略
	// 指定分组
	if strategyInfo.UpgradeDevType == 1 {
		// 根据策略中的设备分组id 获取到分组信息 跟传过来的进行比对
		devGroupIds, err := common.ConvertToInt64Slice(strategyInfo.UpgradeDevData)
		if err != nil {
			// 如果转换过程中出现错误，则输出错误信息
			return false, err
		}

		// 根据用户传的设备唯一标识查询出设备唯一标识对应的分组信息
		devGroupList, err := l.svcCtx.ResourceCtx.GetDevGroupListByDevKey(l.ctx, devKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {

		} else if err != nil {
			return false, err
		}

		// 循环查询数据，并做数据比对
		if len(devGroupList) > 0 && len(devGroupIds) > 0 {
			for _, devGroupId := range devGroupIds {
				for _, devGroup := range devGroupList {
					if devGroup.Id == devGroupId {
						return true, nil
					}
				}
			}
		}
	}

	// 指定机型
	if strategyInfo.UpgradeDevType == 2 {
		// 根据策略中的设备机型id 获取到机型信息 跟传过来的进行比对
		devModelIds, err := common.ConvertToInt64Slice(strategyInfo.UpgradeDevData)
		if err != nil {
			// 如果转换过程中出现错误，则输出错误信息
			return false, err
		}

		// 循环查询数据，并做数据比对
		for _, devModelId := range devModelIds {
			devModelInfo, err := l.svcCtx.ResourceCtx.GetDevModelInfoById(l.ctx, devModelId)
			if err != nil && errors.Is(err, model.ErrNotFound) {
				continue
			} else if err != nil {
				return false, err
			}

			if devModelInfo.Key == devModelKey {
				return true, nil
			}
		}

	}

	return false, nil
}

// CheckUpgradeStrategyGray 效验灰度策略是否符合规则
func (l *GetLnxUpgradeInfoLogic) CheckUpgradeStrategyGray(strategyInfo *model.UpgradeLnxUpgradeStrategy) (ok bool, err error) {
	// 通过策略信息判断是否开启了灰度策略
	if strategyInfo.IsGray == 0 {
		return true, nil
	}

	// 查询出灰度策略信息
	grayStrategyList, err := l.svcCtx.ResourceCtx.GetLnxGrayStrategyListByIds(l.ctx, strategyInfo.GrayData)
	if err != nil {
		return false, err
	}

	// 循环查询数据，并做数据比对
	for i := 0; i < len(grayStrategyList); i++ {
		// 判断是否是开启
		if grayStrategyList[i].Enable == 0 {
			continue
		}

		// 获取是否在时间范围外
		if common.IsTimeOutside(grayStrategyList[i].BeginDatetime, grayStrategyList[i].EndDatetime) {
			continue
		} else {
			// 判断缓存数据是否符合 使用计数器方法
			increment, err := l.svcCtx.ResourceCtx.LnxGrayStrategyIncrement(l.ctx, grayStrategyList[i])
			if err != nil {
				return false, err
			}
			if increment {
				return true, nil
			}
			// 在时间范围内，如果超出频率就直接返回 不符合
			return false, nil
		}

	}

	return true, nil
}

// CheckUpgradeStrategyFlowLimit 效验频控策略是否符合规则
// true 则代表可以反馈策略信息， false 为被限制
func (l *GetLnxUpgradeInfoLogic) CheckUpgradeStrategyFlowLimit(strategyInfo *model.UpgradeLnxUpgradeStrategy) (ok bool, err error) {
	// 通过策略信息判断是否开启了频控策略
	if strategyInfo.IsFlowLimit == 0 {
		return true, nil
	}

	// 查询出频控策略信息
	flowLimitStrategyList, err := l.svcCtx.ResourceCtx.GetLnxFlowLimitStrategyListByIds(l.ctx, strategyInfo.FlowLimitData)
	if err != nil {
		return false, err
	}

	// 循环查询数据，并做数据比对
	for i := 0; i < len(flowLimitStrategyList); i++ {
		// 判断是否是开启
		if flowLimitStrategyList[i].Enable == 0 {
			continue
		}

		// 判断时间范围
		now := time.Now()
		startTime, err := time.Parse("15:04:05", flowLimitStrategyList[i].BeginTime)
		if err != nil {
			return false, fmt.Errorf("解析开始时间失败: %w", err)
		}
		endTime, err := time.Parse("15:04:05", flowLimitStrategyList[i].EndTime)
		if err != nil {
			return false, fmt.Errorf("解析结束时间失败: %w", err)
		}
		start := time.Date(now.Year(), now.Month(), now.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, now.Location())
		end := time.Date(now.Year(), now.Month(), now.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, now.Location())
		if now.Before(start) || now.After(end) {
			continue
		} else {
			// 判断缓存数据是否符合 使用计数器方法
			increment, err := l.svcCtx.ResourceCtx.LnxFlowLimitStrategyIncrement(l.ctx, flowLimitStrategyList[i])
			if err != nil {
				return false, err
			}
			if increment {
				return true, nil
			}

			// 在时间范围内，如果超出频率就直接返回 不符合
			return false, nil
		}

	}

	return true, nil
}
