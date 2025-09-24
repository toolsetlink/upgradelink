package url

import (
	"context"
	"errors"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/config"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUrlUpgradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUrlUpgradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUrlUpgradeInfoLogic {
	return &GetUrlUpgradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUrlUpgradeInfoLogic) GetUrlUpgradeInfo(req *types.GetUrlUpgradeInfoReq) (resp *types.GetUrlUpgradeInfoResp, err error) {

	// 请求参数效验
	if req.UrlKey == "" {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err200Msg, config.Err200Docs)
	}
	if req.VersionCode == 0 {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrParamInvalid, config.Err201Msg, config.Err201Docs)
	}

	var res types.GetUrlUpgradeInfoResp

	// 通过唯一标识 获取到对应的应用信息
	urlInfo, err := l.svcCtx.ResourceCtx.GetUrlInfoByKey(l.ctx, req.UrlKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrNotFound, config.Err202Msg, config.Err202Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	// 查询应用版本表，判断是否有大于当前版本的  没有的话则代表当前就是最高版本
	_, err = l.svcCtx.ResourceCtx.GetUrlVersionListByUrlIdAndVersionCode(l.ctx, urlInfo.Id, req.VersionCode)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		res.Code = 0
		res.Msg = "已经是最新版本"
		return &res, nil
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	// 推出可使用的最高版本的升级策略
	urlStrategyInfo, err := l.ReturnUpgradeStrategyInfo(urlInfo.Id, req.VersionCode, req.AppointVersionCode, req.DevModelKey, req.DevKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		res.Code = 0
		res.Msg = "已经是最新版本"
		return &res, nil
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	// 到这个地方的时候 说明前置条件已经都通过了，在这个位置再去判断 策略的频控配置是否符合
	flowLimitOk, err := l.CheckUpgradeStrategyFlowLimit(urlStrategyInfo)
	if err != nil {
		return nil, err
	}
	if !flowLimitOk {
		// 被频控拦住 返回 429
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrStrategyTooManyReq, config.Err42901Msg, config.Err42901Docs)
	}

	// 通过升级版本 id 查询出对应版本信息 获取文件下载地址
	urlVersionInfo, err := l.svcCtx.ResourceCtx.GetUrlVersionInfoById(l.ctx, urlStrategyInfo.UrlVersionId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	} else if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}

	// 插入获取日志上报
	timestamp, err := common.ParseRFC3339ToTime(time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, config.Err1Msg, config.Err1Docs)
	}
	// 获取应用版本 id
	appVersionId, err := l.svcCtx.ResourceCtx.GetAppVersionIdByReport(l.ctx, resource.GetAppVersionIdByReportReq{
		AppKey:           urlInfo.Key,
		AppVersionCode:   req.VersionCode,
		DevModelKey:      req.DevModelKey,
		DevKey:           req.DevKey,
		AppVersionTarget: "",
		AppVersionArch:   "",
	})
	_, err = l.svcCtx.ResourceCtx.AddAppUpgradeGetStrategyReportLog(l.ctx, resource.AddAppUpgradeGetStrategyReportLogReq{
		CompanyId:           urlInfo.CompanyId,
		AppKey:              urlInfo.Key,
		AppType:             "url",
		Timestamp:           *timestamp,
		AppVersionId:        appVersionId,
		AppVersionCode:      req.VersionCode,
		DevModelKey:         req.DevModelKey,
		DevKey:              req.DevKey,
		AppVersionTarget:    "",
		AppVersionArch:      "",
		StrategyVersionId:   urlStrategyInfo.UrlVersionId,
		StrategyVersionCode: urlVersionInfo.VersionCode,
		StrategyId:          urlStrategyInfo.Id,
	})

	res.Code = 200
	res.Msg = "获取到新版本"
	res.Data = types.GetUrlUpgradeInfoRespData{
		UrlKey:               urlInfo.Key,
		VersionName:          urlVersionInfo.VersionName,
		VersionCode:          urlVersionInfo.VersionCode,
		UrlPath:              urlVersionInfo.UrlPath,
		UpgradeType:          urlStrategyInfo.UpgradeType,
		PromptUpgradeContent: urlStrategyInfo.PromptUpgradeContent,
	}
	return &res, nil
}

// ReturnUpgradeStrategyInfo 根据设备信息 获取对应可用的策略
func (l *GetUrlUpgradeInfoLogic) ReturnUpgradeStrategyInfo(urlId int64, versionCode int64, appointVersionCode int64, devModelKey, devKey string) (resp *model.UpgradeUrlUpgradeStrategy, err error) {

	// 获取到大于请求版本的所有策略信息
	urlStrategyList, err := l.svcCtx.ResourceCtx.GetUrlStrategyListByUrlIdAndVersion(l.ctx, urlId, versionCode)
	if err != nil {
		return nil, err
	}
	// 如果获取不到 则代表当前已经为最新版本
	if len(urlStrategyList) == 0 {
		return nil, model.ErrNotFound
	}

	// 判断是否设置了期望版本， 如果设置了期望版本则直接获取期望版本的策略
	if appointVersionCode > 0 {
		for i := 0; i < len(urlStrategyList); i++ {
			// 通过 Version id 获取到对应的版本信息， 通过 VersionId 获取到 VersionCode
			urlVersionInfo, err := l.svcCtx.ResourceCtx.GetUrlVersionInfoById(l.ctx, urlStrategyList[i].UrlVersionId)
			if err != nil {
				return nil, err
			}
			if urlVersionInfo.VersionCode == appointVersionCode {
				// 判断期望版本的策略的 设备相关策略是否符合
				strategyOk, err := l.CheckUpgradeStrategy(urlStrategyList[i], devModelKey, devKey)
				if err != nil {
					return nil, err
				}

				// 如果符合条件则，返回期望版本的策略，如果不符合策略则直接返回无升级策略
				if strategyOk {
					// 判断是否符合灰度策略
					grayOk, err := l.CheckUpgradeStrategyGray(urlStrategyList[i])
					if err != nil {
						return nil, err
					}
					if grayOk {
						return urlStrategyList[i], nil
					}
				}

				break
			}
		}
		return nil, model.ErrNotFound
	}

	// 未设置期望版本 则获取当前版本的策略
	for i := 0; i < len(urlStrategyList); i++ {
		// 判断是否符合条件
		strategyOk, err := l.CheckUpgradeStrategy(urlStrategyList[i], devModelKey, devKey)
		if err != nil {
			return nil, err
		}

		// 如果符合条件则，返回期望版本的策略
		if strategyOk {
			// 判断是否符合灰度策略
			grayOk, err := l.CheckUpgradeStrategyGray(urlStrategyList[i])
			if err != nil {
				return nil, err
			}
			if grayOk {
				return urlStrategyList[i], nil
			}
		}
	}

	return nil, model.ErrNotFound

}

// CheckUpgradeStrategy 效验此版本是否符合规则
func (l *GetUrlUpgradeInfoLogic) CheckUpgradeStrategy(strategyInfo *model.UpgradeUrlUpgradeStrategy, devModelKey, devKey string) (ok bool, err error) {
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
func (l *GetUrlUpgradeInfoLogic) CheckUpgradeStrategyGray(strategyInfo *model.UpgradeUrlUpgradeStrategy) (ok bool, err error) {
	// 通过策略信息判断是否开启了灰度策略
	if strategyInfo.IsGray == 0 {
		return true, nil
	}

	// 查询出灰度策略信息
	grayStrategyList, err := l.svcCtx.ResourceCtx.GetUrlGrayStrategyListByIds(l.ctx, strategyInfo.GrayData)
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
			increment, err := l.svcCtx.ResourceCtx.UrlGrayStrategyIncrement(l.ctx, grayStrategyList[i])
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
func (l *GetUrlUpgradeInfoLogic) CheckUpgradeStrategyFlowLimit(strategyInfo *model.UpgradeUrlUpgradeStrategy) (ok bool, err error) {
	// 通过策略信息判断是否开启了频控策略
	if strategyInfo.IsFlowLimit == 0 {
		return true, nil
	}

	// 查询出频控策略信息
	flowLimitStrategyList, err := l.svcCtx.ResourceCtx.GetUrlFlowLimitStrategyListByIds(l.ctx, strategyInfo.FlowLimitData)
	if err != nil {
		return false, err
	}
	fmt.Println("flowLimitStrategyList", flowLimitStrategyList)

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
			increment, err := l.svcCtx.ResourceCtx.UrlFlowLimitStrategyIncrement(l.ctx, flowLimitStrategyList[i])
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
